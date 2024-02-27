package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"compareDB/pkg/entityes"
	"compareDB/pkg/repository/dbs"
)

type DBReader interface {
	GetAllRecipes() (*entityes.RecipesBook, error)
}

func ConnectDB(fileName string) (DBReader, error) {
	if fileName == "" {
		return nil, fmt.Errorf("нужно передать файл аргументом при помощи флага -f")
	}

	tmpDBFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("не получилось открыть или прочитать файл")
	}

	extension := filepath.Ext(fileName)
	switch extension {
	case ".json":
		return dbs.JsonBase{DbFile: &tmpDBFile}, nil
	case ".xml":
		return dbs.XmlBase{DbFile: &tmpDBFile}, nil
	default:
		return nil, fmt.Errorf("неизвестный формат, допустимы: .xml и .json")
	}
}
