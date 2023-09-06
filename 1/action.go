package main

type Action struct {
	*Human
}

func NewAction(h *Human) *Action {
	return &Action{h}
}
