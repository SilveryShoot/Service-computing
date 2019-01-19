package service

import (
	"fmt"
	"net/http"
)

func unknown() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "501 Not Implemented")
	}
}
