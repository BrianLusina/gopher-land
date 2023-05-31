package rpn

import (
	"gopherland/datastructures/stack"
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := stack.NewStack[int](len(tokens))

	for _, char := range tokens {
		switch char {
		case "+":
			a, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			b, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			res := a + b
			stack.Push(res)
		case "-":
			a, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			b, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			res := b - a
			stack.Push(res)
		case "*":
			a, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			b, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			res := a * b
			stack.Push(res)
		case "/":
			a, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			b, err := stack.Pop()
			if err != nil {
				log.Fatalf("Nothing to pop: %s", err)
			}
			res := b / a
			stack.Push(res)
		default:
			if num, err := strconv.Atoi(char); err == nil {
				stack.Push(num)
			} else {
				log.Fatalf("invalid number supplied: %s", char)
			}
		}
	}

	if value, err := stack.Peek(); err != nil {
		log.Fatalf("Failed to evaluate RPN")
		return 0
	} else {
		return value
	}
}

func evalRPNWithSlice(tokens []string) int {
	stack := []int{}

	for _, char := range tokens {
		switch char {
		case "+":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := a + b
			stack = append(stack, res)
		case "-":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := b - a
			stack = append(stack, res)
		case "*":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := a * b
			stack = append(stack, res)
		case "/":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := b / a
			stack = append(stack, res)
		default:
			if num, err := strconv.Atoi(char); err == nil {
				stack = append(stack, num)
			} else {
				log.Fatalf("invalid number supplied: %s", char)
			}
		}
	}

	return stack[0]
}
