//
//
//

package hashmap

import (
	"fmt"
	"testing"
	// "gotest.tools/assert"
)

func int_hash(a int) uint64 {
	return uint64(a)
}

func Test01(t *testing.T) {
	h := New[int, string](int_hash, 0, 0, 0)

	node, ok := h.Insert(1, func() string { return "lalala" })
	fmt.Printf("INSERT: %v %v\n", node, ok)

	node, ok = h.Insert(1, func() string { return "lalala" })
	fmt.Printf("INSERT: %v %v\n", node, ok)

	node, ok = h.Insert(2, func() string { return "bububu" })
	fmt.Printf("INSERT: %v %v\n", node, ok)

	node, ok = h.Delete(1)
	fmt.Printf("DELETE: %v %v\n", node, ok)

	node, ok = h.Insert(1, func() string { return "lalala" })
	fmt.Printf("INSERT: %v %v\n", node, ok)

	fmt.Printf("REHASH1: %v %v\n", h.Size(), h.Buckets())
	h.rehash(8 * 4 / 3)
	fmt.Printf("REHASH2: %v %v\n", h.Size(), h.Buckets())

	node, ok = h.Find(1)
	fmt.Printf("FIND: %v %v\n", node, ok)

	node, ok = h.Find(2)
	fmt.Printf("FIND: %v %v\n", node, ok)

	node, ok = h.Find(3)
	fmt.Printf("FIND: %v %v\n", node, ok)
}
