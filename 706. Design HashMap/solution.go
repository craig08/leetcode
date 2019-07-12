import (
	"hash/crc32"
	"strconv"
)

type MyHashMap struct {
	Buckets []*Node
}

type Node struct {
	Key, Val int
	Next     *Node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	hash := MyHashMap{}
	hash.Buckets = make([]*Node, 10000)
	return hash
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	bucket := hash(key)
	if this.Buckets[bucket] == nil {
		this.Buckets[bucket] = &Node{Next: &Node{Key: key, Val: value}}
	} else if prev, ok := this.find(bucket, key); ok {
		prev.Next.Val = value
	} else {
		this.Buckets[bucket].Next = &Node{
			Key:  key,
			Val:  value,
			Next: this.Buckets[bucket].Next,
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if prev, ok := this.find(hash(key), key); ok {
		return prev.Next.Val
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	bucket := hash(key)
	if this.Buckets[bucket] == nil {
		return
	}
	if prev, ok := this.find(bucket, key); ok {
		prev.Next = prev.Next.Next
	}
	if this.Buckets[bucket].Next == nil {
		this.Buckets[bucket] = nil
	}
}

func hash(key int) int {
	return int(crc32.ChecksumIEEE([]byte(strconv.Itoa(key))) % 10000)
}

func (this *MyHashMap) find(bucket, key int) (*Node, bool) {
	if head := this.Buckets[bucket]; head == nil {
		return nil, false
	} else {
		for prev := this.Buckets[bucket]; prev.Next != nil; prev = prev.Next {
			if prev.Next.Key == key {
				return prev, true
			}
		}
	}
	return nil, false
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
