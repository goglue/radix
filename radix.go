package radix

import (
	"errors"
	"sort"
	"sync"
)

var (
	ErrNodeNotFound  = errors.New("can not find node")
	ErrDuplicateNode = errors.New("duplicate node")
	ErrNodeLabel     = errors.New("node label required")
	ErrNodeValue     = errors.New("node value required")
)

type (
	Tree struct {
		root  *Node
		rwMux *sync.RWMutex
	}

	Node struct {
		label byte
		next  []*Node

		val interface{}
	}
)

func (t *Tree) Get(s string) (interface{}, error) {
	r, lr := stringToBytes(s)
	if nil == t.root || t.root.label != r[0] {
		return nil, ErrNodeNotFound
	}

	return t.lookup(r, t.root, 1, lr)
}

// lookup method performs the actual search in the tree
func (t *Tree) lookup(r []byte, prev *Node, i, l int) (interface{}, error) {
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

	r, lr := stringToBytes(s)

	if nil == t.root {
		t.root = newNode(r[0], nil)
	}

	return t.process(r, val, t.root, 1, lr)
}

// process processes the given string
func (t *Tree) process(r []byte, v interface{}, prev *Node, i, l int) error {
	if i == l-1 {
		if nil != prev.val {
			return ErrDuplicateNode
		}
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
func (n *Node) withLabel(r byte) (*Node, error) {
	nl := len(n.next)
	id := sort.Search(nl, func(i int) bool {
		return n.next[i].label >= r
	})
	if id < nl && n.next[id].label == r {
		return n.next[id], nil
	}

	return nil, ErrNodeNotFound
}

func stringToBytes(s string) ([]byte, int) {
	r := []byte(s)
	lr := len(r)

	return r, lr
}

// withLabel iterates through the next nodes,
func (n *Node) addNext(i *Node) {
	n.next = append(n.next, i)
	sort.Slice(n.next, func(i, j int) bool { return n.next[i].label < n.next[j].label })
}

func newNode(l byte, v interface{}) *Node {
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
