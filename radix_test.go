package radix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Add(t *testing.T) {
	tree := NewTree()
	tree.Add("am/here", 1)

	v, err := tree.Get("am/here")

	assert.Equal(t, 1, v.(int), "Found value is not as expected")
	assert.Equal(t, nil, err, "Error found while lookup")
}

func BenchmarkTree_Add(b *testing.B) {
	tree := NewTree()
	for i := 0; i < b.N; i++ {
		tree.Add("this/is/something/that/is/long/enough/to/mimic/an/endpoint", 100)
	}
}

func BenchmarkTree_Get(b *testing.B) {
	tree := NewTree()
	tree.Add("this/is/something/that/is/long/enough/to/mimic/an/endpoint", 100)
	for i := 0; i < b.N; i++ {
		tree.Get("this/is/something/that/is/long/enough/to/mimic/an/endpoint")
	}
}
