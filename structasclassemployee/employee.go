package structasclassemployee

import "fmt"

type employee struct{
	firstName string
	lastName string
	totalLeaves int
	leavesTaken int
}

func New(fname,laname string, toleave, leataken int) employee{
	employ := employee{
		fname,laname,toleave,leataken,
}
	return employ
}

func(e employee) LeavesRemaining(){
	fmt.Printf("%s %s has %d leaves remaining \n", e.firstName,e.lastName,(e.totalLeaves-e.leavesTaken))
}