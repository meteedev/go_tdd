package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// server_test.go
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	//fmt.Println("in stub")
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayers(t *testing.T) {

	store := StubPlayerStore{
		scores: map[string]int{
			"A": 20,
			"B": 10,
		},
	}

	//init server
	server := &PlayerServer{Store: &store}
	t.Run("returns Mr A 's score", func(t *testing.T) {
		req, err := NewGetRequestScore("A")
		if err != nil {
			fmt.Println(err.Error())
		}
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Body.String()
		want := "20"

		assert.Equal(t, want, got)
	})

	t.Run("returns B's score", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/players/B", nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Body.String()
		want := "10"

		assert.Equal(t, want, got)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/players/C", nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusNotFound

		assert.Equal(t, want, got)
	})

}

func TestStoreScore(t *testing.T) {

	store := StubPlayerStore{
		scores:   map[string]int{},
		winCalls: nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {

		player := "D"

		req, err := newPostWinRequest("D")
		if err != nil {
			fmt.Println(err.Error())
		}
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusAccepted
		assert.Equal(t, want, got)

		assert.Equal(t, 1, len(store.winCalls))
		assert.Equal(t, player, store.winCalls[0])

	})
}

func NewGetRequestScore(name string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return req, nil
}

func newPostWinRequest(name string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return req, nil
}
