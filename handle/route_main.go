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
	vals := request.URL.Query()
	_, err := model.CheckSession(request)
	if err != nil {
		web.GenerateHTML(writer, vals, "layout", "public.navbar", "index")
	} else {
		web.GenerateHTML(writer, vals, "layout", "private.navbar", "index")
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
