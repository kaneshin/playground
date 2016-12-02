package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("pingHandler", func(t *testing.T) {
		t.Parallel()

		s := httptest.NewServer(http.HandlerFunc(pingHandler()))
		defer s.Close()

		res, err := http.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, "text/plain", res.Header.Get("Content-Type"))
		assert.Equal(t, 200, res.StatusCode)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "pong", string(body))
	})

	t.Run("echoHandler", func(t *testing.T) {
		candidates := []struct {
			query    string
			expected string
		}{
			{"", ""},
			{"foo=bar", ""},
			{"msg=foo", "foo"},
		}
		for _, c := range candidates {
			c := c
			t.Run(c.query, func(t *testing.T) {
				t.Parallel()

				s := httptest.NewServer(http.HandlerFunc(echoHandler()))
				defer s.Close()

				res, err := http.Get(fmt.Sprintf("%v?%v", s.URL, c.query))
				assert.NoError(t, err)
				assert.Equal(t, "text/plain", res.Header.Get("Content-Type"))
				assert.Equal(t, 200, res.StatusCode)

				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				assert.NoError(t, err)
				assert.Equal(t, c.expected, string(body))
			})
		}
	})
}

func TestHandlerWithRecorder(t *testing.T) {
	t.Run("echoHandler", func(t *testing.T) {
		candidates := []struct {
			url      string
			expected string
		}{
			{"http://example.com/?", ""},
			{"http://example.com/?foo=bar", ""},
			{"http://example.com/?msg=foo", "foo"},
		}

		for _, c := range candidates {
			c := c
			t.Run(c.url, func(t *testing.T) {
				t.Parallel()

				res := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodGet, c.url, nil)
				assert.NoError(t, err)

				handler := echoHandler()
				handler(res, req)

				assert.Equal(t, "text/plain", res.HeaderMap.Get("Content-Type"))
				assert.Equal(t, 200, res.Code)

				body, err := ioutil.ReadAll(res.Body)
				assert.NoError(t, err)
				assert.Equal(t, c.expected, string(body))
			})
		}
	})
}
