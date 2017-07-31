package radix

import (
	"errors"
	"sync"
)

var (
	ErrNextNodeNotFound = errors.New("can not find next node")
	ErrNodeLabel        = errors.New("node label required")
	ErrNodeValue        = errors.New("node value required")
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

	r := []rune(s)

	if nil == t.root {
		t.root = newNode(r[0], nil)
	}

	t.process(r, val, t.root, 1, len(r))
	return nil
}

// process processes the given string
func (t *Tree) process(r []rune, v interface{}, prev *Node, i, l int) error {
	if i == l-1 {
		prev.val = v
		return nil
	}

	n, err := prev.nextWithLabel(r[i])
	if nil != err {
		n = newNode(r[i], nil)
		prev.addNext(n)
	}

	return t.process(r, v, n, i+1, l)
}

// nextWithLabel iterates through the next nodes,
func (n *Node) nextWithLabel(r rune) (*Node, error) {
	for i := 0; i < len(n.next); i++ {
		if r == n.next[i].label {
			return n.next[i], nil
		}
	}

	return nil, ErrNextNodeNotFound
}

// nextWithLabel iterates through the next nodes,
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
