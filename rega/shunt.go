package main

import (
	"fmt"
)

func intoPostFix(infix string) string {
	specials := map[rune]int{"*":10, ".":9, "|":8}

	postFix, stack := []rune{}, []rune{}

	for _, r := range infix{

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

	//Answer: ab.c*.
	fmt.Println("Infix:    ", "a.(b|d).c*")
	fmt.Println("Postfix:  ", intoPostFix("a.(b|d).c*"))

	//Answer: ab.c*.
	fmt.Println("Infix:    ", "a.(b|b)+.c")
	fmt.Println("Postfix:  ", intoPostFix("a.(b|b)+.c"))
}
