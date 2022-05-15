package handle

import (
	"minepin/com/web"
	"minepin/model"
	"net/http"
)

// Err GET /err?msg=
// shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	msg := val.Get("msg")
	_, err := model.CheckSession(request)
	if err != nil {
		web.GenerateHTML(writer, msg, "layout", "public.navbar", "error")
	} else {
		web.GenerateHTML(writer, msg, "layout", "private.navbar", "error")
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	_, err := model.CheckSession(request)
	if err != nil {
		web.GenerateHTML(writer, nil, "layout", "public.navbar", "index")
	} else {
		web.GenerateHTML(writer, nil, "layout", "private.navbar", "index")
	}
}
