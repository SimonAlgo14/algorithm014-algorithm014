package main

import (
	"fmt"
)

func main() {
	s := "()[]{}"
	fmt.Printf("%s %v\n", s, isValid(s))
}

// https://leetcode.com/problems/valid-parentheses/submissions/

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	index := -1
	last := ' '

	for _, c := range s {
		switch c {
		case '(':
			stack, index = push(stack, ')', index)
		case '[':
			stack, index = push(stack, ']', index)
		case '{':
			stack, index = push(stack, '}', index)
		default:
			if index < 0 {
				return false
			}

			if stack, index, last = pop(stack, index); c != last {
				return false
			}
		}
	}

	return index < 0
}

func push(stack []rune, c rune, index int) ([]rune, int) {
	stack = append(stack, c)
	index++
	return stack, index
}

func pop(stack []rune, index int) ([]rune, int, rune) {
	last := stack[index]
	stack = stack[0:index]
	index--
	return stack, index, last
}
