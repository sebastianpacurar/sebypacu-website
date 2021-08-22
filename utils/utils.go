package utils

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, bp, _, _ = runtime.Caller(0)
	BasePath    = filepath.Dir(bp)
)

func GetFilePaths(filePath string) ([]string, error) {

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
		rootPath := BasePath + "/web/style/scss/" + file.Name()
		fp = append(fp, rootPath)
	}
	return fp, err
}

func ScrollToTop(ctx app.Context, e app.Event) {
	app.Window().ScrollToID("app-header")
}
