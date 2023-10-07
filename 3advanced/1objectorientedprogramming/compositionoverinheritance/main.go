package main

import "fmt"

type alive struct {
}

type walkable struct {
}

type swimmable struct {
}

type duck struct {
	a alive
	w walkable
	s swimmable
}

func (d duck) eat() {
	d.a.eat()
}

func (d duck) sleep() {
	d.a.sleep()
}

func (d duck) walk() {
	d.w.walk()
}

func (d duck) swim() {
	d.s.swim()
}

func (alive) eat() {
	fmt.Println("Alive & eating")
}

func (alive) sleep() {
	fmt.Println("Alive & sleeping")
}

func (walkable) walk() {
	fmt.Println("Strolling through")
}

func (swimmable) swim() {
	fmt.Println("Taking a dip")
}

type goose struct {
	a alive
	w walkable
	s swimmable
	f flyable
}

type flyable struct {
}

func main() {
	d := duck{}
	d.eat()
	d.sleep()
	d.walk()
	d.swim()
}
