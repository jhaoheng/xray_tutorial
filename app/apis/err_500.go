package apis

import (
	"log"
	"net/http"
)

/*
Err500 ...
*/
type Err500 struct{}

func (h *Err500) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("error 500")
	writer.WriteHeader(http.StatusBadGateway)
}
