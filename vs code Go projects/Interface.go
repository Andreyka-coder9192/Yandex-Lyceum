package main

import "fmt"

type Human interface {
	SayHello()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main1() {
	var h Human = Person{Name: "John", Age: 25}
	h.SayHello()
}
