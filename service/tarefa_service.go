package service

import (
	"github.com/jhonnydsl/Api-de-tarefas/database"
	"github.com/jhonnydsl/Api-de-tarefas/model"
)

func CriaTarefa(t model.Tarefa) error {
	return database.InsertTarefa(t)
}

func ListarTarefas() []model.Tarefa {
	return database.SelectTarefas()
}

func AtualizaTarefa(t model.Tarefa) error {
	return database.UpdateTarefa(t)
}

func DeletaTarefa(id int) error {
	return database.DeleteTarefa(id)
}