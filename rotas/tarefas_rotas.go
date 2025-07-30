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
	err := service.CriaTarefa(novaTarefa)
	if err != nil {
		http.Error(w, "Erro ao adicionar tarefa", http.StatusInternalServerError)
		return
	}
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
		return
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
	if tarefaAtualizada.Id == 0 {
		http.Error(w, "Id invalido para atualização.", http.StatusBadRequest)
		return
	}
	err := service.AtualizaTarefa(tarefaAtualizada)
	if err != nil {
		http.Error(w, "Erro ao atualizar tarefa", http.StatusInternalServerError)
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
	err := service.DeletaTarefa(int(tarefaParaExcluir.Id))
	if err != nil {
		http.Error(w, "Erro ao excluir tarefa", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	listaDeTarefas := service.ListarTarefas()
	err = json.NewEncoder(w).Encode(listaDeTarefas)
	if err != nil {
		http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
	}
}