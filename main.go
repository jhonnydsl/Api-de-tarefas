package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jhonnydsl/Api-de-tarefas/database"
	"github.com/jhonnydsl/Api-de-tarefas/rotas"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar variaveis de ambiente")
		return
	}
	
	err = database.Conexao()
	if err != nil {
		fmt.Println("Não foi possivel conectar ao banco de dados:", err)
		return
	}
	defer database.DB.Close()

	http.HandleFunc("/tarefas", rotas.RotaListaTarefa)
	http.HandleFunc("/tarefas/criar", rotas.RotaCriarTarefa)
	http.HandleFunc("/tarefas/atualizar", rotas.RotaAtualizaTarefa)
	http.HandleFunc("/tarefas/deletar", rotas.RotaDeletaTarefa)

	fmt.Println("Servidor rodando em http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}

