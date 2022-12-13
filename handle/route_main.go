package handle

import (
	"minepin/com/utils"
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
		web.GetInstance().GenerateHTML(writer, msg, "layout", "public.navbar", "error")
	} else {
		web.GetInstance().GenerateHTML(writer, msg, "layout", "private.navbar", "error")
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	sess, err := model.CheckSession(request)
	if err != nil {
		web.GetInstance().GenerateHTML(writer, nil, "layout", "public.navbar", "index")
	} else {
		user, err := sess.User()
		if err != nil {
			utils.ErrorMessage(writer, request, "failed to get user.")
		}
		web.GetInstance().GenerateHTML(writer, model.IndexModel{
			User:       user,
			PinCount:   user.GetPinCount(),
			GroupCount: user.GetGroupCount(),
		}, "layout", "private.navbar", "private.index")
	}
}
