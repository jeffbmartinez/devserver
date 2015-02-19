package handler

import (
	"fmt"
	"math/rand"
	"net/http"
)

func Random(response http.ResponseWriter, request *http.Request) {
	HandleDelay(request)

	fmt.Fprint(response, rand.Int())
}
