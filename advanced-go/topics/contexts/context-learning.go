package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type contextKey string

var UserIDKey contextKey = "userId"

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

func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("Processing Truck %+v \n", truck)

	// access the user id from context
	// userId := ctx.Value(UserIDKey)
	// log.Println(userId)

	// Creating a timeout from context to cancel long running operations
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel() // Calling cancel after the function is finished

	// Simulate a long running code
	// delay := time.Second * 3
	delay := time.Second * 1
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

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

func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			err := processTruck(ctx, t)

			if err != nil {
				log.Println(err)
			}

			wg.Done()
		}(t)
	}

	wg.Wait()

	return nil
}

func main() {
	// A context has a box that carries contexual information about that request that mostly used to carry information or cancel long running requests.
	// 1. Context allows us to cancel long running operations
	// 2. Context is also used to pass meta data between API boundaries
	// 3. The same context can be passed to functions running in different go-routines.
	// 4. Context is immutable, so to attach a value we need to create a new context and attach the value

	ctx := context.Background() // Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests
	// ctx := context.TODO() // TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).
	ctx = context.WithValue(ctx, UserIDKey, 42)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v \n", err)
		return
	}

	fmt.Println("All trucks processed successfully!")
}
