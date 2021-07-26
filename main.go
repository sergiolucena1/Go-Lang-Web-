package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=sergio_loja password=Sergim25. host=localhost sslmode=disable" // as informações que tem q ser passadas
	db, err := sql.Open("postgres", conexao)                                                        //abrindo uma conexao com o banco de dados

	//verificando se tem algum erro
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produtos struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

//template.Must (encapsula todas as nossas templates e devolve 2 retornos(o template e o erro))
//template.ParseGlob (Passa o caminho dos templates)"nome da pata /* tudo que for html"

func main() {
	db := conectaComBancoDeDados()
	defer db.Close() // denfer: ele executa depois de tudo para fechar a execucao
	http.HandleFunc("/", index)                  // toda vez que tiver uma "/" a função index vai atender
	log.Fatal(http.ListenAndServe(":3000", nil)) // porta do servidor
}

//Parametros(quem vai escrever, quem vai exibir )
func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produtos{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tênis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto Novo", "Muito legal", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
