package internal

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

const (
	uriPathPatternRaw = `/storage/v1/b/([a-z\d\-_.]{3,222})(/o/([\d\w][\d\w\-_.]{0,1000}))?`
)

var (
	uriPathPattern = regexp.MustCompile(uriPathPatternRaw)
)

func Handler(dataDir string) func(respWriter http.ResponseWriter, req *http.Request) {
	storage := &Storage{
		dataDir: dataDir,
	}
	return func(respWriter http.ResponseWriter, req *http.Request) {
		reqPath := req.URL.Path
		result := uriPathPattern.FindStringSubmatch(reqPath)
		bucket := result[1]
		object := result[3]
		if len(object) > 0 {
			handleObjectRequest(storage, respWriter, req, bucket, object)
		} else {
			handleBucketRequest(storage, respWriter, req, bucket)
		}
	}
}

func handleBucketRequest(storage *Storage, respWriter http.ResponseWriter, req *http.Request, bucket string) {
	switch req.Method {
	case "DELETE":
		err := storage.DeleteBucket(bucket)
		if err == nil {
			fmt.Print(respWriter, "")
		} else {
			respWriter.WriteHeader(500)
			fmt.Print(respWriter, "")
		}
	case "GET":
		if len(bucket) > 0 {
			err := storage.ReadBucket(bucket)
		} else {
			query := req.URL.Query()
			maxResults, err := strconv.Atoi(query.Get("maxResults"))
			if err != nil || maxResults < 1 {
				maxResults = 1000
			}
			pageToken := query.Get("pageToken")
			prefix := query.Get("prefix")
			storage.ListBuckets(maxResults, pageToken, prefix)

		}
	case "POST":
		bucketResource, err := internal.CreateBucket()
	}
}

func handleObjectRequest(
	storage *Storage, respWriter http.ResponseWriter, req *http.Request, bucket string, object string) {

}
