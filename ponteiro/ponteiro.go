package main

import "fmt"

func main() {

	var varivavel1 int = 10
	var varivavel2 int = varivavel1

	fmt.Println(varivavel1, varivavel2)
	// so incrementa o primeiro, pois o valor do 2 n é acessado, classico de ponteiro obrigado faculdade
	varivavel1++

	//ponteiro é ref de memoria
	var varivavel3 int
	var ponteiro *int

	// dessa forma pega o endereço de memoria onde ta salvo o valor 100 no ponteiro
	varivavel3 = 100
	ponteiro = &varivavel3

	// para pegar o valor salvo nesse endereço de memoria, precisa colocar * na frente da variavel que é o ponteiro
	fmt.Println(varivavel3, *ponteiro)

}
