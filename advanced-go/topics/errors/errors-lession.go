package main

import (
	"errors"
	"fmt"
	"log"
)

// Custom error
var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck struct {
	id    string
	cargo int
}

// @NOTE: We can also write custom error
// type customError interface {
// 	Error() string
// 	// We can add optional meta data or any other things
// }

// @NOTE: Errors are usually the last return from a function
// So if we have multiple returns, it would do
// func processTruck(truck Truck) (string, float64, error)

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err) // -> Other way of Creating an error
	}

	// return errors.New("some error") -> Creating an error
	// return ErrNotImplemented
	return nil // We either return an error or return "nil"
}

func RunErrorsLesson() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Trucks %s arrived. \n", truck.id)

		// Option 1 to process error: More vurbose
		// err := processTruck(truck)
		// if err != nil {
		// 	log.Fatalf("Error processing truck: %s", err) // @NOTE: log.Fatalf() will terminate the program and use Printf() to print/log the error
		// }

		// Option 2: Inline error handling
		// @NOTE: Difference between Option 1 & Option 2: with option 2, the err object will get garbage collected after the nil check
		if err := processTruck(truck); err != nil {
			// Handling of custom error
			if errors.Is(err, ErrNotImplemented) {
				// We do this
			}

			if errors.Is(err, ErrTruckNotFound) {
				// We do this
			}

			// Generic Error Handling
			log.Fatalf("Error processing truck: %s", err)
		}
		// log.Fatalf("Error processing truck: %s", err) -> This will not work when we use Option 2
	}
}
