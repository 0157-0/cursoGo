package main

import (
	"log"
	"net/http"
)

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	// HandleFunc - recebe dois parametros, nome da rota e a função q lida com ela
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // raiz
		w.Write([]byte("Página Raiz!"))
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	// pode fazer essa funcao fora e chamar ela na rota
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
