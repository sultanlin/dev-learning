package datastructs

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*Bucket
}

type Bucket struct {
	Head *BucketNode
}

type BucketNode struct {
	Key  int
	Next *BucketNode
}

// func Init() *HashTable {
// 	result := &HashTable{}
// 	for i:= range result {
// 		result.array[i] = &Bucket{}
// 	}
// }

func NewHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func (ht *HashTable) Insert(key string) {
	// index := Hash(key)
	// ht.array.insertInBucket(key)
}

func (ht *HashTable) Search(key string) bool {
	// index := Hash(key)
	return true
}

func (ht *HashTable) Delete(key string) {
	// index := Hash(key)

}

// func (b *Bucket) insertInBucket(key string) {
// 	newNode := &BucketNode{key: key}
// 	newNode.Next = b.Head
//
// }

func Hash(key string) int {
	sum := 0
	for _, letter := range key {
		sum += int(letter)
	}
	return sum % ArraySize
}
