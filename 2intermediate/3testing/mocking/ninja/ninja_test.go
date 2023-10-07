package ninja

import (
	"fmt"
	"testing"
)

type MockSecretGen struct {
	secret int
}

func (m MockSecretGen) Get() int {
	return m.secret
}

func TestNinja_Greeting(t *testing.T) {
	ninjaName := "Tommy"
	secret := 9001
	sg := MockSecretGen{secret}
	tommy := Ninja{
		Name:      ninjaName,
		SecretGen: sg,
	}

	secretGreeting := tommy.Greeting()
	expectedSecretGreeting :=
		fmt.Sprintf("I'm ninja %s & my secret is %d.",
			ninjaName, secret)
	if secretGreeting != expectedSecretGreeting {
		t.Error("Unexpected secret greeting!")
	}
}
