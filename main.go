package main

	import (
		"fmt"
		"net/http"
	)

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