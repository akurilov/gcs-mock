package internal

import (
	"encoding/json"
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

func handleBucketRequest(storage *Storage, resp http.ResponseWriter, req *http.Request, bucket string) {
	switch req.Method {
	case "DELETE":
		err := storage.DeleteBucket(bucket)
		if err == nil {
			resp.WriteHeader(http.StatusNoContent)
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
	case "GET":
		if len(bucket) > 0 {
			err := storage.ReadBucket(bucket)
			if err == nil {
				resp.WriteHeader(http.StatusOK)
				resp.Header().Set("Content-Type", "application/json")
				encoder := json.NewEncoder(resp)
				encoder.Encode(NewBucketResource(bucket))
			} else {
				resp.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			query := req.URL.Query()
			maxResults, err := strconv.Atoi(query.Get("maxResults"))
			if err != nil || maxResults < 1 {
				maxResults = 1000
			}
			pageToken := query.Get("pageToken")
			prefix := query.Get("prefix")
			buckets, err := storage.ListBuckets(maxResults, pageToken, prefix)
			if err == nil {
				resp.Header().Set("Content-Type", "application/json")
				encoder := json.NewEncoder(resp)
				encoder.Encode(NewBucketListResource(buckets))
			} else {
				resp.WriteHeader(http.StatusInternalServerError)
			}
		}
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var res bucketResource
		err := decoder.Decode(&res)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
		} else {
			bucket := res.Name
			err := storage.CreateBucket(bucket)
			if err == nil {
				resp.WriteHeader(http.StatusCreated)
			} else {
				resp.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

func handleObjectRequest(
	storage *Storage, respWriter http.ResponseWriter, req *http.Request, bucket string, object string) {
	// TODO
}
