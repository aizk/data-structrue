package data_structrue

// container interface is all data structures need.
type Container interface {
	Empty() bool // judge container is empty
	Size() int // get internal size
	Clear() // clear all key
	Values() []interface{} // get all values
}
