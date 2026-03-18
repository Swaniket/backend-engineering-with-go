package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

// @IMPORTANT: maps in go are not concurrent safe
type truckManager struct {
	trucks map[string]*Truck
	// mu     sync.RWMutex // Mutex allows us to block reading operation to a key when it's being modified / written
	sync.RWMutex // We can also apply mutex without adding a variable and embedding directly into a struct - its called COMPOSITION
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (m *truckManager) AddTruck(id string, cargo int) error {
	// .Lock(), locks the map so that until the operation is completed, the other go routines can't access this.
	m.Lock()         // Because of COMPOSITION, we directly have access to the mutex functions inside of 'm'
	defer m.Unlock() // Unclock at the end of the func execution

	m.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (m *truckManager) GetTruck(id string) (Truck, error) {
	m.RLock() // .RLock() only locks the read operation
	defer m.RUnlock()

	truck, ok := m.trucks[id]

	if !ok {
		return Truck{}, ErrTruckNotFound
	}

	return *truck, nil
}

func (m *truckManager) RemoveTruck(id string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.trucks, id)
	return nil
}

func (m *truckManager) UpdateTruckCargo(id string, cargo int) error {
	m.Lock()
	defer m.Unlock()

	m.trucks[id].Cargo = cargo
	return nil
}
