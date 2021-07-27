package models

import (
	"fmt"
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
	db:= db.ConectaComBancoDeDados() // Abrindo a conexao com o banco de dados

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
		p.Id = id
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

//deleta os produtos do banco de dados
func DeletaProduto (id string){
	db := db.ConectaComBancoDeDados() // abrindo conexao com o banco

	deletarOProduto, err := db.Prepare("delete  from produtos where id=$1")
	if err != nil{
		panic(err.Error())
	}

	//deletando o produto executando pelo id
	deletarOProduto.Exec(id)

	defer  db.Close() // fechando conexao com o banco
}

//Edita os produtos do banco de dados
func EditaProduto(id string)Produtos {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produtos{}

	for produtoDoBanco.Next(){ // next prepara o resultados de uma linha e depois buscamos o resultado no scan

		var id, quantidade int
		var nome, descricao string
		var preco float64

		//Scan = diz q a variavel que criei acime tem q ser exatamente igual a do banco (usando &)
		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil{
			panic(err.Error())
		}
		//espelhando as variaveis com os dados do banco
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close() // fechando conexao com banco
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int)  {
	db := db.ConectaComBancoDeDados()

	fmt.Println(fmt.Sprintf("%s %s %s %s %s", id, nome,descricao,preco,quantidade))
	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4, where id=$5 ")
	if err != nil{
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
