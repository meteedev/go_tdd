package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayers(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/players", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	res := httptest.NewRecorder()
	PlayerServer(res, req)

	t.Run("returns Mr A 's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/A", nil)
		response := httptest.NewRecorder()
	
		PlayerServer(response, request)
	
		got := response.Body.String()
		want := "20"
	
		assert.Equal(t,want,got)
	})



	t.Run("returns B's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/B", nil)
		response := httptest.NewRecorder()
	
		PlayerServer(response, request)
	
		got := response.Body.String()
		want := "10"
	
		assert.Equal(t,want,got)

	})

}
