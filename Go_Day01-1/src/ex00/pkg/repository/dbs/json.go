package dbs

import (
	"encoding/json"
	"readDB/pkg/entityes"
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
