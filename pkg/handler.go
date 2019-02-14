package pkg

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

const (
	uriPathPatternRaw = `/storage/v1/b(/([a-z\d\-_.]{3,222})(/o/([\d\w][\d\w\-_.]{0,1000}))?)?`
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
		bucket := result[2]
		object := result[4]
		if len(object) > 0 {
			handleObjectRequest(storage, respWriter, req, bucket, object)
		} else {
			handleBucketRequest(storage, respWriter, req, bucket)
		}
	}
}

func handleBucketRequest(storage *Storage, resp http.ResponseWriter, req *http.Request, bucket string) {
	switch req.Method {
	case http.MethodDelete:
		err := storage.DeleteBucket(bucket)
		if err == nil {
			resp.WriteHeader(http.StatusNoContent)
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
	case http.MethodGet:
		if len(bucket) > 0 {
			err := storage.ReadBucket(bucket)
			if err == nil {
				resp.WriteHeader(http.StatusOK)
				resp.Header().Set("Content-Type", "application/json")
				encoder := json.NewEncoder(resp)
				err := encoder.Encode(NewBucketResource(bucket))
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				if err == os.ErrNotExist {
					resp.WriteHeader(http.StatusNotFound)
				} else {
					resp.WriteHeader(http.StatusInternalServerError)
				}
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
				err := encoder.Encode(NewBucketListResource(buckets))
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				resp.WriteHeader(http.StatusInternalServerError)
			}
		}
	case http.MethodPost:
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
	default:
		resp.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleObjectRequest(storage *Storage, resp http.ResponseWriter, req *http.Request, bucket string, object string) {
	switch req.Method {
	case http.MethodDelete:
		break
	case http.MethodGet:
		break
	case http.MethodPost:
		break
	default:
		resp.WriteHeader(http.StatusMethodNotAllowed)
	}
}
