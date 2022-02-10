package apis

import (
	"fmt"
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
		text := fmt.Errorf("the method is wrong, %v", request.Method)
		fmt.Println("=>", text)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	c := utility.CARS{
		Name: time.Now().String(),
	}
	err := c.Insert(request.Context())

	//
	if err != nil {
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
}
