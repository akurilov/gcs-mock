package internal

import "time"

type ResourceBase struct {
	Kind string
}

type Owner struct {
	Entity   string
	EntityId string
}

type NamedResource struct {
	ResourceBase
	Etag        string
	Id          string
	Name        string
	Owner       Owner
	SelfLink    string
	TimeCreated time.Time
	Updated     time.Time
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
	NextPageToken string
	Items         []*bucketResource
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
	Bucket string
	Size   uint64
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
