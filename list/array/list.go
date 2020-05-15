package array

type List struct {
    nodes []interface{}
    size  int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.PushBack(values)
	}
	return list
}

func (list *List) PushBack(values ...interface{}) {
	list.nodes = append(list.nodes, values)
	list.size += len(values)
}