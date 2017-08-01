package radix

import (
	"errors"
	"sync"
)

var (
	ErrNodeNotFound = errors.New("can not find node")
	ErrNodeLabel    = errors.New("node label required")
	ErrNodeValue    = errors.New("node value required")
)

type (
	Tree struct {
		root  *Node
		rwMux *sync.RWMutex
	}

	Node struct {
		label rune
		next  []*Node

		val interface{}
	}
)

func (t *Tree) Get(s string) (interface{}, error) {
	t.rwMux.RLock()
	defer t.rwMux.RUnlock()

	r, lr := stringToRune(s)
	if nil == t.root || t.root.label != r[0] {
		return nil, ErrNodeNotFound
	}

	return t.lookup(r, t.root, 1, lr)
}

// lookup method performs the actual search in the tree
func (t *Tree) lookup(r []rune, prev *Node, i, l int) (interface{}, error) {
	if i == l-1 {
		if nil == prev.val {
			return nil, ErrNodeNotFound
		}

		return prev.val, nil
	}

	n, err := prev.withLabel(r[i])
	if nil != err {
		return nil, ErrNodeNotFound
	}

	return t.lookup(r, n, i+1, l)
}

// Appends a string into the tree and attaches the value to the last node
func (t *Tree) Add(s string, val interface{}) error {
	if "" == s {
		return ErrNodeLabel
	}
	if nil == val {
		return ErrNodeValue
	}

	t.rwMux.Lock()
	defer t.rwMux.Unlock()

	r, lr := stringToRune(s)

	if nil == t.root {
		t.root = newNode(r[0], nil)
	}

	return t.process(r, val, t.root, 1, lr)
}

// process processes the given string
func (t *Tree) process(r []rune, v interface{}, prev *Node, i, l int) error {
	if i == l-1 {
		prev.val = v
		return nil
	}

	n, err := prev.withLabel(r[i])
	if nil != err {
		n = newNode(r[i], nil)
		prev.addNext(n)
	}

	return t.process(r, v, n, i+1, l)
}

// withLabel iterates through the next nodes,
func (n *Node) withLabel(r rune) (*Node, error) {
	// BenchmarkTree_Add-8   	  200000	      8522 ns/op	     248 B/op	       2 allocs/op
	// BenchmarkTree_Get-8   	  200000	      8058 ns/op	     240 B/op	       1 allocs/op
	//
	for i := 0; i < len(n.next); i++ {
		if r == n.next[i].label {
			return n.next[i], nil
		}
	}

	return nil, ErrNodeNotFound
}

// withLabel iterates through the next nodes,
func (n *Node) addNext(i *Node) {
	n.next = append(n.next, i)
}

func newNode(l rune, v interface{}) *Node {
	return &Node{
		label: l,
		val:   v,
	}
}

// NewTree retru
func NewTree() *Tree {
	return &Tree{
		rwMux: &sync.RWMutex{},
	}
}

func stringToRune(s string) ([]rune, int) {
	r := []rune(s)
	lr := len(r)

	return r, lr
}
