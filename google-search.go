package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/llitfkitfk/demo/pkg/google"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	start := time.Now()
	results, err := google.Search("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, err)
}