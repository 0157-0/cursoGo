package main

import "fmt"

func closure() func() {
	text := "Dentro da f closure"

	funcao := func() {
		fmt.Println(text)

	}
	return funcao
}
func main() {
	text := "Dentro da main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()

}
