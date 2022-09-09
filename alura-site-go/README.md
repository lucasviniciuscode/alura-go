# alura-site-go

Projeto realizado em paralelo ao curso da Alura [Go: Crie uma aplicação web](https://cursos.alura.com.br/course/go-lang-web)

## Objetivo
Estudar e despertar curiosidade sobre a Golang

## Desenvolvido
Foi desenvolvido uma aplicação Web utilizando Html, Bootstrap, e Goolang e Mysql que faz um crud na tabela `produtos`.

## Mysql
```sql
CREATE DATABASE alura_site;
USE alura_site;
CREATE TABLE produtos (id int NOT NULL PRIMARY KEY,nome varchar(255) NOT NULL, descricao varchar(255),preco double, quantidade int);
```

![Captura de tela de 2022-08-24 16-21-43](https://user-images.githubusercontent.com/92001463/186505215-6dd06479-c39a-4d71-a0ce-dc0ec37e80ed.png)
