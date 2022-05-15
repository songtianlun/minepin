package handle

import (
	"minepin/com/utils"
	"minepin/com/web"
	"minepin/model"
	"net/http"
)

// GET /err?msg=
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
	threads, err := model.Threads()
	if err != nil {
		utils.Error_message(writer, request, "Cannot get threads")
	} else {
		_, err := model.CheckSession(request)
		if err != nil {
			web.GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			web.GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
