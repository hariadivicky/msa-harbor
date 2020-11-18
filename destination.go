package main

import "fmt"

// Visitable defines structure with enter channel.
type Visitable interface {
	GetEnterChannel() chan<- *Car
	fmt.Stringer // visitable place must be printable.
}

// StartingPoint .
type StartingPoint interface {
	Visitable
	IsStartingPoint() bool
}

// Place is openable destination (listening for cars queue).
type Place interface {
	Visitable
	Open()
}
