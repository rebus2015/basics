package structs

import (
	mh "github.com/dolthub/maphash"
)

type hashTable struct {
	buckets []linkedList
	hasher  mh.Hasher[string]
	base    int
	len     int
}

type data struct {
	key string
	val any
}

func NewHashTable(n int) hashTable {
	return hashTable{
		buckets: make([]linkedList, n),
		hasher:  mh.NewHasher[string](),
		base:    n,
	}
}

func (h *hashTable) index(hashValue uint64) uint64 {
	return (hashValue % uint64(h.base))
}

func (h *hashTable) Add(key string, val any) {
	k := h.index(h.hasher.Hash(key))
	if h.buckets[k].Head != nil {
		current := h.buckets[k].Head
		for {
			if current.Data.(data).key == key {
				current.Data = data{key, val}
				return
			}
			if current.Next == nil {
				break
			}
			current = current.Next
		}
	} else {
		h.buckets[k] = NewLinkedList()
	}

	h.buckets[k].AddToTail(data{key, val})
	h.len++
}

func (h *hashTable) Get(key string) (any, bool) {
	k := h.index(h.hasher.Hash(key))
	if h.buckets[k].len == 0 {
		return nil, false
	}
	current := h.buckets[k].Head
	for {
		if current.Data.(data).key == key {
			return current.Data.(data).val, true
		}
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	return nil, false
}

func (h *hashTable) Delete(key string) bool {
	k := h.index(h.hasher.Hash(key))
	if h.buckets[k].len == 0 {
		return false
	}
	val, ok := h.Get(key)
	if !ok {
		return false
	}

	err := h.buckets[k].Remove(data{key, val})
	return err == nil
}

func (h *hashTable) Print() {
	for _, v := range h.buckets {
		if v.len != 0 {
			v.Traverse()
		}
	}
}
