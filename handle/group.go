package handle

import (
	"minepin/com/log"
	"minepin/com/utils"
	"minepin/com/web"
	"minepin/model"
	"net/http"
)

func PinGroupIndex(writer http.ResponseWriter, request *http.Request) {
	pins, err := model.GetGroupList(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get threads")
		return
	}
	web.GenerateHTML(writer, pins, "layout", "private.navbar", "index.group")
}

func NewGroup(writer http.ResponseWriter, request *http.Request) {
	web.GenerateHTML(writer, &model.Pin{}, "layout", "private.navbar", "new.group")
}

func CreateGroup(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Error("Cannot parse form")
	}
	user, err := model.GetUser(request)
	if err != nil {
		log.Error("Cannot get user from session")
	}
	group := model.GroupBind{
		Name: request.PostFormValue("name"),
		Note: request.PostFormValue("note"),
	}
	if _, err := user.CreateGroup(group); err != nil {
		log.Error("Cannot create group")
	}
	http.Redirect(writer, request, "/", 302)
}

func EditGroup(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	uuid := val.Get("pid")

	user, err := model.GetUser(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get user")
	}
	group, err := user.GetGroupByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get group")
	}
	web.GenerateHTML(writer, &group, "layout", "private.navbar", "private.group")
}