package main

import "fmt"

func main() {
	// @NOTE: Maps are not concurrent safe

	// Creating a map
	m := make(map[string]int) // Key will be type of string, value will be type of int

	// Access value from map
	val, exists := m["a"] // here "a" is the key, if the "val" doesn't exist we will use "ok" to handle that

	if !exists {
		// handle it here
	} else {
		fmt.Printf("Val is %v", val)
	}

	// Or the inline syntax
	// if _, ok := m["a"]; ok {

	// }

	delete(m, "val") // Deletes by the key
	clear(m)         // Clears the whole map
}
