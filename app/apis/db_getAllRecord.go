package apis

import (
	"app/utility"
	"encoding/json"
	"log"
	"net/http"
)

/*
GetAllRecordHandler ...
*/
type GetAllRecordHandler struct{}

func (h *GetAllRecordHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetAllRecordHandler")
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"state":"error"}`))
		return
	}

	c := utility.CARS{}
	objs := c.GetAll(request.Context())
	// objs := utility.GetAllByXray()
	writer.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(objs)
	writer.Write(b)
}
