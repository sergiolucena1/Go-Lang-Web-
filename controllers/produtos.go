package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/sergio/Go-Lang-Web-/models"
)

//variavel com todos os nossos templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

//template.Must (encapsula todas as nossas templates e devolve 2 retornos(o template e o erro))
//template.ParseGlob (Passa o caminho dos templates)"nome da pata /* tudo que for html"

//Parametros(quem vai escrever, quem vai exibir )
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	//EXECUTA NO TEMPLATE
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

//executando o template
func New(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "New", nil)
}

//Buscando, convertendo e inserindo os dados
func Insert (w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{ // se a requisição for do metodo POST, criar um novo produto
		//buscando os dados
		nome:= r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")


	//convertendo os valores necessarios(os valores que vem pra nós sao do tipo string) temos q converter;
		//convertendo string para float64
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		//convertendo string par int
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		//Pegar os valores e criar (funcao do models)
		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	//depois de passar os dados redirecionar para a pagina inicial
	http.Redirect(w, r, "/", 301)
	return
}

//Deletando dados do banco
func Delete(w http.ResponseWriter, r *http.Request){

	//pegando a url atraves da requisiçãp (.Get = pegar )
	idDoProduto := r.URL.Query().Get("id") // pegando o id da url

	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

//editando os dados do banco
func Edit(w http.ResponseWriter, r *http.Request){
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

// atualizando os dados do banco
func Update (w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{ //  retorna tudo em string
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//convertendo os valores

		idConvertidoParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("erro na conversão do ID para int:", err)
		}

		precoConvertidoParaFloat , err := strconv.ParseFloat(preco,64)
		if err != nil {
			log.Println("erro na conversão do preço para float64:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("erro na conversão da quantidade para int:", err)
		}

		//Mostrando quais os valores pra atualizar
		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)

}