package handler

import (
	"fmt"
	"net/http"
)

type Echo struct{}

func (r Echo) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	HandleDelay(request)

	fmt.Fprint(response, request.URL.RequestURI())
}

func NewEcho(prefix string) http.Handler {
	return http.StripPrefix(prefix, &Echo{})
}
