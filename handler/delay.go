package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

const DELAY_HEADER_KEY string = "delay-milliseconds"

/*
	Checks the request for a header which indicates that there should be a delay before
	returning the response. Useful for simulating high latency.

	Returns true if a delay header was found, parsed, and processed successfully.
	Returns false otherwise.
*/
func HandleDelay(request *http.Request) bool {
	delayMillisString := request.Header.Get(DELAY_HEADER_KEY)

	if delayMillisString != "" {
		const BASE_10 int = 10
		delayTimeMillis, err := strconv.ParseInt(delayMillisString, BASE_10, 0)

		if err != nil {
			log.Printf("Received unparseable delay time in header: \"%v\"\n", delayMillisString)
		} else {
			delayDuration := time.Duration(delayTimeMillis) * time.Millisecond
			time.Sleep(delayDuration)
			return true
		}
	}

	return false
}
