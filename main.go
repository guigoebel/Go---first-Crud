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
	func rotaPrincipal(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Bem vindo")
	}

	func configurarRotas(){
		http.HandleFunc("/", rotaPrincipal)
	}

	func configurarServidor(){
		configurarRotas()

		fmt.Println("Servidor está rodando na porta 1337")
		http.ListenAndServe(":1337", nil) //defaultServerMux
	}


	func main(){
		configurarServidor()
	}