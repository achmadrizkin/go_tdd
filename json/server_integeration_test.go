package json

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThemSUCCESS(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func TestRecordingWinsAndRetrievingThemFAILBadStatus(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatusFalse(t, response.Code, http.StatusBadGateway)

	assertResponseBody(t, response.Body.String(), "3")
}
