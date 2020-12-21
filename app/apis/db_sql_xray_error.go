package apis

import (
	"app/utility"
	"encoding/json"
	"log"
	"net/http"
)

/*
SQLByXrayWithError ...
*/
type SQLByXrayWithError struct{}

func (h *SQLByXrayWithError) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("SQLByXrayWithError")

	obj, err := utility.SQLByXrayWithError(request.Context())
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		writer.Write([]byte(err.Error()))
		return
	}
	b, _ := json.Marshal(obj)
	writer.Write(b)
	writer.WriteHeader(http.StatusOK)
}
