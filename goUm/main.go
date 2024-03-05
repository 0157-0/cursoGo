package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo no main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devtestegmail.com")
	fmt.Println(erro)
}
