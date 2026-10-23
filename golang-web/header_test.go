package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	header := request.Header.Get("Content-Type")
	fmt.Fprint(writer, header)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recoder := httptest.NewRecorder()
	RequestHeader(recoder, request)
	
	response := recoder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}