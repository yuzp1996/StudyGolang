package proxy

import "fmt"

// GitCmd represent the proxy between coder and github
type GitCmd struct {
	Server Clone
	Before string
	Name string
	After string
}

func NewGitCmd(server Clone, before string, name string, after string) *GitCmd {
	return &GitCmd{Server: server, Before: before, Name: name, After: after}
}

// Clone will do the clone The proxy function
func (g GitCmd) Clone() {
	fmt.Printf("Before %s should do 『%s』\n",g.Name,g.Before)
	g.Server.Clone()
	fmt.Printf("\nAfter %s should do 『%s』\n",g.Name,g.After)
}