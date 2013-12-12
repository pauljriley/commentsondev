package main

import (
	"code.google.com/p/go-inject"
	"fmt"
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
	if s.car == nil {
		// s.car = new(FordMondeo)
		fmt.Println("Car?")
	}

	return s.car
}

func (s FordFactory) getMake() string {
	return "Ford"
}

// Injector tags

type GetCarFactory struct{}

// Main function

func main() {
	injector := inject.CreateInjector()

	injector.Bind(GetCarFactory{}, func(context inject.Context, container inject.Container) interface{} {
		return FordFactory{car: new(FordMondeo)}
	})
	container := injector.CreateContainer()

	myCarFactory := container.GetInstance(nil, GetCarFactory{}).(CARFACTORY)
	myCar := myCarFactory.makeCar()
	fmt.Println(myCar.getModel())
}
