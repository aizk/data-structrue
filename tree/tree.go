package tree

import data_structrue "github.com/liunian1004/data-structrue"

type Tree interface {
	data_structrue.Container
	Add(key, value interface{})
	Get(key interface{}) (value interface{}, find bool)
	Remove(key interface{})
}
