package nfa

import (
	"github.com/ddosakura/gRegEx/fa"
)

// NFA - Nondeterministic Finite Automaton
type NFA struct {
	transition map[fa.TransitionInput]map[fa.State]bool
	inputMap   map[fa.Input]bool
	initState  fa.State

	currentState map[fa.State]bool
	totalStates  []fa.State
	finalStates  []fa.State
}

// New a new NFA
func New(s fa.State) fa.FA {

	retNFA := &NFA{
		transition:   make(map[fa.TransitionInput]map[fa.State]bool),
		inputMap:     make(map[fa.Input]bool),
		initState:    s,
		currentState: make(map[fa.State]bool),
	}

	retNFA.currentState[s] = true
	retNFA.AddState(s)
	return retNFA
}

//AddState in this NFA
func (d *NFA) AddState(s fa.State) {
	if s == nil {
		panic("state should not be nil")
	}

	d.totalStates = append(d.totalStates, s)
	if s.IsFinal() {
		d.finalStates = append(d.finalStates, s)
	}
}

//AddTransition new transition function into NFA
func (d *NFA) AddTransition(src fa.State, i fa.Input, dst ...fa.State) {
	find := false

	for _, v := range d.totalStates {
		if v == src {
			find = true
		}
	}

	if !find {
		panic("No such state in current NFA")
	}

	if i == nil {
		panic("Not allow nil input in NFA")
	}

	//find input if exist in NFA input List
	if _, ok := d.inputMap[i]; !ok {
		//not exist, new input in this NFA
		d.inputMap[i] = true
	}

	dstMap := make(map[fa.State]bool)
	for _, destState := range dst {
		dstMap[destState] = true
	}

	d.transition[fa.TransitionInput{
		S: src,
		I: i,
	}] = dstMap
}

// Reset NFA state to initilize state
func (d *NFA) Reset() {
	initState := make(map[fa.State]bool)
	initState[d.initState] = true
	d.currentState = initState
}

// IsFinal of current state
func (d *NFA) IsFinal() bool {
	for _, val := range d.finalStates {
		for cState := range d.currentState {
			if val == cState {
				return true
			}
		}
	}
	return false
}

// Next state of NFA
func (d *NFA) Next(i fa.Input) []fa.State {
	updateCurrentState := make(map[fa.State]bool)
	for current := range d.currentState {
		if valMap, ok := d.transition[fa.TransitionInput{
			S: current,
			I: i,
		}]; ok {
			for dst := range valMap {
				updateCurrentState[dst] = true
			}
		} else {
			//dead state, remove in current state
			//do nothing.
		}
	}

	//update curret state
	d.currentState = updateCurrentState

	//return result
	var ret []fa.State
	for state := range updateCurrentState {
		ret = append(ret, state)
	}
	return ret
}
