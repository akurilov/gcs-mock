package internal

import "time"

type Owner struct {
	Entity   string
	EntityId string
}

type Resource struct {
	Etag        string
	Id          string
	Kind        string
	Name        string
	Owner       Owner
	SelfLink    string
	TimeCreated time.Time
	Updated     time.Time
}

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
