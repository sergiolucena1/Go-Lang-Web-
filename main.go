package main

import (
	"log"
	"net/http"

	"github.com/sergio/Go-Lang-Web-/routes"
)

func main() {
	routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":3000", nil)) // porta do servidor
}
