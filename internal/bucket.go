package internal

import (
	"time"
)

type bucketResource struct {
	Resource
}

func NewBucketResource(id string) *bucketResource {
	return &bucketResource{
		Resource{
			Etag: "",
			Id:   id,
			Kind: "storage#bucket",
			Name: id,
			Owner: Owner{
				Entity:   "",
				EntityId: "",
			},
			SelfLink:    "",
			TimeCreated: time.Now(),
			Updated:     time.Now(),
		},
	}
}

func DeleteBucket(bucket string) error {
	return nil
}

func ReadBucket(bucket string) (*bucketResource, error) {
	return nil, nil
}

func ListBuckets(pageToken string) ([...]*bucketResource, error) {

}

func CreateBucket() (*bucketResource, error) {

}
