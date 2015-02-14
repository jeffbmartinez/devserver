package handler

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Random struct{}

func (r Random) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, rand.Int())
}

func NewRandom() http.Handler {
	return &Random{}
}
