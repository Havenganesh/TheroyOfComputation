///*  DFA, which accepts binary strings containing an even number of 0s: */

package main

import "fmt"

// DFA struct representing a deterministic finite automaton
type DFA struct {
	initialState int
	acceptStates map[int]bool
	transition   map[int]map[rune]int
	currentState int
}

// NewDFA initializes and returns a new DFA
func NewDFA(initialState int, acceptStates []int, transition map[int]map[rune]int) *DFA {
	acceptStatesMap := make(map[int]bool)
	for _, state := range acceptStates {
		acceptStatesMap[state] = true
	}

	return &DFA{
		initialState: initialState,
		acceptStates: acceptStatesMap,
		transition:   transition,
		currentState: initialState,
	}
}

// Reset resets the DFA to its initial state
func (d *DFA) Reset() {
	d.currentState = d.initialState
}

// Process processes a single character and transitions to the next state
func (d *DFA) Process(c rune) {
	if nextState, ok := d.transition[d.currentState][c]; ok {
		d.currentState = nextState
	} else {
		// No valid transition, stay in current state (or handle as needed)
		d.currentState = -1 // Invalid state
	}
}

// Accept checks if the DFA is in an accept state
func (d *DFA) Accept() bool {
	return d.acceptStates[d.currentState]
}

// ProcessString processes an entire string through the DFA
func (d *DFA) ProcessString(s string) bool {
	d.Reset()
	for _, c := range s {
		d.Process(c)
		if d.currentState == -1 {
			return false
		}
	}
	return d.Accept()
}

func main() {
	// Define transitions for the DFA
	transition := map[int]map[rune]int{
		0: {'0': 1, '1': 0},
		1: {'0': 0, '1': 1},
	}

	// Create a DFA with initial state 0, accept state 0
	dfa := NewDFA(0, []int{0}, transition)

	// Test strings
	testStrings := []string{
		"0",
		"1",
		"00",
		"01",
		"10",
		"11",
		"000",
		"001",
	}

	for _, s := range testStrings {
		fmt.Printf("String: %s, Accepted: %t\n", s, dfa.ProcessString(s))
	}
}
