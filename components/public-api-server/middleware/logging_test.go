package middleware

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	logInMemory := &bytes.Buffer{}
	log.SetOutput(logInMemory)
	expectedBody := `hello world`

	someHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(expectedBody))
	})
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder() // this records the response

	m := NewLoggingMiddleware()
	wrappedHandler := m(someHandler)
	wrappedHandler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Something went wrong with status code %v", status)
	}

	if rec.Body.String() != expectedBody {
		t.Errorf("Unexpected body: %v", rec.Body.String())
	}

	assert.Equal(t, logInMemory, logInMemory)
}
