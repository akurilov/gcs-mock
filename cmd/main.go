package cmd

import (
	"github.com/uber-go/zap"
	"log"
	"net/http"
	"regexp"
)

const (
	uriPathPatternRaw = `/storage/v1/b/([a-z\d\-_.]{3,222})(/o/([\d\w][\d\w\-_.]{0,1000}))?`
)

var (
	uriPathPattern = regexp.MustCompile(uriPathPatternRaw)
)

type Operation int

const (
	Create Operation = iota + 1
	Read
	Delete
	List
)

func gcsHandler(respWriter http.ResponseWriter, req *http.Request) {
	reqPath := req.URL.Path
	result := uriPathPattern.FindStringSubmatch(reqPath)
	bucket := result[1]
	object := result[3]
	if len(bucket) > 0 {
		if len(object) > 0 {
			handleObjectRequest(respWriter, req.Method, bucket, object)
		} else {
			handleBucketRequest(respWriter, req.Method, bucket)
		}
	} else {

	}
}

func handleBucketRequest(respWriter http.ResponseWriter, method string, bucket string) {

}

func handleObjectRequest(respWriter http.ResponseWriter, method string, bucket string, object string) {

}

func main() {
	http.HandleFunc("/", gcsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
