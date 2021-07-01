package hashset

type MyHashSet struct {
	Map map[int]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Map: make(map[int]bool),
	}
}

func (hashSet *MyHashSet) Add(key int) {
	if !hashSet.Contains(key) {
		hashSet.Map[key] = true
	}
}

func (hashSet *MyHashSet) Remove(key int) {
	if hashSet.Contains(key) {
		delete(hashSet.Map, key)
	}
}

/** Returns true if hashSet set contains the specified element */
func (hashSet *MyHashSet) Contains(key int) bool {
	return hashSet.Map[key]
}
