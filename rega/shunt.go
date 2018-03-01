package main

import (
	"fmt"
)


func intoPostFix(infix string) string {
	//precedence map
	specials := map[rune]int{'*':10, '.':9, '|':8}
	//initialsing stacks
	postFix, stack := []rune{}, []rune{}
	//range loop converts string to char array (UTF-8)
	for _, r := range infix{
		switch{
		case r == '(':
			stack = append(stack, r)
		case r == ')':
			for stack[len(stack)-1] != '('{
				//next 2 lines pop character off stack and append it to returned postfix
				postFix = append(postFix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			//pops off next character after brackets
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			//if stack isn't empty and the precedence off current char is less/equals to the precedence of whats on the stack
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]]{
				//next 2 lines pop character off stack and append it to returned postfix
				postFix = append(postFix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		default: 
			postFix = append(postFix, r)
		}
	}
	for len(stack)>0{
		//next 2 lines pop character off stack and append it to returned postfix
		postFix = append(postFix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(postFix)
}

func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:    ", "a.b.c*")
	fmt.Println("Postfix:  ", intoPostFix("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix:    ", "(a.(b|d))*")
	fmt.Println("Postfix:  ", intoPostFix("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix:    ", "a.(b|d).c*")
	fmt.Println("Postfix:  ", intoPostFix("a.(b|d).c*"))

	//Answer: abb|+.c.
	fmt.Println("Infix:    ", "a.(b|b)+.c")
	fmt.Println("Postfix:  ", intoPostFix("a.(b|b)+.c"))
}
