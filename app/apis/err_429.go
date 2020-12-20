package apis

import (
	"log"
	"net/http"
)

/*
Err429 ...
*/
type Err429 struct{}

func (h *Err429) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("error 429, too many requests")
	writer.WriteHeader(http.StatusTooManyRequests)
}
