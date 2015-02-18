package handler

import (
	"fmt"
	"net/http"
)

type Counter int

func (c *Counter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	HandleDelay(request)

	*c++
	fmt.Fprint(response, *c)
}

func NewCounter() http.Handler {
	return new(Counter)
}
