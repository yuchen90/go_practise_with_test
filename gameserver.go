package gameserver

import (
	"net/http"
	"net/http/httptest"
)

func PlayerServer(request *http.Request, response *httptest.ResponseRecorder) {
	response.Write([]byte("20"))
}
