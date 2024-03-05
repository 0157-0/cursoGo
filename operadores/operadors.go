package main

import "fmt"

func main() {
	soma := 1 + 2
	subs := 1 - 2
	div := 10 / 2
	mult := 10 * 5
	restDiv := 10 % 2

	fmt.Println(soma, subs, div, mult, restDiv)

	// só podem fazer conta se forem do msm tipo, se fosse int8, int32 etc, ia da erro
	var num1 int16 = 10
	var num2 int16 = 25
	soma2 := num1 + num2
	fmt.Println(soma2)

	// op relacional
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//op logicos

	// && é "e"
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	// ou padrão
	fmt.Println(verdadeiro || falso)
	// inverter o booll
	fmt.Println(!verdadeiro) // a saida vai vir false

	// op unarios (serve para incrementar e algo do tipo)
	numero := 10
	numero++     //ai vai sair 11
	numero += 15 //ai vai sair 25, ele vai somar os 15
	numero--     // ai diminui um e assim vai

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)

	// operadores ternarios n existem, tipo texto > 5 ? "aaaa" : "bbb"
}
