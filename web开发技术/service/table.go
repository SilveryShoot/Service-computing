package service

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

type User struct {
	Username string
	Password string
}

func tableHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {

		}
		user := new(User)
		decoder := schema.NewDecoder()
		err = decoder.Decode(user, req.PostForm)
		if err != nil {

		}
		formatter.HTML(w, http.StatusOK, "table", struct {
			Username string
			Password string
		}{Username: user.Username, Password: user.Password})
	}
}
