package app

import (
	"fmt"
	"log"
	"myFind/config"
	"os"
	"path/filepath"
)

func Run(conf config.Params) {
	err := filepath.Walk(conf.DirectoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().Perm()&0444 != 0444 {
			return nil
		}

		if conf.DirFlag && info.IsDir() {
			fmt.Println(path)
		} else if conf.SymlinksFlag && info.Mode().Type() == os.ModeSymlink {
			linkVal, err := filepath.EvalSymlinks(path)
			if err != nil {
				fmt.Println(path, "-> [broken]")
			} else {
				fmt.Println(path, "->", linkVal)
			}
		} else if conf.FileFlag && info.Mode().IsRegular() {
			if conf.FilesExtension != "" {
				if conf.FilesExtension == filepath.Ext(path) {
					fmt.Println(path)
				}
			} else {
				fmt.Println(path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
