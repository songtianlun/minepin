package handle

import (
	"fmt"
	"minepin/com/log"
	"minepin/com/utils"
	"minepin/com/web"
	"minepin/model"
	"net/http"
)

// GET /threads/new
// Show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := model.CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		web.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {
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
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			log.Error("Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := model.ThreadByUUID(uuid)
	if err != nil {
		utils.Error_message(writer, request, "Cannot read thread")
	} else {
		_, err := model.CheckSession(request)
		if err != nil {
			web.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			web.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
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
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := model.ThreadByUUID(uuid)
		if err != nil {
			utils.Error_message(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			log.Error("Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
