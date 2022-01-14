//
//
//

package hashmap

import (
	"testing"

	"gotest.tools/assert"
)

func int_hash(a int) uint64 {
	return uint64(a)
}

func Test01(t *testing.T) {
	h := New[int, string](int_hash, 0, 0, 0)

	node, ok := h.Insert(1, func() string { return "lalala" })
	assert.Assert(t, ok == true)
	assert.Assert(t, node.Key == 1)

	node, ok = h.Insert(1, func() string { return "lalala" })
	assert.Assert(t, ok == false)
	assert.Assert(t, node.Key == 1)

	node, ok = h.Insert(2, func() string { return "bububu" })
	assert.Assert(t, ok == true)
	assert.Assert(t, node.Key == 2)

	node, ok = h.Delete(1)
	assert.Assert(t, ok == true)
	assert.Assert(t, node.Key == 1)

	node, ok = h.Insert(1, func() string { return "lalala" })
	assert.Assert(t, ok == true)
	assert.Assert(t, node.Key == 1)

	h.rehash(8 * 4 / 3)

	node, ok = h.Find(1)
	assert.Assert(t, ok == true)
	assert.Assert(t, node.Key == 1)

	node, ok = h.Find(2)
	assert.Assert(t, ok == true)
	assert.Assert(t, node.Key == 2)

	node, ok = h.Find(3)
	assert.Assert(t, ok == false)
	assert.Assert(t, node == nil)
}
