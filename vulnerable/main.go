package main

import (
    "fmt"
    "net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

    // искусственно вызываем панику он покажет стек в ответе/логах?
    panic("some internal error happened")
}

func configHandler(w http.ResponseWriter, r *http.Request) {
    //secret is returned directly its bad 
    fmt.Fprintf(w, "API Key: %s", APIKey)
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/config", configHandler)
    http.HandleFunc("/create-user", createUser) // из deserialize.go

    fmt.Println("[vulnerable] Listening on :8080")


    http.ListenAndServe(":8080", nil)     //  server 
}


