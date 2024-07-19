package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()
	fmt.Println("Rodando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(":3000", r))
}
