package handle

import (
	"minepin/com/log"
	"minepin/com/utils"
	"minepin/data"
	"net/http"
)

// GET /login
// Show the Login page
func Login(writer http.ResponseWriter, request *http.Request) {
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the Signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	utils.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Error("Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		log.Error("Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		log.Error("Cannot find user")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Error("Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}

}

// GET /logout
// Logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		log.Warn("Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
