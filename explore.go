// Create and seed the generator.
// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
// Using a fixed seed will produce the same output on every run.
// r := rand.New(rand.NewSource(99))

import (
	"math/rand"
	"time"
)

func seedRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
