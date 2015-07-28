package handler

import (
	"fmt"
	"math/rand"
	"net/http"
)

func Random(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, rand.Int())
}
