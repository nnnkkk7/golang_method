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
	fmt.Printf("I go to %s, name:%s, age: %s, my power is %s, and speed is %s", s.school, s.name, s.age, s.power, s.speed)
}

func main() {
	you := Student{Human{"john", 18, "strong", "fast"}, "univ"}
	me := Human{"nao", 3, "weak", "slow"}

	you.sayHello()
	me.sayHello()
}
