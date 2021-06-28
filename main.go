package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

	type Livro struct {
		Id int
		Titulo string
		Autor string
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

	func configurarRotas(){
		http.HandleFunc("/", rotaPrincipal)
		http.HandleFunc("/livros", listarLivros)
	}

	func configurarServidor(){
		configurarRotas()

		fmt.Println("Servidor está rodando na porta 1337")
		http.ListenAndServe(":1337", nil) //defaultServerMux
	}


	func main(){
		configurarServidor()
	}