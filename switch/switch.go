package main

import "fmt"

// tem as duas maneiras de usar
func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	default:
		return "Numero invalido"
	}
}

func diaDaSeman2a(num int) string {
	var dia string
	switch {
	case num == 1:
		dia = "Domingo"
	case num == 2:
		dia = "Segunda"
	default:
		dia = "Numero invalido"
	}
	return dia
}

func diaDaSemana3(num int) string {
	var dia string
	switch {
	case num == 1:
		dia = "Domingo"
		fallthrough // faz jogar o valor para o proximo, logo se eu passar 1, a respsota vem segunda
	case num == 2:
		dia = "Segunda"
	default:
		dia = "Numero invalido"
	}
	return dia
}

func main() {
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSeman2a(1))
	fmt.Println(diaDaSemana3(1))
}
