package main

type Human struct {
	name string
	age  int
}

func NewHuman(name string, age int) *Human {
	return &Human{name, age}
}

func (h *Human) Name() string {
	return h.name
}

func (h *Human) Age() int {
	return h.age
}

func (h *Human) SetAge(age int) {
	h.age = age
}
