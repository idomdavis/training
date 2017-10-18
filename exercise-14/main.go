package main

// Implement stateMachine.
type state struct {
	name string
	zero *state
	one  *state
}

func (s *state) listen(zero chan struct{}, one chan struct{}, state chan *state) {
	select {
	case <-zero:
		state <- s.zero
	case <-one:
		state <- s.one
	}
}

type stateMachine struct {
	states  []*state
	current *state
	zero    chan struct{}
	one     chan struct{}
	s       chan *state
}

func (s *stateMachine) send(i int) {
	switch i {
	case 0:
		s.zero <- struct{}{}
	case 1:
		s.one <- struct{}{}
	default:
		panic("send only support 0 or 1")
	}

	s.current = <-s.s
	go s.current.listen(s.zero, s.one, s.s)
}

func (s *stateMachine) state() string {
	close(s.one)
	close(s.zero)
	close(s.s)
	return s.current.name
}

func newStateMachine() *stateMachine {
	a := &state{name: "A"}
	b := &state{name: "B"}
	c := &state{name: "C"}

	a.zero = a
	a.one = b
	b.zero = c
	b.one = a
	c.zero = b
	c.one = a

	zero := make(chan struct{})
	one := make(chan struct{})
	statechan := make(chan *state)

	go a.listen(zero, one, statechan)

	s := &stateMachine{
		states:  []*state{a, b, c},
		current: a,
		zero:    zero,
		one:     one,
		s:       statechan,
	}

	return s
}

func main() {
	sm := newStateMachine()
	sm.send(1)          // "state A + 1 => state B"
	sm.send(0)          // "state B + 0 => state C"
	println(sm.state()) // "state C"
}
