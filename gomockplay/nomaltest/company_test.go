package nomaltest

import "testing"

func TestCompany_Meeting(t *testing.T) {
	company := NewCompany(NewPerson("小妹"))
	t.Log(company.Meeting("老板"))
}