package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"Jota", 35, 170, "Joo"}
	fmt.Println(p1)
	el := estudante{p1, "medicina", "usp"}
	fmt.Println(el)
	// caso queira pegar os dados do estudante, sem pela pessoa
	fmt.Println(el.nome)

}
