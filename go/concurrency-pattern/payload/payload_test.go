package payload

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceGet(t *testing.T) {
	assert := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}))
	defer server.Close()

	data := Instance{URL: server.URL}
	body, err := data.Get()
	assert.NoError(err)
	assert.Equal("hello", string(body))
}
