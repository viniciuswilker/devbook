package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
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
