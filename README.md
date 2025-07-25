# API de Tarefas em Go

Este é um projeto de uma API RESTful simples para gerenciamento de tarefas, desenvolvida em Go (Golang). É possível criar, listar, atualizar e deletar tarefas.

## Funcionalidades

- Criar uma nova tarefa
- Listar todas as tarefas
- Atualizar uma tarefa existente
- Deletar uma tarefa

## Estrutura do Projeto

## Rotas da API

| Método | Rota                 | Descrição                     |
| ------ | -------------------- | ----------------------------- |
| GET    | `/tarefas`           | Lista todas as tarefas        |
| POST   | `/tarefas/criar`     | Cria uma nova tarefa          |
| PUT    | `/tarefas/atualizar` | Atualiza uma tarefa existente |
| DELETE | `/tarefas/deletar`   | Deleta uma tarefa             |
