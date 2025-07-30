package model

type Tarefa struct {
	Titulo    string `json:"titulo"`
	Id        int64  `json:"id"`
	Concluida bool   `json:"concluida"`
}