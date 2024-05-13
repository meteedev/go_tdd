package main

import (
	"log"
	"net/http"

	"github.com/meteedev/go_tdd/http-server/server"
)

func main(){
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000",handler))
}
