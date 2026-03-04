package main

import (
	"fmt"
	"log"
)

func main() {
	truckId := 42              // This is an integer
	anotherTruckId := &truckId // This is a pointer to an integer value

	fmt.Println("truckId", truckId)
	fmt.Println("truckId memory address", &truckId)

	fmt.Println("anotherTruckId", anotherTruckId)
	fmt.Println("anotherTruckId Value", *anotherTruckId) // Dereferencing the memory address to get the value

	truckId = 24
	fmt.Println("anotherTruckId Value after change", *anotherTruckId)

	TruckExample()

	var userID int
	log.Println(userID) // O/P -> 0

	// But if we do pointer reciever
	var userID2 *int
	log.Println(userID2) // O/P -> <nil> (Which is the default value for any variable without memory allocation)

}
