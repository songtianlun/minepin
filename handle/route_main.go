package handle

import (
	"minepin/com/utils"
	"minepin/data"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := utils.Session(writer, request)
	if err != nil {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		utils.Error_message(writer, request, "Cannot get threads")
	} else {
		_, err := utils.Session(writer, request)
		if err != nil {
			utils.GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
