package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

	type Livro struct {
		Id int `json:"id"`
		Titulo string `json:"titulo"`
		Autor string `json:"autor"`
	}

	var Livros []Livro = []Livro{
		Livro{
			Id: 1,
			Titulo: "O Guarani",
			Autor: "Jose de Alencar",
		},
		Livro{
			Id: 2,
			Titulo: "Cazuza",
			Autor: "Viriato Correia",
		},
		Livro{
			Id: 3,
			Titulo: "Dom Casmurro",
			Autor: "Machado de Assis",
		},
	}

	func rotaPrincipal(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Bem vindo")
	}

	func listarLivros(w http.ResponseWriter, r *http.Request){
		encoder := json.NewEncoder(w)
		encoder.Encode(Livros)
	}

	func cadastrarLivro(w http.ResponseWriter, r *http.Request){

		body, erro := ioutil.ReadAll(r.Body)

		if erro != nil{
			//TODO lidar com erro
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

	func excluirLivro (w http.ResponseWriter, r *http.Request){
		partes := strings.Split(r.URL.Path, "/")
		id, _ := strconv.Atoi(partes[2])

		indiceLivro := -1
		for indice, livro := range Livros {
			if livro.Id == id {
				indiceLivro = indice
				break
			}
		}

		if (indiceLivro < 0) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		Livros = append(Livros[0:indiceLivro], Livros[indiceLivro + 1: len(Livros)]...)

		w.WriteHeader(http.StatusNoContent)
	}

	func rotearLivros(w http.ResponseWriter, r *http.Request){
		//ta mto feio isso, existe algum método melhor?
		//livros
		w.Header().Set("Content-Type", "application/json")
		partes := strings.Split(r.URL.Path, "/")

		if len(partes) == 2 || len(partes) == 3 && partes[2] == "" {
			if r.Method == "GET" {
				listarLivros(w, r)
			} else if r.Method == "POST" {
				cadastrarLivro(w, r)
			}
		} else if len(partes) == 3 || len(partes) == 4 && partes[3] == ""{
			if r.Method == "GET" {
				buscarLivro(w, r)
			}else if r.Method == "DELETE"{
				excluirLivro(w, r)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	}

	func buscarLivro(w http.ResponseWriter, r *http.Request){
		partes := strings.Split(r.URL.Path, "/")

		id, erro := strconv.Atoi(partes[2])
		//TODO error handler
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, livro := range Livros {
			if livro.Id == id {
				json.NewEncoder(w).Encode(livro)
			}
		}
		w.WriteHeader(http.StatusNotFound)

	}

	func configurarRotas(){
		http.HandleFunc("/", rotaPrincipal)
		http.HandleFunc("/livros", rotearLivros)
		http.HandleFunc("/livros/", rotearLivros)
	}

	func configurarServidor(){
		configurarRotas()
		//TODO if error dont show this message.
		fmt.Println("Servidor está rodando na porta 1337")
		log.Fatal(http.ListenAndServe(":1337", nil)) //defaultServerMux
	}


	func main(){
		configurarServidor()
	}