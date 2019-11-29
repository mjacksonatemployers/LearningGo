package main

import "fmt"

type Stack struct {

	stackList []int
}

func NewStack() *Stack {
	return &Stack{
		stackList: make([]int,0,10),
	}

}

func (s *Stack) isEmpty() bool {
	if cap(s.stackList) == 0 {
		return true
	}
	return false
}

func (s *Stack) pop() {
	if len(s.stackList) > 0 {
		s.stackList = s.stackList[:len(s.stackList) - 1]
	}
	fmt.Println("Stack Popped")
}

func (s *Stack) push(value int) {
	if cap(s.stackList) == len(s.stackList) {
		s.stackList = make([]int, 10)
	}
	s.stackList = append(s.stackList, value)
	fmt.Println("Stack Pushed")
}

func (s *Stack) count() {
	fmt.Println("Count is:", len(s.stackList))
}

func (s *Stack) print() {
	fmt.Println("Print:" ,s.stackList)
}
