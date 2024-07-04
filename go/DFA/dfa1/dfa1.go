package main

import (
	"fmt"
)

type DFA struct {
	initialState string
	acceptStates map[string]bool
	transition   map[string]map[rune]string
	currentState string
}

func NewDFA(initialState string, acceptStates []string, transition map[string]map[rune]string) *DFA {
	acceptStatesMap := make(map[string]bool)
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

func (d *DFA) Reset() {
	d.currentState = d.initialState
}

func (d *DFA) Process(symbol rune) {
	if nextState, ok := d.transition[d.currentState][symbol]; ok {
		d.currentState = nextState
	} else {
		d.currentState = "ERROR"
	}
}

func (d *DFA) Accept() bool {
	return d.acceptStates[d.currentState]
}

func (d *DFA) ProcessString(input string) bool {
	d.Reset()
	for _, symbol := range input {
		d.Process(symbol)
		if d.currentState == "ERROR" {
			return false
		}
	}
	return d.Accept()
}

func main() {
	transition := map[string]map[rune]string{
		"q0": {'0': "q1", '1': "q1"},
		"q1": {'0': "q2", '1': "q2"},
		"q2": {},
	}

	acceptStates := []string{"q2"}
	dfa := NewDFA("q0", acceptStates, transition)

	testStrings := []string{
		"0", "1", "00", "01", "10", "11",
		"000", "001", "010", "011", "100", "101", "110", "111",
	}

	for _, s := range testStrings {
		fmt.Printf("String: %s, Accepted: %v\n", s, dfa.ProcessString(s))
	}
}
