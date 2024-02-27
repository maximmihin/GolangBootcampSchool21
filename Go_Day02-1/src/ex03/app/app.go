package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)

func Run(path string, logFiles []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(logFiles))
	for _, file := range logFiles {
		go createArchive(path, file, wg)
	}
	wg.Wait()
}

func createArchive(prefixPath, fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	// получить MTIME файла
	fi, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	mTime := fi.ModTime()

	cleanName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	strUnixTimeStamp := strconv.FormatInt(mTime.Unix(), 10)
	archiveName := prefixPath + cleanName + "_" + strUnixTimeStamp + ".tar.gz"

	tarFlags := "-czvf"
	tarArgs := []string{tarFlags, archiveName, fileName}

	cmd := exec.Command("tar", tarArgs...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))

}

//
//func createArchive(prefixPath, fileName string, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	// открыть ридер исходного файла
//	sourceLog, err := os.Open(fileName)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer sourceLog.Close()
//
//	// получить MTIME файла
//	fi, err := os.Stat(fileName)
//	if err != nil {
//		log.Fatal(err)
//	}
//	mTime := fi.ModTime()
//
//	// создать новый файл под архив
//	cleanName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
//	strUnixTimeStamp := strconv.FormatInt(mTime.Unix(), 10)
//	archiveName := prefixPath + cleanName + "_" + strUnixTimeStamp + ".tar.gz"
//	archiveFile, err := os.Create(archiveName)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// открыть райтер в архивный файл
//	gw := gzip.NewWriter(archiveFile)
//	defer gw.Close()
//	tw := tar.NewWriter(gw)
//	defer tw.Close()
//
//	// писать в него архив
//
//	// записать заголовок
//	tarHeader, err := tar.FileInfoHeader(fi, fi.Name())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	tarHeader.Name = fileName
//
//	err = tw.WriteHeader(tarHeader)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// запись
//	_, err = io.Copy(tw, sourceLog)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
