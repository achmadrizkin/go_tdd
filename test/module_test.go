package test

import (
	logic "go_tdd/logic"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckHealthSUCCESS(t *testing.T) {
	t.Run("Check health status success", func(t *testing.T) {
		// request header
		r := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		w := httptest.NewRecorder()
		logic.CheckHealth(w, r)

		response := w.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "text/plain; charset=utf-8", response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}

func TestCheckHealthFAILBadResponse(t *testing.T) {
	t.Run("Check health status success", func(t *testing.T) {
		// request header
		r := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		w := httptest.NewRecorder()
		logic.CheckHealth(w, r)

		response := w.Result()
		body, err := io.ReadAll(response.Body)

		assert.NotEqual(t, "BAD RESPONSE", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "text/plain; charset=utf-8", response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}

func TestCheckHealthFAILBadStatusCode(t *testing.T) {
	t.Run("Check health status success", func(t *testing.T) {
		// request header
		r := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		w := httptest.NewRecorder()
		logic.CheckHealth(w, r)

		response := w.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.NotEqual(t, 400, response.StatusCode)
		assert.Equal(t, "text/plain; charset=utf-8", response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}

func TestCheckHealthFAILBadContentType(t *testing.T) {
	t.Run("Check health status success", func(t *testing.T) {
		// request header
		r := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		w := httptest.NewRecorder()
		logic.CheckHealth(w, r)

		response := w.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.NotEqual(t, "video/test; charset=utf-8", response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}
