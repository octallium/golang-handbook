package main

import "fmt"

// BUCKETS defines the total number of buckets in the table
const BUCKETS = 15

// Node defines the value and pointer to the next item
// in the bucket
type Node struct {
	Value int
	Next  *Node
}

// HashTable defines the structure of the hash tables, it
// gives access to buckets and restricts the buckets to
// the predetermined size
type HashTable struct {
	Table map[int]*Node
	Size  int
}

// determine the position where the value has to be inserted
func position(i, size int) int {
	return i % size
}

// insert a value in the hash table
func insert(hash *HashTable, value int) int {
	index := position(value, hash.Size)
	element := Node{
		Value: value,
		Next:  hash.Table[index],
	}
	hash.Table[index] = &element
	return index
}

// traverse the hash table and print out the values
func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d ->", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

// check for a value in the hash table
func lookup(hash *HashTable, value int) bool {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				if t.Value == value {
					return true
				}
				t = t.Next
			}
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, BUCKETS)
	hashTable := &HashTable{
		Table: table,
		Size:  BUCKETS,
	}
	fmt.Println("Total no of buckets =", hashTable.Size)
	for i := 0; i < 120; i++ {
		insert(hashTable, i)
	}
	traverse(hashTable)
	checkFor := 120
	if !lookup(hashTable, checkFor) {
		fmt.Printf("%d is not in hashtable!\n", checkFor)
	} else {
		fmt.Printf("%d found in hashtable!\n", checkFor)
	}
}
