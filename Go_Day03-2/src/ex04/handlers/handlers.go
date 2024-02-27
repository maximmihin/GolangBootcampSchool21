package handlers

import (
	"encoding/json"
	"ex04/db"
	"ex04/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func HandleHtmlPlaces(es *db.Elastic, indexTemplate *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := r.URL.Query().Get("page")
		page, err := strconv.Atoi(val)

		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid 'page' value: '%v'", val), http.StatusBadRequest)
			return
		}

		if page == 0 {
			page = 1
		}
		//start := time.Now()
		places, total, err := es.GetPlaces(10, 10*(page-1))
		//duration := time.Since(start)
		//log.Printf("время выполнения запроса %d ms\n", duration.Milliseconds())
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid 'page' value: '%v'", val), http.StatusBadRequest)
			return
		}
		lastPage := int(math.Ceil(float64(total) / float64(10)))
		if page > lastPage {
			http.Error(w, fmt.Sprintf("Invalid 'page' value: '%v'", val), http.StatusBadRequest)
			return
		}

		testStruct := models.PlacesModel{
			Name:     "Places",
			Total:    total,
			Places:   places,
			LastPage: lastPage,
		}

		if page < 2 {
			testStruct.PrevPage = 0
		} else {
			testStruct.PrevPage = page - 1
		}

		if page == lastPage {
			testStruct.NextPage = 0
		} else {
			testStruct.NextPage = page + 1
		}

		err2 := indexTemplate.Execute(w, testStruct)
		if err2 != nil {
			log.Fatal(err)
		}
	}

}

func HandleApiPlaces(es *db.Elastic) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		val := r.URL.Query().Get("page")
		page, err := strconv.Atoi(val)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("Invalid 'page' value: '%v'", val),
			})
			return
		}

		if page == 0 {
			page = 1
		}
		//start := time.Now()
		places, total, err := es.GetPlaces(10, 10*(page-1))

		lastPage := int(math.Ceil(float64(total) / float64(10)))
		if page > lastPage {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("Invalid 'page' value: '%v'", val),
			})
			return
		}

		//duration := time.Since(start)
		//log.Printf("время выполнения запроса %d ms\n", duration.Milliseconds())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("Invalid 'page' value: '%v'", val),
			})
			return
		}

		testStruct := models.PlacesModel{
			Name:     "Places",
			Total:    total,
			Places:   places,
			LastPage: lastPage,
		}

		if page < 2 {
			testStruct.PrevPage = 0
		} else {
			testStruct.PrevPage = page - 1
		}

		if page == lastPage {
			testStruct.NextPage = 0
		} else {
			testStruct.NextPage = page + 1
		}

		err = json.NewEncoder(w).Encode(testStruct)
		if err != nil {
			http.Error(w, "encode error", http.StatusInternalServerError)
			return
		}
	}
}

func HandleApiGetToken(secretKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key := []byte(secretKey)
		token := jwt.New(jwt.SigningMethodHS256)

		bearerToken, err := token.SignedString(key)
		if err != nil {
			http.Error(w, "sign token error", http.StatusInternalServerError)
			return
		}

		tokenModel := models.TokenModel{Token: bearerToken}

		err = json.NewEncoder(w).Encode(tokenModel)
		if err != nil {
			http.Error(w, "encode error", http.StatusInternalServerError)
			return
		}
	}
}

func HandleApiRecommend(es *db.Elastic) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rLat := r.URL.Query().Get("lat")
		lat, err1 := strconv.ParseFloat(rLat, 64)

		rLon := r.URL.Query().Get("lon")
		lon, err2 := strconv.ParseFloat(rLon, 64)

		if err1 != nil || err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("Invalid values: lat: '%v'; lon: '%v'", lat, lon),
			})
			return
		}

		places, _, err := es.GetRecommendPlaces(3, lat, lon)
		if err != nil {
			return
		}

		testStruct := models.RecommendedPlacesModel{
			Name:   "places",
			Places: places,
		}

		err = json.NewEncoder(w).Encode(testStruct)
		if err != nil {
			http.Error(w, "encode error", http.StatusInternalServerError)
			return
		}
	}
}

func MiddlewareJwtCheck(secretKey string, next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(map[string]string{
				"error": fmt.Sprintf("Where is your token?"),
			})
			return
		}

		gotToken := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := jwt.Parse(gotToken, func(token *jwt.Token) (interface{}, error) {

			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secretKey), nil
		})
		if err != nil {
			// write error auth
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(map[string]string{
				"error": fmt.Sprintf("Token is shit"),
			})
			return
		}
		next(writer, request)
		return
	}
}
