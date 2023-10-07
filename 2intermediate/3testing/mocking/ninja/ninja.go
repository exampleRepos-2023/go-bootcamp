package ninja

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/2intermediate/3testing/mocking/secretgen"
)

type Ninja struct {
	Name      string
	SecretGen secretgen.SG
}

func (n Ninja) Greeting() string {
	return fmt.Sprintf("I'm ninja %s & my secret is %d.",
		n.Name, n.SecretGen.Get())
}
