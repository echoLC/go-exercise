package main

import (
	"fmt"

	"go-exercise/exercise"
)

func main() {
	fmt.Printf("%s", exercise.ToRNA("A"))
	val, _ := exercise.Square(2)
	fmt.Printf("%d", val)
}