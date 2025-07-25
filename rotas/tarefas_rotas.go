package rotas

import (
	"encoding/json"
	"net/http"

	"github.com/jhonnydsl/Api-de-tarefas/model"
	"github.com/jhonnydsl/Api-de-tarefas/service"
)

// Metodo para validar se a requisição é valida.
func MetodoValido(w http.ResponseWriter, r *http.Request, metodo string) bool {
	if r.Method != metodo {
		http.Error(w, "Método não permitido :(", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

// Metodo para validar se o JSON foi decodificado com sucesso.
func DecodificaJSON(w http.ResponseWriter, r *http.Request, destino interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(destino)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return false
	}
	return true
}

func RotaCriarTarefa(w http.ResponseWriter, r *http.Request) {
	if !MetodoValido(w, r, "POST") {
		return
	}
	var novaTarefa model.Tarefa
	
	if !DecodificaJSON(w, r, &novaTarefa) {
		return
	}

	service.CriaTarefa(novaTarefa)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)   // => Retorna para o usuario um status 201 dizendo que a tarefa foi criada com sucesso.
	json.NewEncoder(w).Encode(novaTarefa)    // => Envia para o usuario o JSON da tarefa criada.
}

func RotaListaTarefa( w http.ResponseWriter, r * http.Request) {
	if !MetodoValido(w, r, "GET") {
		return
	}

	listaDeTarefas := service.ListarTarefas()
	w.Header().Set("Content-Type", "application/json")   // => Define o tipo do cabeçalho que sera enviado como resposta.
	err := json.NewEncoder(w).Encode(listaDeTarefas)
	if err != nil {
		http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
	}
}

func RotaAtualizaTarefa(w http.ResponseWriter, r *http.Request) {
	if !MetodoValido(w, r, "PUT") {
		return
	}

	var tarefaAtualizada model.Tarefa
	if !DecodificaJSON(w, r, &tarefaAtualizada) {
		return
	}
	ok := service.AtualizaTarefa(tarefaAtualizada)
	if !ok {															// => Verifica se a tarefa que o usuario deseja atualizar realmente existe.
		http.Error(w, "Tarefa não encontrada.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tarefaAtualizada)
}

func RotaDeletaTarefa(w http.ResponseWriter, r *http.Request) {
	if !MetodoValido(w, r, "DELETE") {
		return
	}

	var tarefaParaExcluir model.Tarefa
	if !DecodificaJSON(w, r, &tarefaParaExcluir) {
		return
	}
	ok := service.DeletaTarefa(tarefaParaExcluir.Id)
	if !ok {
		http.Error(w, "Tarefa não encontrada para ser excluída", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	listaDeTarefas := service.ListarTarefas()
	err2 := json.NewEncoder(w).Encode(listaDeTarefas)
	if err2 != nil {
		http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
	}
}