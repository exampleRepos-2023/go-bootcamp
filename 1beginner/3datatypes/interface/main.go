package main

import "fmt"

type ninjaStar struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star...")
}

type ninjaSword struct {
	owner string
}

func (ninjaSword) attack() {
	fmt.Println("Swinging ninja sword...")
}

type ninjaWeapon interface {
	attack()
}

// Duck typing - If it walks like a duck and it quacks like a duck, then it must be a duck

func main() {
	// Interfaces specify the behaviors of an object,
	// and categorize ones with the same behaviors as the same type
	stars := []ninjaStar{{"Tommy"}, {"Bobby"}}
	for _, star := range stars {
		star.attack()
	}
	swords := []ninjaSword{{"Tommy"}, {"Bobby"}}
	for _, sword := range swords {
		sword.attack()
	}

	weapons := []ninjaWeapon{ninjaStar{"Tommy"}, ninjaSword{"Bobby"}}
	for _, weapon := range weapons {
		weapon.attack()
	}
}
