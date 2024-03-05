package main

import "fmt"

func main() {
	//tipando variaveis, todos precisam ser tipadas
	var variavel1 string = "Variavel 1"
	// deixando implicito
	variavel2 := "Variavel 2"

	fmt.Println(variavel1, variavel2)
	// declarando varias
	var (
		variavel3 string = "aaa"
		variavel4 string = "bbb"
	)
	fmt.Println(variavel3, variavel4)

	// declarando varias implicito

	variavel5, variavel6 := "Variavel5", "Variavel6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

}
