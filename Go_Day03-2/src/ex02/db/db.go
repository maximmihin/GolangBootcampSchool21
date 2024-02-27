package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"ex02/entities"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/segmentsortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"log"
	"os"
	"strconv"
)

const maxResultWindow = 20000

type Store interface {
	// GetPlaces returns a list of items, a total number of hits and (or) an error in case of one
	GetPlaces(limit int, offset int) ([]entities.Place, int, error)
	GetRecommendPlaces(limit uint) ([]entities.Place, int, error)
	TakePlaces(places []entities.Place) (int, error)
}

type Elastic struct {
	classicClient *elasticsearch.Client
	typedClient   *elasticsearch.TypedClient
}

func (e Elastic) GetPlaces(limit int, offset int) ([]entities.Place, int, error) {

	if limit <= 0 {
		return nil, 0, errors.New("limit mast be > 0")
	}
	if offset < 0 {
		return nil, 0, errors.New("offset mast be >= 0")
	}

	qry := &search.Request{
		Size: &limit,
		From: &offset,
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{Boost: nil, QueryName_: nil},
		},
		TrackTotalHits: true,
		Sort: []types.SortCombinations{
			types.SortOptions{
				SortOptions: map[string]types.FieldSort{
					"id": {
						Order: &sortorder.Asc,
					},
				},
			},
		},
		//Explain: some.Bool(true),
	}

	//tmp, _ := json.MarshalIndent(qry, "", "   ")
	//fmt.Println(string(tmp))

	res, err := e.typedClient.Search().
		Index("places").
		Request(qry).
		Do(context.Background())

	if err != nil {
		return nil, 0, err
	}

	//tmp, _ := json.MarshalIndent(res, "", "   ")
	//fmt.Println(string(tmp))

	rawPlaces := res.Hits.Hits

	places := make([]entities.Place, len(rawPlaces))
	for i := 0; i < len(res.Hits.Hits); i++ {
		err := json.Unmarshal(rawPlaces[i].Source_, &places[i])
		if err != nil {
			return nil, 0, err
		}
	}

	//data3, _ := json.MarshalIndent(places, "", "    ")
	//fmt.Println(string(data3))

	return places, int(res.Hits.Total.Value), nil
}

func (e Elastic) GetRecommendPlaces(limit int, lat, lon float64) ([]entities.Place, int, error) {

	qry := &search.Request{
		Size: &limit,
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{Boost: nil, QueryName_: nil},
		},
		TrackTotalHits: true,
		Sort: []types.SortCombinations{
			types.SortOptions{
				GeoDistance_: &types.GeoDistanceSort{
					GeoDistanceSort: map[string][]types.GeoLocation{
						"location": {
							types.LatLonGeoLocation{
								Lat: *some.Float64(lat),
								Lon: *some.Float64(lon),
							},
						},
					},
					Order:          &sortorder.SortOrder{Name: "asc"},
					Unit:           &distanceunit.DistanceUnit{Name: "km"},
					Mode:           &sortmode.SortMode{Name: "min"},
					DistanceType:   &geodistancetype.GeoDistanceType{Name: "arc"},
					IgnoreUnmapped: some.Bool(true),
				},
				//SortOptions: map[string]types.FieldSort{
				//	"id": {
				//		Order: &sortorder.Asc,
				//	},
				//},
			},
		},
	}

	//tmp, _ := json.MarshalIndent(qry, "", "   ")
	//fmt.Println(string(tmp))

	res, err := e.typedClient.Search().
		Index("places").
		Request(qry).
		Do(context.Background())

	if err != nil {
		return nil, 0, err
	}

	rawPlaces := res.Hits.Hits
	places := make([]entities.Place, len(rawPlaces))
	for i := 0; i < len(res.Hits.Hits); i++ {
		err := json.Unmarshal(rawPlaces[i].Source_, &places[i])
		if err != nil {
			return nil, 0, err
		}
	}

	//tmp, _ = json.MarshalIndent(res, "", "   ")
	//fmt.Println(string(tmp))

	return places, int(res.Hits.Total.Value), nil
}

func (e Elastic) TakePlaces(places []entities.Place) (uint64, error) {

	indexName := "places"

	IndexExist, err := e.isIndexExist(indexName)
	if err != nil {
		return 0, err
	}

	if IndexExist {
		err := e.deleteIndex(indexName)
		if err != nil {
			return 0, err
		}
		log.Println("delete old index")
		err = e.createIndex(indexName)
		if err != nil {
			return 0, err
		}
		log.Println("create new index")
	} else {
		err := e.createIndex(indexName)
		if err != nil {
			return 0, err
		}
		log.Println("create new index")
	}

	indexed, err := e.bulkPlaces(places)
	if err != nil {
		return indexed, err
	}
	log.Printf("upload %d places\n", indexed)

	return indexed, nil
}

func (e Elastic) isIndexExist(indexN string) (bool, error) {
	res, err := e.classicClient.Indices.Exists([]string{indexN})
	defer res.Body.Close()
	if err != nil {
		return false, errors.New(fmt.Sprintf("Cannot check index exists: %s", err))
	}
	return !res.IsError(), nil
}

func (e Elastic) createIndex(indexN string) error {

	tmpFile, err := os.ReadFile("schema.json")
	if err != nil {
		return err
	}

	var tmpMap *types.TypeMapping
	err = json.Unmarshal(tmpFile, &tmpMap)
	if err != nil {
		return err
	}

	req := &create.Request{
		Mappings: tmpMap,
		Settings: &types.IndexSettings{
			MaxResultWindow: some.Int(maxResultWindow),
			Sort: &types.IndexSegmentSort{
				Field: []string{"id"},
				Order: []segmentsortorder.SegmentSortOrder{
					{Name: "asc"},
				},
			},
		},
	}

	//tmp, _ := json.MarshalIndent(req, "", "    ")
	//fmt.Println(string(tmp))

	res, err := e.typedClient.Indices.Create(indexN).
		Request(req).
		Do(nil)

	if err != nil {
		return errors.New(fmt.Sprintf("Cannot create index: %s", err))
	}

	if !res.Acknowledged && res.Index != indexN {
		return errors.New(fmt.Sprintf("unexpected error during index creation, got : %#v", res))
	}
	return nil
}

func (e Elastic) deleteIndex(indexN string) error {
	res, err := e.classicClient.Indices.Delete([]string{indexN}, e.classicClient.Indices.Delete.WithIgnoreUnavailable(true))
	if err != nil || res.IsError() {
		log.Fatalf("Cannot delete index: %s", err)
	}
	res.Body.Close()

	return nil
}

func (e Elastic) bulkPlaces(places []entities.Place) (uint64, error) {
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      "places",
		Client:     e.classicClient,
		NumWorkers: 5,
	})
	if err != nil {
		return 0, err
	}
	defer bulkIndexer.Close(context.Background())

	for _, place := range places {

		jsonPlace, err := json.Marshal(place)
		if err != nil {
			return 0, err
		}

		err = bulkIndexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.FormatUint(uint64(place.ID), 10),
				Body:       bytes.NewReader(jsonPlace),
			})
		if err != nil {
			return 0, err
		}
	}

	biStats := bulkIndexer.Stats()
	if biStats.NumAdded != uint64(len(places)) {
		return 0, errors.New(fmt.Sprintf("добавлены не все файлы: %d вместо %d", biStats.NumAdded, len(places)))
	}

	return biStats.NumAdded, nil
}

func New() (*Elastic, error) {

	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	classicClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	typedClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return nil, err
	}

	res, err := classicClient.Info()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &Elastic{
			classicClient,
			typedClient,
		},
		nil
}
