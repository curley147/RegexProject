package main

import (
	"fmt"
)

//creating struct for each state, it contains 2 pointers
//to states(arrows) and a symbol variable for the label
//on the arrow
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
	//boolean to set states to reject states for ? char
	rej bool
}

//struct to represent smaller subsets of overall nfa
type nfaFrag struct {
	initial *state
	accept  *state
}

//function to convert infix regular expressions to postfix (e.g. infix regex "a.b.c*" = "ab.c*." in postfix)
func IntoPostFix(infix string) string {
	//precedence map
	specials := map[rune]int{'*': 10, '+': 9, '?': 8, '.': 7, '|': 6}
	//initialsing stack
	postFix, stack := []rune{}, []rune{}
	//range loop converts string to char array (UTF-8)
	for _, r := range infix {
		switch {
		case r == '(':
			stack = append(stack, r)
		case r == ')':
			for stack[len(stack)-1] != '(' {
				//next 2 lines pop character off stack and append it to returned postfix
				postFix = append(postFix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			//pops off next character after brackets
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			//if stack isn't empty and the precedence off current char is less/equals to the precedence of whats on the stack
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				//next 2 lines pop character off stack and append it to returned postfix
				postFix = append(postFix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		default:
			postFix = append(postFix, r)
		}
	}
	for len(stack) > 0 {
		//next 2 lines pop character off stack and append it to returned postfix
		postFix = append(postFix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(postFix)
}

//function that manipulates stack of characters to converet postfix
//regular expression to a non-deterministic finite automata
func poregtonfa(postFix string) *nfaFrag {
	//declaring empty array of pointer objects to nfaFrag structs
	nfastack := []*nfaFrag{}
	//looping through expression and switching between each special
	for _, r := range postFix {
		switch r {
		case '.':
			//popping off 2 fragments
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			//joining fragments
			frag1.accept.edge1 = frag2.initial
			//appending new concatenated fragment to stack
			nfastack = append(nfastack, &nfaFrag{initial: frag1.initial, accept: frag2.accept})

		case '|':
			//popping off 2 fragments
			frag2 := nfastack[len(nfastack)-1]
			nfastack := nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			//creating new state at start of nfa joining new state to initial states of frag 1 & 2
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//creating new accept state at end of nfa
			accept := state{}
			//joining accept states of frag 1 & 2 to new accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//appending new concatenated fragment to stack
			nfastack = append(nfastack, &nfaFrag{initial: &initial, accept: &accept})

		case '*':
			//popping one fragment off nfastack
			frag := nfastack[len(nfastack)-1]
			nfastack := nfastack[:len(nfastack)-1]
			//creating new accept state
			accept := state{}
			//create new initial state with edges to fragment and new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//setting fragment accept state arrows back to start of fragment and to the new accept state
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//appending new concatenated fragment to stack
			nfastack = append(nfastack, &nfaFrag{initial: &initial, accept: &accept})

		case '+':
			//popping one fragment off nfastack
			frag := nfastack[len(nfastack)-1]
			nfastack := nfastack[:len(nfastack)-1]
			//creating new accept state
			accept := state{}
			//create new initial state with edge to fragment
			initial := state{edge1: frag.initial}
			//setting fragment accept state arrows back to start of fragment and to the new accept state
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//appending new concatenated fragment to stack
			nfastack = append(nfastack, &nfaFrag{initial: &initial, accept: &accept})

		case '?':
			//popping one fragment off nfastack
			frag := nfastack[len(nfastack)-1]
			nfastack := nfastack[:len(nfastack)-1]
			//assinging symbol of popped fragment to variable s to be used later
			s := frag.initial.symbol
			//creating new state in the case of 2 or more characters
			reject := state{rej: true}
			//creating new accept state pointing to reject state
			accept := state{edge1: &reject, symbol: s}
			//create new initial state with edges to fragment and new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//joining fragment accept to new accept state
			frag.accept.edge1 = &accept
			//appending new fragment to stack
			nfastack = append(nfastack, &nfaFrag{initial: &initial, accept: &accept})

		//default case is for all normal characters (i.e a,b,c)
		default:
			//creating new accept state
			accept := state{}
			//creating new initial state and setting the symbol and only using edge1
			initial := state{symbol: r, edge1: &accept}
			//appending new fragment to stack
			nfastack = append(nfastack, &nfaFrag{initial: &initial, accept: &accept})

		}

	}
	//returning fully concatenated nfa as a whole
	return nfastack[0]

}

//function to add states to array of pointers (current/next)
func addState(l []*state, s *state, a *state) []*state {
	//adding desired state to array
	l = append(l, s)
	//checking it isnt the accept state and its arrow label is e and its not a reject state
	if s != a && s.symbol == 0 && s.rej == false {
		//adding states meeting above condition
		l = addState(l, s.edge1, a)
		//if there's a 2nd edge add relevant state
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	//returning new current array
	return l
}

//function returning boolean on a regular expression matching a given string
func postFixMatch(postFix string, str string) bool {
	//set returned boolean to false to start
	matched := false

	//convert postfix regex to an non-determistic finite automata
	pfixNfa := poregtonfa(postFix)

	//array of state pointers containing all the current states of nfa
	current := []*state{}

	//after reading a character from the string this contains all next states
	next := []*state{}

	current = addState(current[:], pfixNfa.initial, pfixNfa.accept)
	//looping through given string one rune at a time
	for _, r := range str {
		//looping through array of current states
		for _, c := range current {
			//if current rune is the same as the arrow labels of the current state
			if c.symbol == r && c.rej == false {
				next = addState(next[:], c.edge1, pfixNfa.accept)
			}
		}

		//setting current array to next(making move) for next rune being read in and resetting next array to null
		current, next = next, []*state{}
	}
	//loop through current states after nfa is finished
	for _, c := range current {
		//if one of the current states is an accept state its a mathc
		if c == pfixNfa.accept {
			matched = true
			break
		}
	}

	//returning whether it's a match or not
	return matched
}

func main() {
	//boolean variables used to control UI loop
	var exit bool
	var change bool

	for exit == false{
		//ask for user input
		fmt.Print("Enter Regular Expression: ")
		var regex string
		fmt.Scanln(&regex)
		if regex == "" {
			fmt.Println("No empty strings please")
			break
		}
		for change ==false{
			fmt.Print("Enter String('e' to exit or 'c' to change regular expression): ")
			var str string
			fmt.Scanln(&str)
			if str == "e"{
				exit = true
				return
			} else if str == "c"{
				break
			} else {
				//convert given regular expression into postfix notation
				postFixRegex := IntoPostFix(regex)
				//check if its a match by converting postfix to nfa and
				//then going through the nfa and seeing if you finish in an accept state
				matched := postFixMatch(postFixRegex, str)
				//output result to user
				if matched == true {
					fmt.Print("It's a match \n")
				} else {
					fmt.Print("It's not a match \n")
				}
			}
		}
		
	}
	
}
