//go:build !darwin

package main

import "fmt"

func main() {
	fmt.Println("This Star Wars calculator GUI runs on macOS only. Build and run on macOS: go run .")
}
