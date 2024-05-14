package main

import (
	"log"
	"net/http"

	"github.com/meteedev/go_tdd/http-server/player/server"
	"github.com/meteedev/go_tdd/http-server/player/store"
)

func main() {
	server := &server.PlayerServer{Store: &store.InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
