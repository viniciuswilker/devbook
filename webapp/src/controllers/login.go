package controllers

import (
	"net/http"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TELA DE LOGIN"))
}
