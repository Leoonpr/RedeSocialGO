package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config.Carregar()
	// fmt.Println(config.StringConexaoBanco)
	// fmt.Println(config.Porta) 
	fmt.Println("API rodando no servidor")
	r := router.Gerar()

	portaString := ":" + strconv.Itoa(config.Porta)
	log.Fatal(http.ListenAndServe(portaString, r))
}
