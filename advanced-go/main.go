package main

import (
	// "context"
	// "errors"
	"fmt"
	// "log"
	"sync"
	"time"
)

// type contextKey string

// var UserIDKey contextKey = "userId"

// var (
// 	ErrNotImplemented = errors.New("not implemented")
// 	ErrTruckNotFound  = errors.New("truck not found")
// )

// type Truck interface {
// 	LoadCargo() error
// 	UnloadCargo() error
// }

// type NormalTruck struct {
// 	id    string
// 	cargo int
// }

// func (t *NormalTruck) LoadCargo() error {
// 	t.cargo += 2
// 	return nil
// }

// func (t *NormalTruck) UnloadCargo() error {
// 	t.cargo = 0
// 	return nil
// }

// type ElectricTruck struct {
// 	id      string
// 	cargo   int
// 	battery float64
// }

// func (e *ElectricTruck) LoadCargo() error {
// 	e.cargo += 1
// 	e.battery = -1

// 	return nil
// }

// func (e *ElectricTruck) UnloadCargo() error {
// 	e.cargo = 0
// 	e.battery += -1

// 	return nil
// }

// func processTruck(ctx context.Context, truck Truck) error {
// 	fmt.Printf("Processing Truck %+v \n", truck)

// 	return ErrTruckNotFound
// }

// func processFleet(ctx context.Context, trucks []Truck) error {
// 	var wg sync.WaitGroup

// 	errorsChannel := make(chan error, len(trucks))

// 	for _, t := range trucks {
// 		wg.Add(1)

// 		go func(t Truck) {
// 			err := processTruck(ctx, t)

// 			if err != nil {
// 				log.Println(err)
// 				errorsChannel <- err
// 			}

// 			wg.Done()
// 		}(t)
// 	}

// 	wg.Wait()
// 	close(errorsChannel)

// 	var errs []error

// 	for err := range errorsChannel {
// 		log.Printf("Error processing truck %v \n", err)
// 		errs = append(errs, err)
// 	}

// 	if len(errs) > 0 {
// 		return fmt.Errorf("fleet processing had %d errors", len(errs))
// 	}

// 	return nil
// }

func main() {
	m := make(map[string]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)

			// This might cause a race condition
			m[fmt.Sprintf("key-%d", i)] = i
		}(i)
	}

	wg.Wait()
	fmt.Println("Map:", m)
}
