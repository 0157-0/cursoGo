package main

import (
	"errors"
	"fmt"
)

// tipos de dados
func main() {
	var number int64 = 1000
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	// alias
	// int32 = rune (q usa caracteres)
	var number3 rune = 123456
	fmt.Println(number3)

	// byte = uint8
	var number4 byte = 123
	fmt.Println(number4)

	//numeros reais
	var numberReal float32 = 123.45
	fmt.Println(numberReal)

	var numberReal2 float64 = 123456789.45
	fmt.Println(numberReal2)

	//inicia no valor zero ou vazio
	var text string
	fmt.Println(text)

	// se quiser iniciar em false, é só não passar nada
	var bool1 bool = true
	fmt.Println(bool1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
