package nomaltest

import "fmt"

type Person struct {
	name string
}

func NewPerson(name string)*Person{
	return &Person{name:name}
}

func (p *Person)SayHello(name string)(response string){
	return fmt.Sprintf("Hello %s, welcome come to here My name is %v ",name,p.name)
}