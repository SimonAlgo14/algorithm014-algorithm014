package main

import (
	"container/heap"
	"fmt"
)

type IntFrequent struct {
	Val int
	Frq int
}
type Heap []IntFrequent


func (h *Heap) Less(i, j int) bool {
	return (*h)[j].Frq < (*h)[i].Frq
}

func (h *Heap) Swap(i, j int)  {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Pop() (v interface{}) {
	v   = (*h)[h.Len()-1]
	*h  = (*h)[:h.Len()-1]
	return v
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(IntFrequent))
}

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	h := new(Heap)

	for _, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num] = hash[num] + 1
		} else {
			hash[num] = 1
		}
	}

	for key, val := range hash {
		h.Push(IntFrequent{Val: key, Frq: val})
	}
	heap.Init(h)

	var r []int
	for i := 0; i < k; i++ {
		r = append(r, heap.Pop(h).(IntFrequent).Val)
	}

	return r
}

func main() {

	a := []int{ 1,1,1,2,2,9,9,9,9,7 }
	b := topKFrequent(a, 2)

	fmt.Println(a)
	fmt.Println(b)
}
