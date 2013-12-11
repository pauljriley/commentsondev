package main

import (
	"fmt"
	"github.com/99designs/goodies/depinject"
)

// Interfaces

type CARFACTORY interface {
	makeCar() CAR
	getMake() string
}

type CAR interface {
	getModel() string
}

// Concrete implementations

type FordFactory struct {
	car CAR
}

type FordMondeo struct {
}

func (s *FordMondeo) getModel() string {
	return "Mondeo"
}

func (s FordFactory) makeCar() CAR {
	return new(FordMondeo)
}

func (s FordFactory) getMake() string {
	return "Ford"
}

// Main function

func main() {
	di := depinject.NewDependencyInjector()

	di.MustRegister(func() CAR {
		return new(FordMondeo)
	})
	di.MustRegister(func(myCar CAR) FordFactory {
		return FordFactory{car: myCar}
	})

	myCarFactory := di.Create(FordFactory{}).(FordFactory)
	myCar := myCarFactory.makeCar()
	fmt.Println(myCar.getModel())
}
