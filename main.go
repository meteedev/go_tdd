package main

import (
	"log"
	"net/http"

	"github.com/meteedev/go_tdd/http-server/player/server"
)

func main(){
	server := &server.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000",server))
}
