package apis

import (
	"log"
	"net/http"
)

/*
Err400 ...
*/
type Err400 struct{}

func (h *Err400) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("error 400")
	writer.WriteHeader(http.StatusBadRequest)
}
