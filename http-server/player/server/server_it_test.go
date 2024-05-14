// server_integration_test.go
package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meteedev/go_tdd/http-server/player/store"
	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	//server_integration_test.go
	store := store.NewInMemoryPlayerStore()
	server := PlayerServer{store}
	
	player := "Lieang"

	req1 , err := newPostWinRequest(player)
	if err != nil {
		fmt.Println(err.Error())
	}
	server.ServeHTTP(httptest.NewRecorder(), req1)

	req2 , err := newPostWinRequest(player)
	if err != nil {
		fmt.Println(err.Error())
	}
	server.ServeHTTP(httptest.NewRecorder(), req2)

	req3 , err := newPostWinRequest(player)
	if err != nil {
		fmt.Println(err.Error())
	}
	server.ServeHTTP(httptest.NewRecorder(), req3)


	response := httptest.NewRecorder()
	
	req4 , err := NewGetRequestScore(player)
	if err != nil {
		fmt.Println(err.Error())
	}
	server.ServeHTTP(response, req4)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t,"3",response.Body.String())	
}