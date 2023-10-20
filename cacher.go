package caches

type Cacher interface {
	Get(key string, model any) *Query
	Store(key string, val *Query) error
}
