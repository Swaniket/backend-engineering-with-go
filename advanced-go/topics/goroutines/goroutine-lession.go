package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

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

	// Simulate some processing time
	time.Sleep(time.Second)

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("Error loading cargo %w", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error un-loading cargo %w", err)
	}

	fmt.Printf("Finish processing Truck %+v \n", truck)
	return nil
}

// processFleet demonstrates concurrent processing of multiple trucks
// GoRoutines are light weight threads that go manages for us
func processFleet(trucks []Truck) error {
	// Running without go-routines
	// for _, t := range trucks {
	// 	processTruck(t)
	// }
	var wg sync.WaitGroup // Step 1: A WaitGroup waits for a collection of goroutines to finish
	// wg.Add(len(trucks)) // Step 2A: -> We can either do this

	// for _, t := range trucks {
	// 	go processTruck(t) // If we just do that, then the processes will get leaked and nothing will get processed
	// }

	for _, t := range trucks {
		// Step 2B: Or add it in the loop
		wg.Add(1)

		go func(t Truck) {
			err := processTruck(t)

			if err != nil {
				log.Println(err)
			}

			wg.Done() // Step 3: We also need to do this
		}(t)
	}

	wg.Wait() // Step 4: Wait for the go routines to get finished

	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	// Process all trucks concurrently
	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleet: %v \n", err)
		return
	}

	fmt.Println("All trucks processed successfully!")

	// Temp fix - we can sleep for more than the waiting time to finish it
	// time.Sleep(5 * time.Second)
}
