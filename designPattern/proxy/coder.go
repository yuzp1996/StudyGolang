package proxy

// Coder repersent the coder with Command
type Coder struct {
	Name string
	Command GitCmd
}

// Clone represent the function for Coder
func (c Coder) Clone() {
	c.Command.Clone()
}

// NewCoder init a coder
func NewCoder(name string, command GitCmd) *Coder {
	return &Coder{Name: name, Command: command}
}


