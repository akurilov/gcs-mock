package main

import (
	"github.com/akurilov/gcs-mock/pkg"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
)

var (
	log = initLogger()
)

func initLogger() *zap.Logger {
	l, e := zap.NewProduction()
	if e != nil {
		panic(e)
	}
	return l
}

func main() {
	defer log.Sync()
	dataDir, _ := os.Getwd()
	if len(os.Args) > 1 {
		dataDir, _ = filepath.Abs(os.Args[1])
	}
	http.HandleFunc("/", pkg.Handler(dataDir))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal("", zap.Error(err))
}
