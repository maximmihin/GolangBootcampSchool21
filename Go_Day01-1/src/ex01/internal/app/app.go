package app

import (
	"compareDB/pkg/repository"
	"fmt"
	"log"
)

func Run(oldDBFileName string, newDBFileName string) {
	oldDB, err := repository.ConnectDB(oldDBFileName)
	if err != nil {
		log.Fatal(err)
	}

	newDB, err := repository.ConnectDB(newDBFileName)
	if err != nil {
		log.Fatal(err)
	}

	oldRecipes, err := oldDB.GetAllRecipes()
	if err != nil {
		log.Fatal(err)
	}

	newRecipes, err := newDB.GetAllRecipes()
	if err != nil {
		log.Fatal(err)
	}

	diffList := oldRecipes.Diff(newRecipes)

	for i := 0; i < len(diffList); i++ {
		fmt.Println(diffList[i])
	}

}
