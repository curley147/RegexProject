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
}

//struct to represent smaller subsets of overall nfa
type nfaFrag struct {
	initial *state
	accept  *state
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

func main() {

	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
