package main

import (
	"code.google.com/p/go-inject"
	"fmt"
	"reflect"
)

type CAR interface {
	Make() string
	Model() string
	Colour() string
}

type Ford struct {
}

func (c *Ford) Make() string {
	return "Ford"
}

func (c *Ford) Model() string {
	return "Modeao"
}

func (c *Ford) Colour() string {
	return "Red"
}

func main() {
	injector := inject.CreateInjector()
	var ford Ford
	injector.BindInstance(reflect.TypeOf((*CAR)(nil)), ford)
	container := injector.CreateContainer()
	car := container.GetInstance(nil, reflect.TypeOf((*CAR)(nil)))
	fmt.Println(fmt.Sprintf("Make : %s, Model : %s, Colour : %s", car.Make(), car.Model(), car.Colour()))
}
