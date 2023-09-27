/**
 * @author: baowj
 * @time: 2023/9/26 17:10
 */
package main

import (
	"container/heap"
)

type LFUCache struct {
	cache     []int
	frequency []int
	timestamp []int
	queue     []keyTime
	clock     int
	capacity  int
	num       int
}

type keyTime struct {
	key       int
	frequency int
	timestamp int
}

func Constructor(capacity int) LFUCache {
	cache := make([]int, 1e5+100)
	frequency := make([]int, 1e5+100)
	timestamp := make([]int, 1e5+100)
	queue := []keyTime{}
	for i := range cache {
		cache[i] = -1
	}

	lfu := LFUCache{
		cache:     cache,
		frequency: frequency,
		timestamp: timestamp,
		queue:     queue,
		clock:     0,
		capacity:  capacity,
		num:       0,
	}
	return lfu
}

func (this *LFUCache) Len() int {
	return len(this.queue)
}

func (this *LFUCache) Pop() any {
	res := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return res
}

func (this *LFUCache) Push(x any) {
	this.queue = append(this.queue, x.(keyTime))
}

func (this *LFUCache) Less(i, j int) bool {
	if this.queue[i].frequency == this.queue[j].frequency {
		return this.queue[i].timestamp < this.queue[j].timestamp
	} else {
		return this.queue[i].frequency < this.queue[j].frequency
	}
}

func (this *LFUCache) Swap(i, j int) {
	this.queue[i], this.queue[j] = this.queue[j], this.queue[i]
}

func (this *LFUCache) Get(key int) int {
	this.clock++
	if v := this.cache[key]; v != -1 {
		this.frequency[key]++
		this.timestamp[key] = this.clock
		heap.Push(this, keyTime{key: key, frequency: this.frequency[key], timestamp: this.timestamp[key]})
		return v
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cache[key] == -1 {
		if this.num >= this.capacity {
			for {
				popKey := heap.Pop(this).(keyTime)
				if popKey.timestamp == this.timestamp[popKey.key] {
					this.cache[popKey.key] = -1
					this.frequency[popKey.key] = 0
					this.timestamp[popKey.key] = 0
					this.num--
					break
				}
			}
		}
		this.num++
	}
	this.clock++
	this.frequency[key]++
	this.cache[key] = value
	this.timestamp[key] = this.clock
	heap.Push(this, keyTime{key: key, frequency: this.frequency[key], timestamp: this.timestamp[key]})
}
