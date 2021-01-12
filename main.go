package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	power string
	speed string
}

type Member struct {
	Human
	team string
}

type Student struct {
	Human
	school string
}

func (h *Human) sayHello() {
	fmt.Printf("name:%s, age: %s, my power is %s, and speed is %s", h.name, h.age, h.power, h.speed)
}

func (s *Student) sayHello() {
	fmt.Printf("I go to %s", s.school)
}

func main() {
	you := &Human{"john", 18, "strong", "fast"}
	you.sayHello
}
