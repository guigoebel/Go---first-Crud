package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

var Livros []Livro = []Livro{
	Livro{
		Id:     1,
		Titulo: "O Guarani",
		Autor:  "Jose de Alencar",
	},
	Livro{
		Id:     2,
		Titulo: "Cazuza",
		Autor:  "Viriato Correia",
	},
	Livro{
		Id:     3,
		Titulo: "Dom Casmurro",
		Autor:  "Machado de Assis",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var novoLivro Livro

	json.Unmarshal(body, &novoLivro)
	novoLivro.Id = len(Livros) + 1

	Livros = append(Livros, novoLivro)

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(novoLivro)
}

func excluirLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, erro := strconv.Atoi(vars["livroId"])

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	indiceLivro := -1
	for indice, livro := range Livros {
		if livro.Id == id {
			indiceLivro = indice
			break
		}
	}

	if indiceLivro < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Livros = append(Livros[0:indiceLivro], Livros[indiceLivro+1:len(Livros)]...)

	w.WriteHeader(http.StatusNoContent)
}

func modificarLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, erro := strconv.Atoi(vars["livroId"])

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	corpo, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var livroModificado Livro
	erroJson := json.Unmarshal(corpo, &livroModificado)

	if erroJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	indiceLivro := -1
	for indice, livro := range Livros {
		if livro.Id == id {
			indiceLivro = indice
			break
		}
	}

	if indiceLivro < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Livros[indiceLivro] = livroModificado
	Livros[indiceLivro].Id = indiceLivro + 1

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(livroModificado)

}

func buscarLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["livroId"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, livro := range Livros {
		if livro.Id == id {
			json.NewEncoder(w).Encode(livro)
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

func configurarRotas(roteador *mux.Router) {
	roteador.HandleFunc("/", rotaPrincipal)
	roteador.HandleFunc("/livros", listarLivros).Methods("GET")
	roteador.HandleFunc("/livros", cadastrarLivro).Methods("POST")
	roteador.HandleFunc("/livros/{livroId}", buscarLivro).Methods("GET")
	roteador.HandleFunc("/livros/{livroId}", modificarLivro).Methods("PUT")
	roteador.HandleFunc("/livros/{livroId}", excluirLivro).Methods("DELETE")
}

func configurarServidor() {
	roteador := mux.NewRouter().StrictSlash(true)
	roteador.Use(jsonMiddleware)
	configurarRotas(roteador)

	fmt.Println("Servidor está rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", roteador)) //Gorilla mux
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func main() {
	configurarServidor()
}