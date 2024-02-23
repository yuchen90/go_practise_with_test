package gameserver_test

import (
	gameserver "go_practise/go_practise_with_test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGamerServer(t *testing.T) {
	t.Run("return Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/Pepper", nil)
		response := httptest.NewRecorder()

		gameserver.PlayerServer(request, response)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
