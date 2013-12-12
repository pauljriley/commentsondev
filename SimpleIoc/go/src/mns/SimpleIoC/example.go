package main

import (
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

// 'Class' factories

type CarFactoryCreator func() CARFACTORY

type CarFactoryFactory struct {
	m map[string]CarFactoryCreator
}

func NewCarFactoryFactory() *CarFactoryFactory {
	return &CarFactoryFactory{m: make(map[string]CarFactoryCreator, 10)}
}

func (cf *CarFactoryFactory) MakeCarFactory(cft string) CARFACTORY {
	fn, ok := cf.m[cft]
	if !ok {
		panic("whatever")
	}
	return fn()
}

func (cf *CarFactoryFactory) Register(cft string, fn CarFactoryCreator) {
	cf.m[cft] = fn
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
		s.car = new(FordMondeo)
	}

	return s.car
}

func (s FordFactory) getMake() string {
	return "Ford"
}

// Main function

func main() {
	factory := NewCarFactoryFactory()
	factory.Register("Ford", func() CARFACTORY { return FordFactory{car: new(FordMondeo)} })

	myCarFactory := factory.MakeCarFactory("Ford")
	myCar := myCarFactory.makeCar()
	fmt.Println(myCar.getModel())
}
