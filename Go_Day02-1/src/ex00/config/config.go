package config

import (
	"errors"
	"flag"
)

type Params struct {
	DirFlag        bool
	FileFlag       bool
	SymlinksFlag   bool
	FilesExtension string
	DirectoryPath  string
}

func ParseInput() (p Params, err error) {
	flag.BoolVar(&p.FileFlag, "f", false, "")
	flag.BoolVar(&p.DirFlag, "d", false, "")
	flag.BoolVar(&p.SymlinksFlag, "sl", false, "")
	flag.StringVar(&p.FilesExtension, "ext", "", "")
	flag.Parse()

	p.DirectoryPath = flag.Arg(0)

	if p.DirectoryPath == "" {
		return p, errors.New("необходимо указать путь до директории")
	}

	if p.FileFlag == false && p.FilesExtension != "" {
		return p, errors.New("флаг -ext можно использовать только одновременно с флагом -f")
	}

	if p.FilesExtension != "" {
		p.FilesExtension = "." + p.FilesExtension
	}

	if p.FileFlag == false && p.DirFlag == false && p.SymlinksFlag == false {
		p.FileFlag, p.DirFlag, p.SymlinksFlag = true, true, true
	}

	return
}
