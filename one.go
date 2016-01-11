package main

import "fmt"

type Animal interface {
	eat()
	walk()
}
type Dog struct {
}

func (dog Dog) eat() {
	fmt.Println("I am eating")
}
func (dog Dog) walk() {
	fmt.Println("Dog is walking")
}

type Cat struct {
}

func (cat Cat) eat() {
	fmt.Println("I am eating")
}
func (cat Cat) walk() {
	fmt.Println("Cat is walking")
}
func main() {
	d := Dog{}
	c := Cat{}
	animals := [...]Animal{d, c}
	for _, v := range animals {
		v.eat()
		v.walk()
	}
}
