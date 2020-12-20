package apis

import (
	"app/utility"
	"log"
	"net/http"
	"time"
)

/*
ManyFuncsHandler ...
*/
type ManyFuncsHandler struct{}

func (h *ManyFuncsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("many funcs ... ")
	//
	utility.DbXrayMiddle(request.Context(), "funcOne", func() error {
		funcOne()
		return nil
	})
	//
	utility.DbXrayMiddle(request.Context(), "funcTwo", func() error {
		funcTwo()
		return nil
	})
	utility.DbXrayMiddle(request.Context(), "funcThree", func() error {
		funcThree()
		return nil
	})
	writer.WriteHeader(http.StatusOK)
}

func funcOne() {
	log.Println("do func one")
	time.Sleep(1 * time.Second)
}

func funcTwo() {
	log.Println("do func two")
	time.Sleep(500 * time.Microsecond)
}

func funcThree() {
	log.Println("do func three")
	time.Sleep(2 * time.Second)
}
