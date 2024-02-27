package entityes

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

type RecipesBook struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

func (fr *RecipesBook) PrettyJSON(indent string) []byte {
	tmpJSON, err := json.MarshalIndent(fr, "", indent)
	if err != err {
		panic(err)
	}
	return tmpJSON
}

func (fr *RecipesBook) PrettyXML(indent string) []byte {
	tmpXML, err := xml.MarshalIndent(fr, "", indent)
	if err != err {
		panic(err)
	}
	return tmpXML
}

func (fr *RecipesBook) Diff(sr *RecipesBook) []string {
	if fr == nil || sr == nil {
		return nil
	}

	firstBookIndex := createCakesIndex(fr)
	secondBookIndex := createCakesIndex(sr)

	setFirstBookCakes := mapset.NewSetFromMapKeys(firstBookIndex)
	setSecondBookCakes := mapset.NewSetFromMapKeys(secondBookIndex)

	addedCakes := setSecondBookCakes.Difference(setFirstBookCakes)
	removedCakes := setFirstBookCakes.Difference(setSecondBookCakes)
	intersectedCakes := setFirstBookCakes.Intersect(setSecondBookCakes)

	diffList := make([]string, 0, addedCakes.Cardinality()+removedCakes.Cardinality())

	diffList = append(diffList, getDiffCakeList(addedCakes, "ADDED")...)
	diffList = append(diffList, getDiffCakeList(removedCakes, "REMOVED")...)

	for _, cake := range fr.Cakes {
		if intersectedCakes.Contains(cake.Name) {
			firstCake := firstBookIndex[cake.Name]
			secondCake := secondBookIndex[cake.Name]

			diffList = append(diffList, firstCake.diff(secondCake)...)
		}
	}

	return diffList
}

// Diff utils
func createCakesIndex(Base *RecipesBook) map[string]*Cake {
	CakesIndex := make(map[string]*Cake, len(Base.Cakes))

	for i := 0; i < len(Base.Cakes); i++ {
		CakesIndex[Base.Cakes[i].Name] = &Base.Cakes[i]
	}

	return CakesIndex
}

func getDiffCakeList(cakeList mapset.Set[string], typeDiff string) []string {
	diffList := make([]string, 0, cakeList.Cardinality())

	for cake := range cakeList.Iter() {
		diffList = append(diffList, fmt.Sprintf("%s cake \"%s\"", typeDiff, cake))
	}

	return diffList
}
