package singly

type List struct {
	first *node
	last *node
}

type node struct {
	value interface{}
	next *node
}

