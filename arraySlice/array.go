package main

import "fmt"

func main() {
	var array1 [5]string
	// adicionando valor em determinada posição no array
	array1[0] = "pos1"
	fmt.Println(array1)

	array2 := [5]string{"pors1", "pos2", "pos3", "pos4", "pos5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slice, é mais usado pq o tamanho n é fixo, diferente dos array

	slice := []int{10, 11, 12, 13, 14}
	fmt.Println(slice)

	// adicionando valor no slice
	slice = append(slice, 18)
	fmt.Println(slice)

	// pegando a possicao 1 ao 2 do array e salvando na variavel slice2
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// se alterar no array, atualiza no slice tmb
	array2[1] = "posAlt"
	fmt.Println(slice2)

	// arrays internos
	//make cria um array de 11 capacidade, mas retorna somente as 10 primeiras posições
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade
	// se passar de 11 q é capacidade, ele dobra o valor
	slice3 = append(slice3, 5)
	fmt.Println(slice3)

}
