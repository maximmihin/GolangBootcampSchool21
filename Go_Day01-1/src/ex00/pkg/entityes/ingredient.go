package entityes

import "fmt"

type Ingredient struct {
	Name  string  `json:"ingredient_name" xml:"itemname"`
	Count float64 `json:"ingredient_count,string" xml:"itemcount"`
	Unit  string  `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

func (i *Ingredient) diff(si *Ingredient, parentCakeName string) []string {
	diffList := make([]string, 0)

	if i.Unit == "" && si.Unit != "" {
		tmpDiffStr := fmt.Sprintf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"", si.Unit, i.Name, parentCakeName)
		diffList = append(diffList, tmpDiffStr)
	} else if i.Unit != "" && si.Unit == "" {
		tmpDiffStr := fmt.Sprintf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"", i.Unit, i.Name, parentCakeName)
		diffList = append(diffList, tmpDiffStr)
	} else if i.Unit != si.Unit {
		tmpDiffStr := fmt.Sprintf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"", i.Name, parentCakeName, si.Unit, i.Unit)
		diffList = append(diffList, tmpDiffStr)
	}

	// else if i.Unit == "" && si.Unit == "" { ничего не делать }

	if i.Count != si.Count {
		diffList = append(diffList, fmt.Sprintf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%v\" instead of \"%v\"", i.Name, parentCakeName, si.Count, i.Count))
	}

	return diffList
}
