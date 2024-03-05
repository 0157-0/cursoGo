package main

import "fmt"

func main() {

	// fora do colchete é tipo dos valores
	// dentro do colchete é o tipo das chaves
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	// para acessar as propriedades
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"ultimo":   "silva",
		},
		"curso": {
			"titulo": "engenharia",
			"campus": "campus 1",
		},
	}

	fmt.Println(usuario2)
	// deletando uma chave do map
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	// add chave
	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}
	fmt.Println(usuario2)

}
