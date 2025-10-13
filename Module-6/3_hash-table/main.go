package main

import "fmt"

type HashTable struct {
	table []*Entry
	size  int // this is m
	count int // this is n
}

type Entry struct {
	key     int
	value   int
	deleted bool
}

func main() {
	testTable := NewHashTable(13)
	err := testTable.Insert(124, 21)
	if err != nil {
		panic(err)
	}

	err = testTable.Insert(1, 1212)
	if err != nil {
		panic(err)
	}

	value, ok := testTable.Get(124)
	fmt.Println(value, ok)
}

func NewHashTable(m int) HashTable {
	return HashTable{
		table: make([]*Entry, m),
		size:  m,
		count: 0,
	}
}

// Insert creates a new entry in the hashtable with the specified key and value
// If the key already exists, it overwrites its value
func (ht *HashTable) Insert(key, value int) error {
	if ht.count >= ht.size {
		return fmt.Errorf("unable to insert value, no free space in hashtable")
	}

	for i := 0; i < ht.size; i++ {
		hashValue := ht.hashFunction(key, i, ht.size)
		entry := ht.table[hashValue]

		if entry == nil || entry.deleted {
			ht.table[hashValue] = &Entry{key, value, false}
			ht.count++
			return nil
		}

		if entry.key == key {
			entry.value = value
			return nil
		}
	}

	return fmt.Errorf("unable to insert value at %d for some reason", key)
}

// Get gets the value for an entry based on the specified key if it exists.
func (ht *HashTable) Get(key int) (value int, exists bool) {
	for i := 0; i < ht.size; i++ {
		hashValue := ht.hashFunction(key, i, ht.size)
		entry := ht.table[hashValue]

		if entry == nil {
			break
		}

		if entry.key == key && !entry.deleted {
			return entry.value, true
		}
	}

	return 0, false
}

func (ht *HashTable) Delete(key int) error {
	for i := 0; i < ht.size; i++ {
		hashValue := ht.hashFunction(key, i, ht.size)
		entry := ht.table[hashValue]

		if entry == nil {
			break
		}

		if entry.key == key {
			entry.deleted = true
			ht.count--
			return nil
		}
	}

	return fmt.Errorf("unable to find entry with key: %d", key)
}

func (ht *HashTable) hashFunction(k, i, m int) int {
	return (k + i) % m
}
