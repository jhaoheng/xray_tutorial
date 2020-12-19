package apis

import (
	"log"
	"net/http"
)

/*
RootHandler ...
*/
type RootHandler struct{}

func (h *RootHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("Hello, world!")
	writer.Write([]byte("Hello, world!"))
}
