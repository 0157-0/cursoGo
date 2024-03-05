package main

import "fmt"

func main() {

	num := 10
	if num > 15 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	if outNum := num; outNum > 0 {
		fmt.Println("Maior")

	}
}
