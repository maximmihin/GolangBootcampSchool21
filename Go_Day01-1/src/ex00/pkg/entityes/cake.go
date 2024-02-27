package entityes

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

func (c *Cake) diff(sc *Cake) []string {
	if c.Name != sc.Name {
		return nil
	}

	diffList := make([]string, 0)

	if c.Time != sc.Time {
		diffList = append(diffList, fmt.Sprintf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"", c.Name, sc.Time, c.Time))
	}

	firstIngredientsIndex := createIngredientsIndex(c)
	secondIngredientsIndex := createIngredientsIndex(sc)

	setFirstCakeIngredients := mapset.NewSetFromMapKeys(firstIngredientsIndex)
	setSecondCakeIngredients := mapset.NewSetFromMapKeys(secondIngredientsIndex)

	addedIngredients := setSecondCakeIngredients.Difference(setFirstCakeIngredients)
	removedIngredients := setFirstCakeIngredients.Difference(setSecondCakeIngredients)
	intersectedIngredients := setFirstCakeIngredients.Intersect(setSecondCakeIngredients)

	diffList = append(diffList, getDiffIngredientList(addedIngredients, removedIngredients, c.Name)...)

	for _, ingr := range c.Ingredients {
		if intersectedIngredients.Contains(ingr.Name) {
			oldIngr := firstIngredientsIndex[ingr.Name]
			newIngr := secondIngredientsIndex[ingr.Name]

			diffList = append(diffList, oldIngr.diff(newIngr, c.Name)...)
		}
	}

	return diffList
}

// Diff utils
func createIngredientsIndex(Cake *Cake) map[string]*Ingredient {
	CakesIndex := make(map[string]*Ingredient, len(Cake.Ingredients))

	for i := 0; i < len(Cake.Ingredients); i++ {
		CakesIndex[Cake.Ingredients[i].Name] = &Cake.Ingredients[i]
	}

	return CakesIndex
}

func getDiffIngredientList(addedIngredients, removedIngredients mapset.Set[string], parentName string) []string {
	diffList := make([]string, 0, addedIngredients.Cardinality()+removedIngredients.Cardinality())

	for ingr := range addedIngredients.Iter() {
		diffList = append(diffList, fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"", ingr, parentName))
	}

	for ingr := range removedIngredients.Iter() {
		diffList = append(diffList, fmt.Sprintf("REMOVED ingredient \"%s\" for cake \"%s\"", ingr, parentName))
	}

	return diffList
}
