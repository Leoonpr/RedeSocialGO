package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeUsuarios(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "home.html", nil)

}
