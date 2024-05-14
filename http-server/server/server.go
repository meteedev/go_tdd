package server

import (
	"fmt"
	"net/http"
	"strings"
)


func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string)(score string){
	if name == "A" {
		return	"20"
	}

	if name == "B" {
		return "10"
	}

	return ""
}