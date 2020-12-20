package apis

import (
	"app/utility"
	"log"
	"net/http"
)

/*
SendSQSHandler ...
*/
type SendSQSHandler struct{}

func (h *SendSQSHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("send to SQS")
	_, err := utility.SendSQSMsg(request.Context(), "hello world")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
