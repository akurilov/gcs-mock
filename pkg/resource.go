package pkg

import "time"

type ResourceBase struct {
	Kind string `json:"kind"`
}

type Owner struct {
	Entity   string `json:"entity"`
	EntityId string `json:"entityId"`
}

type NamedResource struct {
	ResourceBase
	Etag        string    `json:"etag"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Owner       Owner     `json:"owner"`
	SelfLink    string    `json:"selfLink"`
	TimeCreated time.Time `json:"timeCreated"`
	Updated     time.Time `json:"updated"`
}

type bucketResource struct {
	NamedResource
}

func NewBucketResource(id string) *bucketResource {
	return &bucketResource{
		NamedResource{
			ResourceBase{
				Kind: "storage#bucket",
			},
			"",
			id,
			id,
			Owner{
				Entity:   "",
				EntityId: "",
			},
			"",
			time.Now(),
			time.Now(),
		},
	}
}

type bucketListResource struct {
	ResourceBase
	NextPageToken string            `json:"nextPageToken"`
	Items         []*bucketResource `json:"items"`
}

func NewBucketListResource(buckets []string) *bucketListResource {
	var bucketResources []*bucketResource
	bucketResources = make([]*bucketResource, 0, len(buckets))
	var lastBucket string
	for _, bucket := range buckets {
		bucketResources = append(bucketResources, NewBucketResource(bucket))
		lastBucket = bucket
	}
	return &bucketListResource{
		ResourceBase{
			Kind: "storage#buckets",
		},
		lastBucket,
		bucketResources,
	}
}

type objectResource struct {
	NamedResource
	Bucket string `json:"bucket"`
	Size   uint64 `json:"size"`
}

func NewObjectResource(bucket string, id string, size uint64) *objectResource {
	return &objectResource{
		NamedResource{
			ResourceBase{
				Kind: "storage#object",
			},
			"",
			id,
			id,
			Owner{
				Entity:   "",
				EntityId: "",
			},
			"",
			time.Now(),
			time.Now(),
		},
		bucket,
		size,
	}
}
