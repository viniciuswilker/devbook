package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao

	publicacao.AutorID = usuarioID

	if erro := publicacao.Preparar(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)

	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, publicacao)

}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)

	publicacao, erro := repositorio.BuscarPorID(publicacaoId)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusOK, publicacao)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

func ExcluirPublicacao(w http.ResponseWriter, r *http.Request) {

}
