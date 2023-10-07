package main

import "fmt"

type ninjaStar struct {
}

type ninjaSword struct {
}

type ninjaWeapon interface {
	attack()
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star")
}

func (ninjaSword) attack() {
	fmt.Println("Swinging ninja sword")
}

func main() {
	// Type switch allows for dynamic type evaluation of an interface type
	weapons := []ninjaWeapon{ninjaStar{}, ninjaSword{}}
	for _, weapon := range weapons {
		switch weapon.(type) {
		case ninjaStar:
			fmt.Println("About to throw ninja star")
		case ninjaSword:
			fmt.Println("About to swing ninja sword")
		}
		weapon.attack()
	}
}
