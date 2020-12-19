package apis

import (
	"log"
	"net/http"
)

/*
ErrPanic ...
*/
type ErrPanic struct{}

func (h *ErrPanic) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("error panic")
	panic("error panic")
	// writer.WriteHeader(http.StatusBadRequest)
}
