package hash

import "fmt"

type HashMap struct {
    m map[interface{}]interface{}
}

func NewHashMap() *HashMap {
	return &HashMap{
		m: make(map[interface{}]interface{}),
	}
}

func (h *HashMap) Get(key interface{}) (value interface{}, exist bool) {
	value, exist = h.m[key]
	return
}

func (h *HashMap) Put(key, value interface{}) {
	h.m[key] = value
}

func (h *HashMap) Remove(key interface{}) {
	delete(h.m, key)
}

// more effect
func (h *HashMap) Keys() (keys []interface{}) {
	for key := range h.m {
		keys = append(keys, key)
	}
	return
}

func (h *HashMap) Values() (values []interface{}) {
	for _, value := range h.m {
		values = append(values, value)
	}
	return
}

func (h *HashMap) Empty() bool {
	return len(h.m) == 0
}

func (h *HashMap) Size() int {
	return len(h.m)
}

func (h *HashMap) Clear() {
	h.m = make(map[interface{}]interface{})
}

func (h *HashMap) String() string {
	return fmt.Sprintf("%+v", h.m)
}