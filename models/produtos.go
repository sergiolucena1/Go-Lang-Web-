package models

import (
	"github.com/sergio/Go-Lang-Web-/db"
)

//criando a struct
type Produtos struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//pegando os dados do banco
func BuscaTodosOsProdutos() []Produtos {
	db := db.ConectaComBancoDeDados() // Abrindo a conexao com o banco de dados

	//selecionando do banco de dados
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produtos{} // p = armazena todo o produto q vem do banco
	produtos := []Produtos{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade) // scanear linha a linha do banco de dados
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		//criando uma lista
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

// criando novos dados no banco
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int){
	db := db.ConectaComBancoDeDados() // abrindo a conexao com o banco

							//db.Prepare = preparando o banco de dados passando a query
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
// se nao tiver nenhum erro inserir os dados no banco(.Exec)

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
//depois de inserir os dados, eu fecho o banco de dados...
	defer db.Close()

}