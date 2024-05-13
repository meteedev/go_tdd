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
		fmt.Errorf(err.Error())
	}

	res := httptest.NewRecorder()
	PlayerServer(res, req)

	t.Run("get player scoere", func(t *testing.T) {
		got := res.Body.String()

		expect := "80"
		assert.Equal(t, expect, got)
	})

}
