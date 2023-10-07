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

func (ninjaStar) pickup() {
	fmt.Println("Picking back up ninja star")
}

func (ninjaSword) attack() {
	fmt.Println("Swinging ninja sword")
}

func main() {
	// Type assertion allows for dynamic type evaluation AND conversion of an interface type
	weapons := []ninjaWeapon{ninjaSword{}, ninjaStar{}}
	for _, weapon := range weapons {
		weapon.attack()
		if ns, ok := weapon.(ninjaStar); ok {
			ns.pickup()
		}
	}
}
