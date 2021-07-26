package routes

import (
	"net/http"

	"github.com/sergio/Go-Lang-Web-/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) // toda vez que tiver uma "/" a função index vai atender

}
