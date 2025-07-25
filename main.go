package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jhonnydsl/Api-de-tarefas/rotas"
)

func main() {
	http.HandleFunc("/tarefas", rotas.RotaListaTarefa)
	http.HandleFunc("/tarefas/criar", rotas.RotaCriarTarefa)
	http.HandleFunc("/tarefas/atualizar", rotas.RotaAtualizaTarefa)
	http.HandleFunc("/tarefas/deletar", rotas.RotaDeletaTarefa)

	fmt.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

