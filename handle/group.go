package handle

import (
	"minepin/com/constvar"
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

func ShowGroup(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	id := utils.StrToInt64(val.Get("id"))

	user, err := model.GetUser(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get user")
		return
	}
	group, err := user.GetGroupByID(id)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get group")
		return
	}
	pins, err := user.ShowPinsByGroupID(id)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get pins by group")
		return
	}
	model.TransformPins(&pins)
	switch group.Type {
	case constvar.PingsMapRoute:
		web.GenerateHTML(writer, model.Pins{
			Group:       group,
			Pins:        pins,
			BaiduAK:     user.BaiduAK(),
			TianDiTuKey: user.TianDiTuKey(),
			MapBoxKey:   user.MapBoxKey(),
		}, "layout", "private.navbar", "index.group.pin.route")
	default:
		web.GenerateHTML(writer, model.Pins{
			Group:       group,
			Pins:        pins,
			BaiduAK:     user.BaiduAK(),
			TianDiTuKey: user.TianDiTuKey(),
			MapBoxKey:   user.MapBoxKey(),
		}, "layout", "private.navbar", "index.group.pin")
	}

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
	http.Redirect(writer, request, "/group", 302)
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

func UpdateGroup(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot parse form")
		return
	}
	uuid := request.PostFormValue("uuid")
	user, err := model.GetUser(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get user")
		return
	}
	group, err := user.GetGroupByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot read group")
		return
	}
	group.Name = request.PostFormValue("name")
	group.Note = request.PostFormValue("note")
	group.Type = request.PostFormValue("type")
	if err := group.UpdateGroup(); err != nil {
		utils.ErrorMessage(writer, request, "Cannot update group - "+err.Error())
		return
	}
	http.Redirect(writer, request, "/group", 302)
}

func DeleteGroup(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	uuid := val.Get("uid")

	user, err := model.GetUser(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get user")
		return
	}
	group, err := user.GetGroupByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get group")
		return
	}
	if err := group.Delete(); err != nil {
		utils.ErrorMessage(writer, request, "Cannot delete group - "+err.Error())
		return
	}
	http.Redirect(writer, request, "/group", 302)
}
