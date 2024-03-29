package handle

import (
	"minepin/com/constvar"
	"minepin/com/log"
	"minepin/com/utils"
	"minepin/com/web"
	"minepin/model"
	"net/http"
)

// GET /login
// Show the Login page
func Login(writer http.ResponseWriter, request *http.Request) {
	// t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	// t.Execute(writer, nil)
	web.GetInstance().GenerateHTML(writer, nil, "login.layout", "public.navbar", "login")
}

// GET /signup
// Show the Signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	web.GetInstance().GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Error("Cannot parse form")
	}
	user := model.User{
		Role:     constvar.UserRegistered,
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		log.Error("Cannot create user - %v", err.Error())
	}
	http.Redirect(writer, request, "/login", 302)
}

// Authenticate the user given the email and password,  POST /authenticate
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := model.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		log.Error("Cannot find user")
	}
	if user.Password == utils.Encrypt(request.PostFormValue("password")) {
		user.CloseAllSession()
		session, err := user.CreateSession()
		if err != nil {
			log.Errorf("Cannot create session - %v", err.Error())
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		user.TransfromWithBD09() // 登录时将账户下的百度坐标系转换为WGS84坐标
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
		model.DeleteSession(cookie.Value)
	}
	http.Redirect(writer, request, "/", 302)
}
