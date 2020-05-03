package state

type GumballGame struct {
	Num int
	Soldoutstate *SoldOutState
	Soldstate *SoldState
	Noquarterstate *NoQuarterState
	State State
}

func (g *GumballGame) Insert() {
	g.State.insert()
	g.SetState(g.Soldoutstate)
}

func (g *GumballGame) Eject() {
	g.State.eject()
	g.SetState(g.Noquarterstate)
}

func (g *GumballGame) Truncrank() {
	g.State.truncrank()
	g.SetState(g.Soldstate)

}

func(g *GumballGame)SetState(state State){
	g.State = state
}

func NewGumballGame(num int, soldoutstate *SoldOutState, soldstate *SoldState, noquarterstate *NoQuarterState, state State) *GumballGame {
	return &GumballGame{Num: num, Soldoutstate: soldoutstate, Soldstate: soldstate, Noquarterstate: noquarterstate, State: state}
}


