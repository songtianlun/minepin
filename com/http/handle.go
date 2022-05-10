package http

import (
	"minepin/com/cfg"
	"minepin/com/log"
	"net/http"
	"time"
)

type Handle func(w http.ResponseWriter, r *http.Request)

var mux *http.ServeMux

// Init initializes the http server
// 导入时自动实例化
func init() {
	mux = http.NewServeMux()
}

func RegisterHandle(path string, handle Handle) {
	mux.HandleFunc(path, handle)
}

func RegisterFile(path string, file string, strip bool) {
	files := http.FileServer(http.Dir(file))
	if strip {
		mux.Handle(path, http.StripPrefix(path, files))
	} else {
		mux.Handle(path, files)
	}
}

func Run(address string) {
	server := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    time.Duration(cfg.GetInt64("ReadTimeout") * int64(time.Second)),
		WriteTimeout:   time.Duration(cfg.GetInt64("WriteTimeout") * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.ErrorF("http server error: %s", err.Error())
		return
	}
}
