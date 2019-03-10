package re

import (
	"errors"
)

// acceptingStates
// transitions

type state struct {
}

type engine struct {
	match      string
	startState *state
}

func newEngineNFA(expr string) (e *engine, err error) {
	return nil, errors.New("TODO")
}

func (e *engine) init() {

}

func (e *engine) next(r rune) error {
	return nil
}
