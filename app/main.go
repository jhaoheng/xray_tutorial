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

	/*
		If the hostname in the request matches the pattern *.example.com, use the hostname. Otherwise, use MyApp.
	*/
	// xraySegmentNamer := xray.NewDynamicSegmentNamer(utility.GetXRAYAppName(), "*.example.com")
	xraySegmentNamer := xray.NewFixedSegmentNamer(utility.GetXRAYAppName())

	// API
	http.Handle("/", xray.Handler(xraySegmentNamer, &apis.RootHandler{}))
	http.Handle("/ping", xray.Handler(xraySegmentNamer, &apis.PingHandler{}))
	// DB
	http.Handle("/new", xray.Handler(xraySegmentNamer, &apis.NewRecordHandler{}))
	http.Handle("/del", xray.Handler(xraySegmentNamer, &apis.DelRecordHandler{}))
	http.Handle("/all", xray.Handler(xraySegmentNamer, &apis.GetAllRecordHandler{}))
	// error
	http.Handle("/error/400", xray.Handler(xraySegmentNamer, &apis.Err400{}))
	http.Handle("/error/429", xray.Handler(xraySegmentNamer, &apis.Err429{}))
	http.Handle("/error/500", xray.Handler(xraySegmentNamer, &apis.Err500{}))
	http.Handle("/error/panic", xray.Handler(xraySegmentNamer, &apis.ErrPanic{}))
	// other interesting
	http.Handle("/many/funcs", xray.Handler(xraySegmentNamer, &apis.ManyFuncsHandler{}))
	http.Handle("/send/sqs", xray.Handler(xraySegmentNamer, &apis.SendSQSHandler{}))
	//
	http.ListenAndServe(":"+defaultPort, nil)
}
