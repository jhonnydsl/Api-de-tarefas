package service

import (
	"github.com/jhonnydsl/Api-de-tarefas/model"
)

var Tarefas = make(map[int]model.Tarefa)
var contador int

// CriaTarefa adiciona uma nova tarefa no mapa Tarefas, atribuindo um ID automático sequencial.
func CriaTarefa(t model.Tarefa) {
	novaTarefa := model.Tarefa {
		Nome: t.Nome,
		Id: contador,
		Concluido: t.Concluido,
	}

	Tarefas[contador] = novaTarefa
	contador ++
}

// ListarTarefas retorna uma cópia atualizada de todas as tarefas armazenadas.
func ListarTarefas() map[int]model.Tarefa {
	var lista = make(map[int]model.Tarefa)

	for _, valor := range Tarefas {
		lista[valor.Id] = valor	
	}
	return lista
}

// AtualizaTarefa atualiza uma tarefa existente com base no ID.
// Retorna false se a tarefa não existir.
func AtualizaTarefa(tarefa model.Tarefa) bool {
	if _, existe := Tarefas[tarefa.Id]; !existe {
		return false
	}
	Tarefas[tarefa.Id] = tarefa
	return true
}

// DeletaTarefa remove a tarefa pelo ID.
// Retorna false se a tarefa não existir.
func DeletaTarefa(id int) bool {
	if _, existe := Tarefas[id]; !existe {
		return false
	}
	delete(Tarefas, id)
	return true
}