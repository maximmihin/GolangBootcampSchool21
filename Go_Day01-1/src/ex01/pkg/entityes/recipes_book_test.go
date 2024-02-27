package entityes

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var originalXmlBase = `<recipes>
	<cake>
		<name>Red Velvet Strawberry Cake</name>
		<stovetime>40 min</stovetime>
		<ingredients>
			<item>
				<itemname>Flour</itemname>
				<itemcount>3</itemcount>
				<itemunit>cups</itemunit>
			</item>
			<item>
				<itemname>Vanilla extract</itemname>
				<itemcount>1.5</itemcount>
				<itemunit>tablespoons</itemunit>
			</item>
			<item>
				<itemname>Strawberries</itemname>
				<itemcount>7</itemcount>
				<itemunit></itemunit>
			</item>
			<item>
				<itemname>Cinnamon</itemname>
				<itemcount>1</itemcount>
				<itemunit>pieces</itemunit>
			</item>
		</ingredients>
	</cake>
	<cake>
		<name>Blueberry Muffin Cake</name>
		<stovetime>30 min</stovetime>
		<ingredients>
			<item>
				<itemname>Baking powder</itemname>
				<itemcount>3</itemcount>
				<itemunit>teaspoons</itemunit>
			</item>
			<item>
				<itemname>Brown sugar</itemname>
				<itemcount>0.5</itemcount>
				<itemunit>cup</itemunit>
			</item>
			<item>
				<itemname>Blueberries</itemname>
				<itemcount>1</itemcount>
				<itemunit>cup</itemunit>
			</item>
		</ingredients>
	</cake>
</recipes>`

var stolenJsonBase = `{
	"cake": [
		{
			"name": "Red Velvet Strawberry Cake",
			"time": "45 min",
			"ingredients": [
				{
					"ingredient_name": "Flour",
					"ingredient_count": "2",
					"ingredient_unit": "mugs"
				},
				{
					"ingredient_name": "Strawberries",
					"ingredient_count": "8"
				},
				{
					"ingredient_name": "Coffee Beans",
					"ingredient_count": "2.5",
					"ingredient_unit": "tablespoons"
				},
				{
					"ingredient_name": "Cinnamon",
					"ingredient_count": "1"
				}
			]
		},
		{
			"name": "Moonshine Muffin",
			"time": "30 min",
			"ingredients": [
				{
					"ingredient_name": "Brown sugar",
					"ingredient_count": "1",
					"ingredient_unit": "mug"
				},
				{
					"ingredient_name": "Blueberries",
					"ingredient_count": "1",
					"ingredient_unit": "mug"
				}
			]
		}
	]
}`

var xmlMastBe = RecipesBook{
	XMLName: xml.Name{
		Space: "",
		Local: "recipes",
	},
	Cakes: []Cake{
		{
			Name: "Red Velvet Strawberry Cake",
			Time: "40 min",
			Ingredients: []Ingredient{
				{Name: "Flour", Count: 3, Unit: "cups"},
				{Name: "Vanilla extract", Count: 1.5, Unit: "tablespoons"},
				{Name: "Strawberries", Count: 7},
				{Name: "Cinnamon", Count: 1, Unit: "pieces"},
			},
		},
		{
			Name: "Blueberry Muffin Cake",
			Time: "30 min",
			Ingredients: []Ingredient{
				{Name: "Baking powder", Count: 3, Unit: "teaspoons"},
				{Name: "Brown sugar", Count: 0.5, Unit: "cup"},
				{Name: "Blueberries", Count: 1, Unit: "cup"},
			},
		},
	},
}

var jsonMastBe = RecipesBook{
	XMLName: xml.Name{
		Space: "",
		Local: "",
	},
	Cakes: []Cake{
		{
			Name: "Red Velvet Strawberry Cake",
			Time: "45 min",
			Ingredients: []Ingredient{
				{Name: "Flour", Count: 2, Unit: "mugs"},
				{Name: "Strawberries", Count: 8},
				{Name: "Coffee Beans", Count: 2.5, Unit: "tablespoons"},
				{Name: "Cinnamon", Count: 1},
			},
		},
		{
			Name: "Moonshine Muffin",
			Time: "30 min",
			Ingredients: []Ingredient{
				{Name: "Brown sugar", Count: 1, Unit: "mug"},
				{Name: "Blueberries", Count: 1, Unit: "mug"},
			},
		},
	},
}

func TestRecipesBook_conversions(t *testing.T) {

	// xml -> struct
	var original RecipesBook
	_ = xml.Unmarshal([]byte(originalXmlBase), &original)

	if !reflect.DeepEqual(original, xmlMastBe) {
		t.Errorf("ошибка при [xml -> struct]")
	}

	// struct -> xml
	tmpXml := string(xmlMastBe.PrettyXML("\t"))
	tmp := strings.Clone(originalXmlBase)
	if tmpXml != tmp {
		t.Errorf("ошибка при [struct -> xml]")
	}

	// json -> struct
	var stolen RecipesBook
	_ = json.Unmarshal([]byte(stolenJsonBase), &stolen)

	if !reflect.DeepEqual(stolen, jsonMastBe) {
		t.Errorf("ошибка при [json -> struct]")
	}

	// struct -> json
	tmpJson := string(jsonMastBe.PrettyJSON("\t"))
	tmp2 := strings.Clone(stolenJsonBase)
	if tmpJson != tmp2 {
		t.Errorf("ошибка при [struct -> json]")
	}
}

var diffListMastBe = []string{
	"ADDED cake \"Moonshine Muffin\"",
	"REMOVED cake \"Blueberry Muffin Cake\"",
	"CHANGED cooking time for cake \"Red Velvet Strawberry Cake\" - \"45 min\" instead of \"40 min\"",
	"ADDED ingredient \"Coffee Beans\" for cake \"Red Velvet Strawberry Cake\"",
	"REMOVED ingredient \"Vanilla extract\" for cake \"Red Velvet Strawberry Cake\"",
	"CHANGED unit for ingredient \"Flour\" for cake \"Red Velvet Strawberry Cake\" - \"mugs\" instead of \"cups\"",
	"CHANGED unit count for ingredient \"Flour\" for cake \"Red Velvet Strawberry Cake\" - \"2\" instead of \"3\"",
	"CHANGED unit count for ingredient \"Strawberries\" for cake \"Red Velvet Strawberry Cake\" - \"8\" instead of \"7\"",
	"REMOVED unit \"pieces\" for ingredient \"Cinnamon\" for cake \"Red Velvet Strawberry Cake\"",
}

func TestRecipesBook_DiffList(t *testing.T) {

	var original RecipesBook
	err := xml.Unmarshal([]byte(originalXmlBase), &original)
	if err != nil {
		panic(err)
	}

	var stolen RecipesBook
	err1 := json.Unmarshal([]byte(stolenJsonBase), &stolen)
	if err1 != nil {
		panic(err1)
	}

	diffList := original.Diff(&stolen)

	for i, diffStr := range diffList {
		if diffStr != diffListMastBe[i] {
			fmt.Println(diffStr, "not equal", diffListMastBe[i])
			t.Errorf("неверный diff")
		}
	}
}
