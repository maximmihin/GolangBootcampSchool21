package dbs

import (
	"compareDB/pkg/entityes"
	"encoding/xml"
)

type XmlBase struct {
	DbFile *[]byte
}

func (base XmlBase) GetAllRecipes() (*entityes.RecipesBook, error) {
	recipes := new(entityes.RecipesBook)
	err := xml.Unmarshal(*base.DbFile, recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
