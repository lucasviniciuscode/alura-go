package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/alura-site-go/produtos-website/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade :=r.FormValue("quantidade")

		precoParsed, err := strconv.ParseFloat(preco, 64)
		quantidadeParsed, err := strconv.Atoi(quantidade)
		if(err != nil) {
			log.Println("Erro na conversão de dados: ", err)
		}

		models.InsertProduct(nome, descricao, precoParsed, quantidadeParsed)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")

	models.DeleteProduct(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request)  {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditProduct(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade :=r.FormValue("quantidade")

		precoParsed, err := strconv.ParseFloat(preco, 64)
		quantidadeParsed, err := strconv.Atoi(quantidade)
		if(err != nil) {
			log.Println("Erro na conversão de dados: ", err)
		}

		result := models.UpdateProduct(id, nome, descricao, precoParsed, quantidadeParsed)
		if result {
			fmt.Println("Update Success")
		}
	}
	http.Redirect(w, r, "/", 301)
}