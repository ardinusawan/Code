// https://leetcode.com/problems/find-median-from-data-stream/
// Heap inspired: https://www.tugberkugurlu.com/archive/usage-of-the-heap-data-structure-in-go-golang-with-examples

package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

// Less always make sure left < right
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	max := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return max
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

// Less always make sure left > right
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	max := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return max
}

type MedianFinder struct {
	small MinHeap
	large MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{small: MinHeap{}, large: MaxHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	if this.large.Len() == 0 || num <= this.large[0] {
		heap.Push(&this.large, num)
	} else {
		heap.Push(&this.small, num)
	}

	// uneven
	if this.small.Len() > this.large.Len() {
		val := heap.Pop(&this.small)
		heap.Push(&this.large, val)
	}
	if this.large.Len() > this.small.Len()+1 {
		val := heap.Pop(&this.large)
		heap.Push(&this.small, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() == this.large.Len() {
		return float64(this.small[0]+this.large[0]) / 2
	}
	return float64(this.large[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	t := Constructor()
	t.AddNum(10)
	t.AddNum(2)
	t.AddNum(8)
	t.AddNum(3)
	fmt.Println(t.FindMedian())
	fmt.Println(t.small.Pop())
	fmt.Println(t.small.Pop())
	fmt.Println(t.large.Pop())
	fmt.Println(t.large.Pop())
}
