package LC24

import "fmt"

func Test() {
	q := Constructor()
	q.Push(1)
	v2 := q.Pop()
	empty := q.Empty()
	fmt.Println(v2)
	fmt.Println(empty)
}

type MyQueue struct {
	s1    Stack
	s2    Stack
	front int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: Stack{},
		s2: Stack{},
	}
}

func (q *MyQueue) Push(x int) {
	q.s1.Push(x)
	if q.s1.Len() == 1 {
		q.front = x
	}
}

func (q *MyQueue) Peek() int {
	return q.front
}

func (q *MyQueue) Pop() int {
	for !q.s1.Empty() {
		q.s2.Push(q.s1.Pop())
	}
	v := q.s2.Pop()
	for !q.s2.Empty() {
		q.s1.Push(q.s2.Pop())
	}
	return v
}

func (q *MyQueue) Empty() bool {
	return q.s1.Empty() && q.s2.Empty()
}

type Stack []int

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}
