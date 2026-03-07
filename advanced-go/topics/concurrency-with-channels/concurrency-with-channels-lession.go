package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
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

	// ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	// defer cancel()

	// delay := time.Second * 5
	// select {
	// case <-ctx.Done():
	// 	return ctx.Err()
	// case <-time.After(delay):
	// 	break
	// }

	// err := truck.LoadCargo()
	// if err != nil {
	// 	return fmt.Errorf("Error loading cargo %w", err)
	// }

	// err = truck.UnloadCargo()
	// if err != nil {
	// 	return fmt.Errorf("Error un-loading cargo %w", err)
	// }

	// fmt.Printf("Finish processing Truck %+v \n", truck)
	// return nil

	return ErrTruckNotFound
}

func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup

	errorsChannel := make(chan error, len(trucks)) // len(trucks) numbers of channels will be created to prevent deadlock
	// defer close(errorsChannel)                     // We need to close a channel that we open

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			err := processTruck(ctx, t)

			if err != nil {
				log.Println(err)
				errorsChannel <- err // Sends the error to the channel
			}

			wg.Done()
		}(t)
	}

	wg.Wait()
	close(errorsChannel)

	// Listen to single error
	// select {
	// case err := <-errorsChannel: // Pipeing any errors that are in errors channel to the err variable.
	// 	return err
	// default:
	// 	return nil
	// }

	// Listen to a slice of errors
	var errs []error

	for err := range errorsChannel {
		log.Printf("Error processing truck %v \n", err)
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("fleet processing had %d errors", len(errs))
	}

	return nil
}

func main() {
	// Go routines lets you run concurrent tasks
	// Channels acts as a pipeline for to send and recieve data between go-routines

	ctx := context.Background()
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
