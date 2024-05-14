package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)


//server_test.go
type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	//fmt.Println("in stub")
	score := s.scores[name]
	return score
}


func TestGetPlayers(t *testing.T) {
	
	store := StubPlayerStore{
		map[string]int{
			"A": 20,
			"B":  10,
		},
	}

	//init server
	server := &PlayerServer{store: &store}
	t.Run("returns Mr A 's score", func(t *testing.T) {
		req, err := NewGetRequestScore("A")
		if err != nil {
			fmt.Println(err.Error())
		}
		res := httptest.NewRecorder()
	
		server.ServeHTTP(res,req)
	
		got := res.Body.String()
		want := "20"
		
		assert.Equal(t,want,got)
	})



	t.Run("returns B's score", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/players/B", nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		res := httptest.NewRecorder()
	
		server.ServeHTTP(res,req)
	
		got := res.Body.String()
		want := "10"
	
		assert.Equal(t,want,got)
	})

}


func NewGetRequestScore(name string)(*http.Request,error){
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil , err
	}
	return req,nil
}



