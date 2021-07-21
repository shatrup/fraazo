package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")

	return func(t *testing.T) {
		t.Log("teardown sub test")
		router = nil
	}
}


func Test_should_return_200(t *testing.T) {
	fmt.Println("My Integration Test")
	handler := func(w http.ResponseWriter, r *http.Request) {
		// here we write our expected response, in this case, we return a
		// JSON string which is typical when dealing with REST APIs
		io.WriteString(w, "{\"key\":\"one\",\"value\":\"1\"}")
	}
	req, err := http.NewRequest("GET", "/one", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"key":"one","value":"1"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
