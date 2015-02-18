package handler

import (
	"net/http"
)

type DelayableFileServer struct {
	fileServer http.Handler
}

func (f DelayableFileServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	HandleDelay(request)

	f.fileServer.ServeHTTP(response, request)
}

func NewDelayableFileServer(urlPrefix string, directoryToServe string) http.Handler {
	httpDirectory := http.Dir(directoryToServe)
	fileServer := http.FileServer(httpDirectory)
	prefixStrippedFileServer := http.StripPrefix(urlPrefix, fileServer)
	delayableFileServer := &DelayableFileServer{fileServer: prefixStrippedFileServer}

	return delayableFileServer
}
