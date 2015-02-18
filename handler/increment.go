package handler

import (
	"fmt"
	"net/http"
)

type Increment struct {
	number int
}

func (r *Increment) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	HandleDelay(request)

	r.number++
	fmt.Fprint(response, r.number)
}

func NewIncrement() http.Handler {
	return &Increment{
		number: 0,
	}
}
