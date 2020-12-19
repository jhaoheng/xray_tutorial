package apis

import (
	"app/utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
DelRecordHandler ...
*/
type DelRecordHandler struct{}

func (h *DelRecordHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("DelRecordHandler")
	if request.Method != "DELETE" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"state":"error"}`))
		return
	}

	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	type DelObj struct {
		ID int `json:"id"`
	}
	delObj := DelObj{}
	err := json.Unmarshal(b, &delObj)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"state":"error"}`))
		return
	}
	c := utility.CARS{}
	err = c.Delete(request.Context(), delObj.ID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"state":"error"}`))
		return
	}
	writer.WriteHeader(http.StatusOK)
}
