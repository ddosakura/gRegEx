package fa

// Input Symbol
type Input interface {
	Eq(Input) bool
}

// Inputer Symbol Stream
type Inputer interface {
	Next() Input
}

// State of FA
type State interface {
	IsFinal() bool
}

// FA - Finite Automaton
type FA interface {
	AddState(s State)
	AddTransition(src State, i Input, dst ...State)

	Reset()
	IsFinal() bool
	Next(i Input) []State
}

// Verify the inputs
func Verify(fa FA, ir Inputer) bool {
	fa.Reset()
	for i := ir.Next(); i != nil; i = ir.Next() {
		fa.Next(i)
	}
	return fa.IsFinal()
}

// TransitionInput for state
type TransitionInput struct {
	S State
	I Input
}
