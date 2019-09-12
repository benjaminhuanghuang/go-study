package main

import "fmt"

type Animal interface {
	Fly()
	Run()
}

type Creature interface {
	Run()
}

type Bird struct {
}

// implment interface
func (bird Bird) Fly() {
	fmt.Println("Bird is flying")
}

func (bird Bird) Run() {
	fmt.Println("Bird can not run")
}

func main() {
	var animal Animal
	var creature Creature

	animal = new(Bird)

	creature = animal
	creature.Run()
}
