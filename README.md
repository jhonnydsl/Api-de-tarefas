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

## Como rodar

1. Configure seu `.env` com os dados do seu banco.
2. Execute `go run main.go` para iniciar o servidor.
3. Acesse `http://localhost:8080` para usar a API.

## Configuração do Banco de Dados

Este projeto utiliza variáveis de ambiente para configurar a conexão com o banco de dados. Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_HOST=localhost
DB_PORT=3306
DB_NAME=nome_do_banco
