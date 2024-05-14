package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/meteedev/go_tdd/http-server/player/store"
)

type PlayerServer struct {
	Store store.PlayerStore
}

func (p PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.Store.GetPlayerScore(player))
}
