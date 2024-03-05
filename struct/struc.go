package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logadouro string
	numero    uint8
}

func main() {
	fmt.Println("main main")

	var u usuario
	u.nome = "teste"
	u.idade = 21
	u.endereco.logadouro = "teste r"
	u.endereco.numero = 6
	fmt.Println(u)

	enderecoEx := endereco{"Rua tal", 5}
	u2 := usuario{"usuario2", 21, enderecoEx}
	fmt.Println(u2)

	//caso queira iniciar só com um dado
	u3 := usuario{idade: 50}
	fmt.Println(u3)

}
