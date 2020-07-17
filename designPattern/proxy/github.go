package proxy

import "fmt"

// GitHub represent the
type GitHub struct {
	Name string
}

func NewGitHub(name string) *GitHub {
	return &GitHub{Name: name}
}

// Clone will execute the clone
func (github GitHub)Clone(){
	fmt.Printf("I am %s, Now you can clone from me",github.Name)
}
