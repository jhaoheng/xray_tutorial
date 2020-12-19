package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-xray-sdk-go/xray"
)

const defaultPort = "9001"

func getServerPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return defaultPort
}

type rootHandler struct{}

func (h *rootHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("Hello, world!")
	writer.Write([]byte("Hello, world!"))
}

type pingHandler struct{}

func (h *pingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("ping requested, responding with HTTP 200")
	writer.WriteHeader(http.StatusOK)
}

func main() {
	log.Println("starting server, listening on port " + defaultPort)
	xraySegmentNamer := xray.NewFixedSegmentNamer(getXRAYAppName())
	http.Handle("/", xray.Handler(xraySegmentNamer, &rootHandler{}))
	http.Handle("/ping", xray.Handler(xraySegmentNamer, &pingHandler{}))
	http.ListenAndServe(":"+defaultPort, nil)
}

func getXRAYAppName() string {
	appName := os.Getenv("XRAY_APP_NAME")
	if appName != "" {
		return appName
	}

	return "myApp"
}

func getXRAYAddr() string {
	XRAY_DAEMON_ADDR := os.Getenv("XRAY_DAEMON_ADDR")
	if XRAY_DAEMON_ADDR != "" {
		return XRAY_DAEMON_ADDR
	}
	return "localhost:2000"
}

func init() {
	xray.Configure(xray.Config{
		DaemonAddr:     getXRAYAddr(), // default
		ServiceVersion: "0.0.1",
	})
}
