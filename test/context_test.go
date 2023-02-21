package test

import (
	"go_tdd/logic"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServerSUCCESS(t *testing.T) {
	data := "hello, world"
	svr := logic.Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	assert.Equal(t, response.Body.String(), data)
}
