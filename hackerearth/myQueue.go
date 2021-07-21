package main

import "fmt"

type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	(*this).queue = append((*this).queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	res := this.Peek()
	(*this).queue = (*this).queue[1:]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return (*this).queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len((*this).queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	param_3 := obj.Empty()
	obj.Push(2)
	param_4 := obj.Empty()

	fmt.Println(param_2, param_3, param_4)
}
