package handler

import (
	"net/http"
)

func NewFileServer(urlPrefix string, directoryToServe string) http.Handler {
	httpDirectory := http.Dir(directoryToServe)
	fileServer := http.FileServer(httpDirectory)
	prefixStrippedFileServer := http.StripPrefix(urlPrefix, fileServer)
	return prefixStrippedFileServer
}
