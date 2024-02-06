// Package example demostraits a larger test function.
package example

import (
	"fmt"
	"log"
)

// Process returs known changes to its parameters for testing.
func Process(factor int, msg string) (int, string, float64) {
	const mulFactor = 2
	const aThird = 1.0 / 3.0
	// Function being tested.
	log.Printf("Entered process(%d, %q)", factor, msg)

	if factor < 1 || factor > 10 {
		log.Panicf("factor (%d) out of bounds: %q", factor, msg)
	}

	fmt.Println("Processing with factor:", factor, "and message:", msg)

	return factor * mulFactor, "Processed: " + msg, float64(factor) * aThird
}
