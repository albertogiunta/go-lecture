package main

import (
	"fmt"
	"math/rand"
	"time"
	"../google"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	start := time.Now()
	results, err := google.SearchParallel("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, err)
}
