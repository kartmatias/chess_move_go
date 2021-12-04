package controller

import (
	"fmt"
	"net/http"
)

type HomeHandler struct {

}

func (h *HomeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Home\n")

	for name, headers := range request.Header {
		for _, h := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, h)
		}
	}

}
