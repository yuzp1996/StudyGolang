package structasclassemployee

import "fmt"

type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

// 你看这个结构体是小写的 在外面是引用不了的  所以在这里给出去了  成功的避免了创建不可用的employee
func New(fname, laname string, toleave, leataken int) employee {
	employ := employee{
		fname, laname, toleave, leataken,
	}
	return employ
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining \n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
