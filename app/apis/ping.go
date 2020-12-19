package apis

import (
	"log"
	"net/http"
)

/*
PingHandler ...
*/
type PingHandler struct{}

func (h *PingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("ping requested, responding with HTTP 200")
	writer.WriteHeader(http.StatusOK)
}
