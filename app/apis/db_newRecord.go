package apis

import (
	"log"
	"net/http"
	"time"

	"app/utility"
)

/*
NewRecordHandler ...
*/
type NewRecordHandler struct{}

func (h *NewRecordHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("NewRecordHandler")
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"state":"error"}`))
		return
	}
	c := utility.CARS{}
	err := c.Insert(request.Context(), time.Now().String())
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"state":"error"}`))
		return
	}
	writer.WriteHeader(http.StatusOK)
}
