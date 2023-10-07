package ninja

import "fmt"

type Ninja struct {
	Name  string
	Level int
}

func (n *Ninja) Greet() {
	fmt.Printf("My name is %s & I'm at level %d\n", n.Name, n.Level)
	// More implementation details
	n.Level++
	n.Name += " Sr"
}

func New() *Ninja {
	return &Ninja{"Tommy", 9001}
}
