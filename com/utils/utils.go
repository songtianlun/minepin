package utils

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func StrToInt64(s string) (i uint64) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		i = 0
	}
	return
}

func StrToFloat64(s string) (i float64) {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		i = 0
	}
	return
}

func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// ErrorMessage Convenience function to redirect to the error message page
func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Checks if the user is logged in and has a Session, if not err is not nil
// func Session(writer web.ResponseWriter, request *web.Request) (sess model.Session, err error) {
//	cookie, err := request.Cookie("_cookie")
//	if err == nil {
//		sess = model.Session{
//			BaseModel: model.BaseModel{UUID: cookie.Value},
//		}
//		if ok, _ := sess.Check(); !ok {
//			err = errors.New("Invalid session")
//		}
//	}
//	return
// }

// parse HTML templates
// pass in a list of file names, and get a template
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// func GenerateHTML(writer web.ResponseWriter, data interface{}, filenames ...string) {
//	var files []string
//	for _, file := range filenames {
//		files = append(files, fmt.Sprintf("templates/%s.html", file))
//	}
//
//	templates := template.Must(template.ParseFiles(files...))
//	err := templates.ExecuteTemplate(writer, "layout", data)
//	if err != nil {
//		log.Errorf("Generate HTML error: %v", err.Error())
//	}
// }

// Version
func Version() string {
	return "0.1"
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

func GetTypeString(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
