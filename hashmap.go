//
//
//

package hashmap

type HashUint64[T comparable] func(T) uint64
type CreateValue[T any] func() T

type Node_t[Key_t comparable, Value_t any] struct {
	Key   Key_t
	Value Value_t
}

type Hashmap_t[Key_t comparable, Value_t any] struct {
	hash_table      [][]*Node_t[Key_t, Value_t]
	hash_func       HashUint64[Key_t]
	load_factor_num int
	load_factor_den int
	count           int
}

func New[Key_t comparable, Value_t any](hash_func HashUint64[Key_t], table_size int, load_factor_num int, load_factor_den int) (self *Hashmap_t[Key_t, Value_t]) {
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
		hash_table:      make([][]*Node_t[Key_t, Value_t], table_size),
		load_factor_num: load_factor_num,
		load_factor_den: load_factor_den,
	}
	return
}

func (self *Hashmap_t[Key_t, Value_t]) rehash(new_len uint64) uint64 {
	hash_table := make([][]*Node_t[Key_t, Value_t], new_len)
	for _, nodes := range self.hash_table {
		for _, node := range nodes {
			bucket := self.hash_func(node.Key) % new_len
			hash_table[bucket] = append(hash_table[bucket], node)
		}
	}
	self.hash_table = hash_table
	return new_len
}

func (self *Hashmap_t[Key_t, Value_t]) Insert(key Key_t, value CreateValue[Value_t]) (node *Node_t[Key_t, Value_t], ok bool) {
	id := self.hash_func(key)
	bucket := id % uint64(len(self.hash_table))
	for _, node = range self.hash_table[bucket] {
		if node.Key == key {
			return
		}
	}
	if buckets := len(self.hash_table); self.count > buckets*self.load_factor_num/self.load_factor_den {
		bucket = id % self.rehash(uint64(buckets)*4/3)
	}
	node = &Node_t[Key_t, Value_t]{
		Key:   key,
		Value: value(),
	}
	self.hash_table[bucket] = append(self.hash_table[bucket], node)
	self.count++
	ok = true
	return
}

func (self *Hashmap_t[Key_t, Value_t]) Delete(key Key_t) (node *Node_t[Key_t, Value_t], ok bool) {
	var i int
	bucket := self.hash_func(key) % uint64(len(self.hash_table))
	for i, node = range self.hash_table[bucket] {
		if node.Key == key {
			// swap, resize, rehash
			temp := len(self.hash_table[bucket])
			self.hash_table[bucket][temp-1], self.hash_table[bucket][i] = self.hash_table[bucket][i], self.hash_table[bucket][temp-1]
			self.hash_table[bucket] = self.hash_table[bucket][:temp-1]
			self.count--
			if temp = len(self.hash_table); temp > self.count*self.load_factor_num/self.load_factor_den {
				self.rehash(uint64(temp - 1))
			}
			ok = true
			return
		}
	}
	return
}

func (self *Hashmap_t[Key_t, Value_t]) Find(key Key_t) (node *Node_t[Key_t, Value_t], ok bool) {
	for _, node = range self.hash_table[self.hash_func(key)%uint64(len(self.hash_table))] {
		if node.Key == key {
			ok = true
			return
		}
	}
	return
}

func (self *Hashmap_t[Key_t, Value_t]) Size() int {
	return self.count
}

func (self *Hashmap_t[Key_t, Value_t]) Buckets() int {
	return len(self.hash_table)
}
