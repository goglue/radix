package radix

import (
	"fmt"
	"testing"
)

func TestTree_Add(t *testing.T) {
	tree := NewTree()
	tree.Add("amhere", 1)

	println(fmt.Sprintf("%s", tree.root))
}
