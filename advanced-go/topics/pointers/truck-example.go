package main

import "fmt"

type NormalTruck struct {
	id    string
	cargo int
}

func TruckExample() {
	t := NormalTruck{cargo: 0}

	fillTruckCargo(&t)

	fmt.Println(t)
}

func fillTruckCargo(t *NormalTruck) {
	t.cargo = 100
}
