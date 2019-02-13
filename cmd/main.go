package cmd

import (
	"github.com/akurilov/gcs-mock/internal"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

const (
	uriPathPatternRaw = `/storage/v1/b/([a-z\d\-_.]{3,222})(/o/([\d\w][\d\w\-_.]{0,1000}))?`
)

var (
	log            = initLogger()
	uriPathPattern = regexp.MustCompile(uriPathPatternRaw)
)

func initLogger() *zap.Logger {
	l, e := zap.NewProduction()
	if e != nil {
		panic(e)
	}
	return l
}

func gcsHandler(dataDir string) func(respWriter http.ResponseWriter, req *http.Request) {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		reqPath := req.URL.Path
		result := uriPathPattern.FindStringSubmatch(reqPath)
		bucket := result[1]
		object := result[3]
		if len(object) > 0 {
			handleObjectRequest(respWriter, req, bucket, object)
		} else {
			handleBucketRequest(respWriter, req, bucket)
		}
	}
}

func handleBucketRequest(respWriter http.ResponseWriter, req *http.Request, bucket string) {
	switch req.Method {
	case "DELETE":
		err := internal.DeleteBucket(bucket)
	case "GET":
		if len(bucket) > 0 {
			bucketResource, err := internal.ReadBucket(bucket)
		} else {
			pageToken := req.URL.Query().Get("pageToken")
			bucketResources, err := internal.ListBuckets(pageToken)
		}
	case "POST":
		bucketResource, err := internal.CreateBucket()
	}
}

func handleObjectRequest(respWriter http.ResponseWriter, req *http.Request, bucket string, object string) {

}

func main() {
	defer log.Sync()
	dataDir, _ := os.Getwd()
	if len(os.Args) > 0 {
		dataDir, _ = filepath.Abs(os.Args[1])
	}
	http.HandleFunc("/", gcsHandler(dataDir))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal("", zap.Error(err))
}
