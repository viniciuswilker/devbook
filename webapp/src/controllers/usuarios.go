package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha")})
	if erro != nil {
		log.Fatal()
	}
	// fmt.Println(bytes.NewBuffer(usuario))

	response, erro := http.Post("http://localhost:5000/usuarios", "apliccation/json", bytes.NewBuffer(usuario))

	if erro != nil {
		log.Fatal()
	}

	defer response.Body.Close()

	fmt.Println(response.Body)
}
