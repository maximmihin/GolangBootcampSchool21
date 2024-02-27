package app

import (
	"errors"
	"fmt"
	"log"
	"readDB/pkg/repository"
	"readDB/pkg/repository/dbs"
)

func Run(dbFileNameWithExtension string) {

	db, err := repository.ConnectDB(dbFileNameWithExtension)
	if err != nil {
		log.Fatal(err)
	}

	recipes, err := db.GetAllRecipes()
	if err != nil {
		log.Fatal(err)
	}

	switch db.(type) {
	case dbs.JsonBase:
		fmt.Println(string(recipes.PrettyXML("    ")))
	case dbs.XmlBase:
		fmt.Println(string(recipes.PrettyJSON("    ")))
	default:
		log.Fatal(errors.New("неожиданный тип базы данных"))
	}
}
