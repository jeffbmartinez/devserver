package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DelayableFileServer struct {
	fileServer http.Handler
}

func (f DelayableFileServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	HandleDelay(request)

	vars := mux.Vars(request)

	pathname := fmt.Sprintf("/%v", vars["pathname"])

	// Forward a modified version of the request to http.FileServer.
	// It already knows how to retrieve a file.
	request.URL.Path = pathname

	f.fileServer.ServeHTTP(response, request)
}

func NewDelayableFileServer(directoryToServe string) http.Handler {
	httpDirectory := http.Dir(directoryToServe)
	fileServer := http.FileServer(httpDirectory)
	delayableFileServer := &DelayableFileServer{fileServer: fileServer}

	return delayableFileServer
}
