package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/response"
	"api/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := autenticacao.GerarToken(usuarioSalvoNoBanco.ID)

	fmt.Println(token)
	w.Write([]byte("você está logado!"))

}
