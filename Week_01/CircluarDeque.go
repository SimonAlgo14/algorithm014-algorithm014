package main

/*

https://leetcode.com/problems/design-circular-queue/submissions/

["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]
[[3],          [1],[2],[3],[4],[],[],[],[4],[]]

Ans: [null,true,true,true,false,3,true,true,true,4]

["MyCircularQueue","enQueue","Rear","Front","deQueue","Front","deQueue","Front","enQueue","enQueue","enQueue","enQueue"]
[[3],          [2],[],[],[],[],[],[],[4],[2],[2],[3]]

Ans: [null,true,2,2,true,-1,false,-1,true,true,true,false]
*/

type MyCircularQueue struct {
	front, rear, cnt, cap int
	array            []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		array: make([]int, k),
		front: 0,
		rear:  0,
		cnt :  0,
		cap:   k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.array[this.front] = value
	this.front = (this.front + 1) % this.cap
	this.cnt ++

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.rear = (this.rear + 1) % this.cap
	this.cnt --
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.array[this.rear] // *** rear ****
	}
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.array[ (this.front - 1  + this.cap) % this.cap ] // *** front - 1 ****
	}
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.cnt == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.cnt == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
