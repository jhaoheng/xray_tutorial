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
	log.Println("SQLByXRayHandler")

	obj, err := utility.SQLByXrayWithError(request.Context())
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		writer.Write([]byte(`bad sql`))
		return
	}
	b, _ := json.Marshal(obj)
	writer.Write(b)
	writer.WriteHeader(http.StatusOK)
}
