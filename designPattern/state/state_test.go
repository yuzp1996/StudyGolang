package state_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/yuzp1996/StudyGolang/designPattern/state"
)

var _ = Describe("State", func() {
	It("simulate gumballgame", func() {
		gumballgame := state.NewGumballGame(10,&state.SoldOutState{},&state.SoldState{},&state.NoQuarterState{},nil)
		gumballgame.Soldstate = state.NewSoldState(*gumballgame)
		gumballgame.Soldoutstate = state.NewSoldOutState(*gumballgame)
		gumballgame.Noquarterstate = state.NewNoQuarterState(*gumballgame)
		gumballgame.SetState(gumballgame.Soldstate)
		gumballgame.Insert()
		gumballgame.Eject()
		gumballgame.Truncrank()

	})


})
