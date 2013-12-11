package inject_test

import (
	"fmt"
	"os"
	"github.com/ParsePlatform/go.inject"
)

// Interfaces

type CARFACTORY interface {
	makeCar() CAR     `inject:""`
	getMake() string  `inject:""`
}

type CAR interface {
	getModel() string `inject:""`
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
	var graph inject.Graph

	var carFactory FordFactory

	err := graph.Provide(
		&inject.Object{Value: &carFactory},
		&inject.Object{Value: http.DefaultTransport},
		)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}



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
