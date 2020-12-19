package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-xray-sdk-go/xray"

	"app/apis"
	"app/utility"
)

const defaultPort = "9001"

func main() {
	log.Println("starting server, listening on port " + defaultPort)
	//
	// xraySegmentNamer := xray.NewFixedSegmentNamer(utility.GetXRAYAppName())
	xraySegmentNamer := xray.NewDynamicSegmentNamer(utility.GetXRAYAppName(), "*")

	// API
	http.Handle("/", xray.Handler(xraySegmentNamer, &apis.RootHandler{}))
	http.Handle("/ping", xray.Handler(xraySegmentNamer, &apis.PingHandler{}))
	// DB
	http.Handle("/new", xray.Handler(xraySegmentNamer, &apis.NewRecordHandler{}))
	http.Handle("/del", xray.Handler(xraySegmentNamer, &apis.DelRecordHandler{}))
	http.Handle("/all", xray.Handler(xraySegmentNamer, &apis.GetAllRecordHandler{}))
	// error
	http.Handle("/error/400", xray.Handler(xraySegmentNamer, &apis.Err400{}))
	http.Handle("/error/panic", xray.Handler(xraySegmentNamer, &apis.ErrPanic{}))
	//
	http.ListenAndServe(":"+defaultPort, nil)
}
