package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a f1")

}
func funcao2() {
	fmt.Println("Executando a f2")

}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("media calculada")

	fmt.Println("na f")

	media := (n1 + n2) / 2

	if media >= 6 {

		return true
	}
	return false

}
func main() {
	// defer adia para a ultima função a ser executada
	funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))

}
