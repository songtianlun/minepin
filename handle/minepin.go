package handle

import (
	"minepin/com/log"
	"minepin/com/utils"
	"minepin/com/web"
	"minepin/model"
	"net/http"
)

func MinePinIndex(writer http.ResponseWriter, request *http.Request) {
	//pins, err := model.Pins()
	sess, err := model.CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		user, _ := sess.User()
		pins, err := user.PinList()
		if err != nil {
			utils.Error_message(writer, request, "Cannot get threads")
			return
		}
		web.GenerateHTML(writer, pins, "layout", "private.navbar", "index.minepin")
	}
}

func NewPin(writer http.ResponseWriter, request *http.Request) {
	_, err := model.CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		web.GenerateHTML(writer, nil, "layout", "private.navbar", "new.pin")
	}
}

func CreatePin(writer http.ResponseWriter, request *http.Request) {
	sess, err := model.CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			log.Error("Cannot parse form")
		}
		user, err := sess.User()
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
}

func EditPin(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	uuid := val.Get("pid")
	pin, err := model.GetPinByUUID(uuid)
	if err != nil {
		utils.Error_message(writer, request, "Cannot read thread")
	} else {
		_, err := model.CheckSession(request)
		if err != nil {
			http.Redirect(writer, request, "/login", 302)
		} else {
			web.GenerateHTML(writer, &pin, "layout", "private.navbar", "private.pin")
		}
	}
}

func UpdatePin(writer http.ResponseWriter, request *http.Request) {
	sess, err := model.CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			log.Error("Cannot parse form")
		}
		uuid := request.PostFormValue("uuid")
		pin, err := model.GetPinByUUID(uuid)
		if err != nil {
			utils.Error_message(writer, request, "Cannot read pin")
			return
		}
		pin.Location = request.PostFormValue("location")
		pin.Lat = request.PostFormValue("latitude")
		pin.Lng = request.PostFormValue("longitude")
		pin.Note = request.PostFormValue("note")

		if pin.UserId != sess.UserId {
			utils.Error_message(writer, request, "It's not your pin, ACL Error!")
			return
		}

		if err := pin.UpdatePin(); err != nil {
			log.Error("Cannot update pin - " + err.Error())
		}
		http.Redirect(writer, request, "/minepin", 302)
	}
}
