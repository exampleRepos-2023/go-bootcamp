package ninja

import "fmt"

type Ninja struct {
	name  string
	level int
}

func (n *Ninja) Greet() {
	fmt.Printf("My name is %s & I'm at level %d\n", n.name, n.level)
}

func New() *Ninja {
	return &Ninja{"Tommy", 9001}
}
