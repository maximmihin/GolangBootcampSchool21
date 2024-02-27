package dbs

import (
	"compareDB/pkg/entityes"
	"encoding/json"
)

type JsonBase struct {
	DbFile *[]byte
}

func (base JsonBase) GetAllRecipes() (*entityes.RecipesBook, error) {
	recipes := new(entityes.RecipesBook)
	err := json.Unmarshal(*base.DbFile, recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
