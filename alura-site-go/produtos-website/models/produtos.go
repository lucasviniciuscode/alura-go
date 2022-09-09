package models

import "github.com/alura-site-go/produtos-website/db"


type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func InsertProduct(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	InsereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	InsereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConectaComBancoDeDados()
	DeleteProduct, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}
	DeleteProduct.Exec(id)
	defer db.Close()
}

func UpdateProduct(id string, nome string, descricao string, preco float64, quantidade int) bool {
	db := db.ConectaComBancoDeDados()
	UpdateDados, err := db.Prepare("update produtos set nome = ?, descricao = ?, preco = ?, quantidade = ? where id = ? ")
	if err != nil {
		panic(err.Error())
	}
	UpdateDados.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
	return true
}

func EditProduct(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoDb, err := db.Query("select * from produtos where id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	produto := Produto{}

	for produtoDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()
	return produto
}
