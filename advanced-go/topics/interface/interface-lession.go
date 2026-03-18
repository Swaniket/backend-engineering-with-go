package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

// We can think about interface as a blueprint about the signature for the methods
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

// @NOTE: Both NormalTruck & ElectricTruck has LoadCargo() & UnloadCargo() methods implemented, hence the Truck interface is automatically assigned with the go compiler.
type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery = -1

	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery += -1

	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing Truck %+v \n", truck)

	// @NOTE: with the Truck interface, we don't have access to it's internal property like battery/cargo

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return nil
}

func main() {
	nt := &NormalTruck{id: "1"}
	et := &ElectricTruck{id: "2"}

	err := processTruck(nt)
	if err != nil {
		log.Fatalf("Error processing the truck: %s", err)
	}

	// Here we are re-assigning the err variable
	err = processTruck(et)
	if err != nil {
		log.Fatalf("Error processing the truck: %s", err)
	}

	log.Println(nt.cargo)
	log.Println(et.battery)
}
