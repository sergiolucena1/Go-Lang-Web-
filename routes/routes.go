package routes

import (
	"net/http"

	"github.com/sergio/Go-Lang-Web-/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) // toda vez que tiver uma "/" a função index vai atender
	http.HandleFunc("/new", controllers.New) // func New
	http.HandleFunc("/insert", controllers.Insert)// func Insert(inserindo os dados )
	http.HandleFunc("/delete", controllers.Delete)// func Delete
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update",controllers.Update)
}
