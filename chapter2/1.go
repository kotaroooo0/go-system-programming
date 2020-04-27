package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%d\n%s\n%f\n", 1234, "5678", 9.0)
}
