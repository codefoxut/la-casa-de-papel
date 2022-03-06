package main

import (
	"fmt"
)

type LRUCache struct {
	capacity int
	data     map[int]int
	keys     []int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		data:     make(map[int]int, capacity),
		keys:     make([]int, 0, capacity),
	}
	return lru
}

func (l *LRUCache) updateQueue(key int) (bool, int) {
	var evicted bool
	var evictionKey int
	keyIndex := -1
	for index, item := range l.keys {
		if item == key {
			keyIndex = index
			break
		}
	}

	//  key is already in the queue
	if keyIndex > -1 {
		l.keys = append(l.keys[:keyIndex], l.keys[keyIndex+1:]...)
	} else {
		//  key is not there in the queue
		if len(l.keys) == l.capacity {
			evictionKey = l.keys[0]
			evicted = true
			l.keys = l.keys[1:]
		}
	}
	l.keys = append(l.keys, key)

	fmt.Println(l.keys)
	return evicted, evictionKey

}

func (l *LRUCache) Get(key int) int {
	if l != nil {
		if v, ok := l.data[key]; ok {
			l.updateQueue(key)
			return v
		}
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if l != nil {
		l.data[key] = value
		eviction, evictionKey := l.updateQueue(key)
		if eviction {
			fmt.Println("evicts ", evictionKey)
			delete(l.data, evictionKey)
		}
	}

}

func main() {
	/**
	* Your LRUCache object will be instantiated and called as such:
	* obj := Constructor(capacity);
	* param_1 := obj.Get(key);
	* obj.Put(key,value);
	 */
	capacity := 2
	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	param1 := obj.Get(1)
	fmt.Printf("getter 1 %d \n", param1)
	obj.Put(3, 3)
	param2 := obj.Get(2)
	fmt.Printf("getter 2 %d \n", param2)
	obj.Put(4, 4)
	param := obj.Get(1)
	fmt.Printf("getter 3 %d \n", param)
	param = obj.Get(3)
	fmt.Printf("getter 4 %d \n", param)
	param = obj.Get(4)
	fmt.Printf("getter 5 %d \n", param)

}
