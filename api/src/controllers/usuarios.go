package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario insere usuario no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var usuario models.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)

	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))

}

// BuscarUsuarios busca usuarios no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos os usuario"))
}

// BuscarUsuario busca um usuario no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar um usuario"))
}

// AtualizarUsuario altera usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando um usuario"))
}

// DeletarUsuario deleta usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletar um usuario"))
}
