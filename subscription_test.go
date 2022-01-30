package huaweistore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendRequest(t *testing.T) {
	c := New()
	assert := assert.New(t)
	ctx := context.Background()
	srv := serverMock()
	defer srv.Close()

	tests := []struct {
		dest         string
		clientID     string
		clientSecret string
		token        string
		checkURL     string
		authURL      string
	}{
		{
			dest:         "success",
			clientID:     "clientid",
			clientSecret: "clientsecret",
			token:        "Basic QVBQQVQ6dG9rZW4=",
			checkURL:     "http://127.0.0.1:8080/v1/check",
			authURL:      "http://127.0.0.1:8080/v1/auth",
		},
		{
			dest:         "empty client id",
			clientID:     "",
			clientSecret: "clientsecret",
			token:        "Basic QVBQQVQ6dG9rZW4=",
			checkURL:     "http://127.0.0.1:8080/v1/check",
			authURL:      "http://127.0.0.1:8080/v1/auth",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("test name %s", test.dest), func(t *testing.T) {

			reqBody := IAPRequest{
				ClientID:     test.clientID,
				ClientSecret: test.clientSecret,
			}

			c.TokenURL = test.authURL
			responseBody, err := c.sendRequest(ctx, test.checkURL, reqBody)
			if err != nil {
				assert.EqualError(err, `Get token fail, {"access_token": ""}`, "expected error string")
			}

			var body = struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"message"`
			}{}

			if err := json.Unmarshal(responseBody, &body); err != nil {
				assert.EqualError(err, `unexpected end of JSON input`, "expected error string")
			}

			if body.StatusCode != 0 {
				t.Errorf("server error. message: %s", body.Message)
			}
		})
	}
}

func serverMock() *httptest.Server {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/v1/check", srvRespMock)
	handler.HandleFunc("/v1/auth", srvAuth)

	srv := httptest.NewUnstartedServer(handler)
	srv.Listener.Close()
	srv.Listener = l
	srv.Start()

	return srv
}

func srvRespMock(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json; charset=UTF-8" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf(`{"status_code": 6, "message": "content type error. got: %s"}`, contentType))
		return
	}

	authHeader := r.Header.Get("Authorization")
	auth := strings.Split(authHeader, " ")[1]
	if auth != "QVBQQVQ6dG9rZW4=" {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, fmt.Sprintf(`{"status_code": 6, "message": "auth error. got: %s"}`, auth))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status_code": 0, "message": "success"}`)
}

func srvAuth(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("client_id") == "" || r.PostFormValue("client_secret") == "" {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, `{"access_token": ""}`)
		return
	}

	_, _ = w.Write([]byte(`{
		"access_token": "token"
	}`))
}
