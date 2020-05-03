package state

import "fmt"

type State interface {
	insert()
	eject()
	truncrank()
}

type SoldOutState struct {
	*GumballGame
}

func NewSoldOutState(game GumballGame) *SoldOutState {
	return &SoldOutState{&game}
}

func (s *SoldOutState) insert() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","SoldOutState","insert")
}

func (s *SoldOutState) eject() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","SoldOutState","eject")
}

func (s *SoldOutState) truncrank() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","SoldOutState","truncrank")
}

type SoldState struct {
	*GumballGame
}

func NewSoldState(game GumballGame) *SoldState {
	return &SoldState{&game}
}

func (s *SoldState) insert() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","SoldState","insert")
}

func (s *SoldState) eject() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","SoldState","eject")
}

func (s *SoldState) truncrank() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","SoldState","truncrank")
}

type NoQuarterState struct {
	*GumballGame
}

func NewNoQuarterState(game GumballGame) *NoQuarterState {
	return &NoQuarterState{&game}
}

func (n *NoQuarterState) insert() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","NoQuarterState","insert")
}

func (n *NoQuarterState) eject() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","NoQuarterState","eject")
}

func (n *NoQuarterState) truncrank() {
	fmt.Printf("Now the state is %s, and the behvior is %s \n","NoQuarterState","truncrank")
}


