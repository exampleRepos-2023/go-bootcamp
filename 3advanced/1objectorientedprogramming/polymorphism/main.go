package main

import "fmt"

type ninjaStar struct {
	owner string
}

type ninjaSword struct {
	owner string
}

// This enables polymorphism
// Duck typing - If it walks like a duck and quacks like a duck, then it must be a duck
type ninjaWeapon interface {
	attack()
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star...")
}

func (ninjaSword) attack() {
	fmt.Println("Swinging ninja sword...")
}

func main() {
	star := ninjaStar{"Tommy"}
	sword := ninjaSword{"Bobby"}
	star.attack()
	sword.attack()

	weapon := ninjaWeapon(star)
	weapon.attack()
	weapon = ninjaWeapon(sword)
	weapon.attack()

	stars := []ninjaStar{star, ninjaStar{"Jimmy"}}
	swords := []ninjaSword{sword, ninjaSword{"Johnny"}}
	for _, star := range stars {
		star.attack()
	}
	for _, sword := range swords {
		sword.attack()
	}

	weapons := []ninjaWeapon{star, sword}
	for _, weapon := range weapons {
		weapon.attack()
	}

}
