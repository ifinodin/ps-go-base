package main

import "fmt"

func main() {
	const USDtoEUR = 0.86
	const USDtoRUB = 77.88
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
}
