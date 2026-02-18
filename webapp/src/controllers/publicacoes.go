package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Titulo   string `json:"titulo"`
		Conteudo string `json:"conteudo"`
	}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	publicacao, err := json.Marshal(p)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(
		r,
		http.MethodPost,
		url,
		bytes.NewBuffer(publicacao),
	)

	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		body, _ := io.ReadAll(response.Body)
		log.Println("Erro API:", string(body))

		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.APIURL, publicacaoId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {

	fmt.Println("CHEGOU AQUI")

	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/descurtir", config.APIURL, publicacaoId)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	var p struct {
		Titulo   string `json:"titulo"`
		Conteudo string `json:"conteudo"`
	}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	publicacao, err := json.Marshal(p)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoId)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
