package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i enquanto for menor que 10")

	// 	time.Sleep(time.Second) // faz um aguarde de 1 segundo
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j enquanto for menor que 10", j)
	// 	time.Sleep(time.Second) // faz um aguarde de 1 segundo

	// }

	// // se quiser aumentar de 2 em 2 ou etc
	// for x := 0; x < 10; x+= 2 {
	// 	fmt.Println("Incrementando x enquanto for menor que 10", x)
	// 	time.Sleep(time.Second) // faz um aguarde de 1 segundo

	// }

	nomes := [3]string{"Joao", "davi", "lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
		// retorna 0 Joao | 1 Davi | 2 Lucas
	}
	// se vc n quiser o indice, é só usar _ exemplo: for _, nome .........

	for indice, letra := range "PALAVRAS" {
		fmt.Println(indice, string(letra))
		// retorna 0 P | 1 A | 3 L ......

	}

	usuario := map[string]string{
		"Nome":      "Leo",
		"sobrenome": "silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
		// retorna Nome Leo sobrenome silva

	}
}
