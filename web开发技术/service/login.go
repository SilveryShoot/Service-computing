package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "login", struct{}{})
	}
}
