package handle

import (
	"minepin/com/log"
	"minepin/com/utils"
	"minepin/com/web"
	"minepin/model"
	"net/http"
)

func MinePinIndex(writer http.ResponseWriter, request *http.Request) {
	pins, err := model.GetPinList(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get threads")
		return
	}
	web.GenerateHTML(writer, pins, "layout", "private.navbar", "index.minepin")
}

func NewPin(writer http.ResponseWriter, request *http.Request) {
	web.GenerateHTML(writer, nil, "layout", "private.navbar", "new.pin")
}

func CreatePin(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Error("Cannot parse form")
	}
	user, err := model.GetUser(request)
	if err != nil {
		log.Error("Cannot get user from session")
	}
	pin := model.PinBind{
		Location: request.PostFormValue("location"),
		Lat:      request.PostFormValue("latitude"),
		Lng:      request.PostFormValue("longitude"),
		Note:     request.PostFormValue("note"),
	}
	if _, err := user.CreatePin(pin); err != nil {
		log.Error("Cannot create thread")
	}
	http.Redirect(writer, request, "/", 302)
}

func EditPin(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	uuid := val.Get("pid")

	user, err := model.GetUser(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get user")
	}
	pin, err := user.GetPinByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get pin")
	}
	web.GenerateHTML(writer, &pin, "layout", "private.navbar", "private.pin")
}

func UpdatePin(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot parse form")
		return
	}
	uuid := request.PostFormValue("uuid")
	user, _ := model.GetUser(request)
	pin, err := user.GetPinByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot read pin")
		return
	}
	pin.Location = request.PostFormValue("location")
	pin.Lat = request.PostFormValue("latitude")
	pin.Lng = request.PostFormValue("longitude")
	pin.Note = request.PostFormValue("note")

	if err := pin.UpdatePin(); err != nil {
		log.Error("Cannot update pin - " + err.Error())
	}
	http.Redirect(writer, request, "/minepin", 302)
}

func DeletePin(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	uuid := val.Get("pid")
	user, err := model.GetUser(request)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot read pin - "+err.Error())
		return
	}
	pin, err := user.GetPinByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot read pin - "+err.Error())
		return
	}

	if err := pin.Delete(); err != nil {
		log.Error("Cannot delete pin - " + err.Error())
	}
	http.Redirect(writer, request, "/minepin", 302)
}
