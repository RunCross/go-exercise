package oop

func (o *Oop) Add() int {
	return 5
}

type oppin interface {
	apd() int
}

type oppout interface {
	oppin
	aad() int
}
