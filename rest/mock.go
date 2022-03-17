package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func MockRequest(httpMethod string, endpoint string, body interface{}) *httptest.ResponseRecorder {
	router := SetupRouter()

	w := httptest.NewRecorder()

	var bodyJSON []byte

	if body != nil {
		bodyJSON, _ = json.Marshal(body)
	}

	fmt.Println(string(bodyJSON))

	req, _ := http.NewRequest(httpMethod, endpoint, bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())

	return w
}
