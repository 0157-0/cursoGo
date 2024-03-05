package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calc(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subs := n1 - n2

	return soma, subs
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto de f")
	fmt.Println(resultado)

	resultadoSum, resultadoSub := calc(5, 7)
	fmt.Println(resultadoSum, resultadoSub)

}
