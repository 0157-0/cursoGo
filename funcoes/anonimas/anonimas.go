package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("passei o texto")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("aaaaaa")
	fmt.Println(retorno)

}
