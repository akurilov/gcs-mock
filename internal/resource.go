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
