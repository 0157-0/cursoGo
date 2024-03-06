package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Exec recuperada com sucesso")

	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperar()

	media := (n1 + n2) / 2

	if media > 6 {

		return true
	} else if media < 6 {
		return false
	}
	panic("a média é exatamente 6")

}

// o panic faz o sistema parar, o panic chama todas as funções, mesmo com defer
// o recover faz o sistema voltar
func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós exec")

}
