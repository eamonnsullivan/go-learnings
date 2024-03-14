package main

import (
	"eamonn/filtering"
	"fmt"
)

func getName() string {
	return "Eamonn"
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("Hello, " + getName())
	filtering.PrintEvens(1, 100)
}
