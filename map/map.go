package _map

type Map interface {
	Get(key interface{}) (interface{}, bool)
	Put(key, value interface{})
	Remove(key interface{})
	Keys() []interface{}
	Values() []interface{}
	Empty() bool
	Size() int
	Clear() error
}
