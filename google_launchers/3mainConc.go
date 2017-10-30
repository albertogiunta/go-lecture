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
	results, err := google.SearchTimeout("golang", 102*time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, err)
}
