//
// go tool go2go test
//

package hashmap

type H64[T any] func(T) uint64
type Equal[T any] func(T, T) bool
type Create[T any] func() T

type Node_t[Key_t any, Value_t any] struct {
	Key   Key_t
	Value Value_t
}

type Hashmap_t[Key_t any, Value_t any] struct {
	hash_func       H64[Key_t]
	equal_func      Equal[Key_t]
	hash_table      [][]*Node_t[Key_t, Value_t]
	load_factor_num int
	load_factor_den int
	table_size      int
	count           int
}

func New[Key_t any, Value_t any](hash_func H64[Key_t], equal_func Equal[Key_t], table_size int, load_factor_num int, load_factor_den int) (self *Hashmap_t[Key_t, Value_t]) {
	if table_size < 8 {
		table_size = 8
	}
	if load_factor_num <= 0 {
		load_factor_num = 1
	}
	if load_factor_den <= 0 {
		load_factor_den = 1
	}
	self = &Hashmap_t[Key_t, Value_t]{
		hash_func:       hash_func,
		equal_func:      equal_func,
		hash_table:      make([][]*Node_t[Key_t, Value_t], table_size),
		table_size:      table_size,
		load_factor_num: load_factor_num,
		load_factor_den: load_factor_den,
	}
	return
}

func (self *Hashmap_t[Key_t, Value_t]) rehash(new_len uint64) uint64 {
	hash_table := make([][]*Node_t[Key_t, Value_t], new_len)
	for _, bucket := range self.hash_table {
		for _, node := range bucket {
			ix := self.hash_func(node.Key) % new_len
			hash_table[ix] = append(hash_table[ix], node)
		}
	}
	self.hash_table = hash_table
	return new_len
}

func (self *Hashmap_t[Key_t, Value_t]) Insert(key Key_t, value Create[Value_t]) (node *Node_t[Key_t, Value_t], ok bool) {
	id := self.hash_func(key)
	ix := id % uint64(len(self.hash_table))
	for _, node = range self.hash_table[ix] {
		if self.equal_func(node.Key, key) {
			return
		}
	}
	if temp := len(self.hash_table); self.count > temp*self.load_factor_num/self.load_factor_den {
		ix = id % self.rehash(uint64(temp)*4/3)
	}
	node = &Node_t[Key_t, Value_t]{
		Key:   key,
		Value: value(),
	}
	self.hash_table[ix] = append(self.hash_table[ix], node)
	self.count++
	ok = true
	return
}

func (self *Hashmap_t[Key_t, Value_t]) Delete(key Key_t) (node *Node_t[Key_t, Value_t], ok bool) {
	var i int
	ix := self.hash_func(key) % uint64(len(self.hash_table))
	for i, node = range self.hash_table[ix] {
		if self.equal_func(node.Key, key) {
			temp := len(self.hash_table[ix])
			self.hash_table[ix][temp-1], self.hash_table[ix][i] = self.hash_table[ix][i], self.hash_table[ix][temp-1]
			self.hash_table[ix] = self.hash_table[ix][:temp-1]
			if temp = len(self.hash_table); temp > self.table_size && self.count*self.load_factor_num/self.load_factor_den < temp {
				self.rehash(uint64(temp - 1))
			}
			self.count--
			ok = true
			return
		}
	}
	return
}

func (self *Hashmap_t[Key_t, Value_t]) Find(key Key_t) (*Node_t[Key_t, Value_t], bool) {
	for _, node := range self.hash_table[self.hash_func(key)%uint64(len(self.hash_table))] {
		if self.equal_func(node.Key, key) {
			return node, true
		}
	}
	return nil, false
}

func (self *Hashmap_t[Key_t, Value_t]) Size() int {
	return self.count
}

func (self *Hashmap_t[Key_t, Value_t]) Buckets() int {
	return len(self.hash_table)
}