package internal

import (
	"time"
)

type objectResource struct {
	Resource
	Bucket string
	Size   uint64
}

func NewObjectResource(bucket string, id string, size uint64) *objectResource {
	return &objectResource{
		Resource{
			Etag: "",
			Id:   id,
			Kind: "storage#object",
			Name: id,
			Owner: Owner{
				Entity:   "",
				EntityId: "",
			},
			SelfLink:    "",
			TimeCreated: time.Now(),
			Updated:     time.Now(),
		},
		bucket,
		size,
	}
}
