package main

import (
	"fmt"
	"github.com/ParsePlatform/go.inject"
	"os"
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
	car *FordMondeo `inject:""`
}

type FordMondeo struct {
}

func (s *FordMondeo) getModel() string {
	return "Mondeo"
}

func (s FordFactory) makeCar() CAR {
	return s.car
}

func (s FordFactory) getMake() string {
	return "Ford"
}

// Main function

func main() {
	var graph inject.Graph

	var carFactory FordFactory

	err := graph.Provide(
		&inject.Object{Value: new(FordMondeo)},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := graph.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("exit 2")
		os.Exit(1)
	}

	myCar := carFactory.makeCar()
	fmt.Println(myCar.getModel())
}
