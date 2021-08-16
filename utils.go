package main

import (
	"github.com/yosssi/gcss"
	"log"
	"os"
)

func GetFilePaths(filePath string) ([]string, error) {
	currDir, _ := os.Getwd()

	var fp []string
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	if err := f.Close(); err != nil {
		return nil, err
	}

	for _, file := range files {
		rootPath := currDir + "/web/style/scss/" + file.Name()
		fp = append(fp, rootPath)
	}
	return fp, err
}

func CompileToSass() error {
	currDir, _ := os.Getwd()
	scssDir := currDir + "/web/style/scss"
	paths, err := GetFilePaths(scssDir)
	if err != nil {
		log.Fatalln("Eroare de compilare la sass " + err.Error())
		return err
	}

	for _, file := range paths {
		if _, err := gcss.CompileFile(file); err != nil {
			panic(err)
		}
	}
	return nil
}
