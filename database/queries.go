package database

import (
	"fmt"

	"github.com/jhonnydsl/Api-de-tarefas/model"
)

// Função que envia um insert e adiciona uma tarefa no banco de dados.
func InsertTarefa(t model.Tarefa) error {
	smt, err := DB.Prepare("insert into tarefas(titulo, concluida) values (?, ?)")
	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.Exec(t.Titulo, t.Concluida)
	if err != nil {
		return err
	}

	return nil
}

// Função que envia um select para o banco de dados e retorna todas as tarefas armazenadas.
func SelectTarefas() []model.Tarefa {
	rows, err := DB.Query("select id, titulo, concluida from tarefas")
	if err != nil {
		fmt.Println("Erro no select", err)
		return nil
	}
	defer rows.Close()
	
	var lista []model.Tarefa
	for rows.Next() {
		var t model.Tarefa
		if err = rows.Scan(&t.Id, &t.Titulo, &t.Concluida); err != nil {
			fmt.Println("Erro no scan", err)
			continue
		}
		lista = append(lista, t)
	}
	return lista
}

// Função que envia um update e atualiza uma tarefa no banco de dados.
func UpdateTarefa(t model.Tarefa) error {
	smt, err := DB.Prepare("update tarefas set titulo = ?, concluida = ? where id = ?")
	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.Exec(t.Titulo, t.Concluida, t.Id)
	if err != nil {
		return err
	}
	return nil
}

// Função que envia um delete e exclui uma tarefa do banco de dados.
func DeleteTarefa(id int) error {
	smt, err := DB.Prepare("delete from tarefas where id = ?")
	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}