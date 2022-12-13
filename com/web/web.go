package web

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sync"
)

type Weber interface {
	Tpler
	Server
	HTTPer

	RegisterTplEmbedFs(efs *embed.FS)
}

type Server interface {
	Init()
	Run(addr string) error
	Stop() error
}

type HTTPer interface {
	RegisterEmbedHandleFs(path string, fs *embed.FS)
	RegisterGet(path string, handler http.HandlerFunc)
	RegisterEmbedHandleWithStripPrefix(path string, preFix string, fs *embed.FS)
	RegisterEmbedWithSub(path, stripPrefix, subPath string, efs *embed.FS)
	RegisterGlobelMiddleware(middlewares ...func(http.Handler) http.Handler)
	RegisterGroup(fn func(r chi.Router))
	RegisterHandle(pattern string, handler http.HandlerFunc)
	RegisterHandleFuncWithMethod(method, pattern string, handler http.HandlerFunc)
	RegisterHandleWithMiddlewares(pattern string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler)

	GetRouters() (routes map[string]string)
}

func NewWebServer() Weber {
	return new(chiWeb)
}

var instance Weber
var once sync.Once

func GetInstance() Weber {
	once.Do(func() {
		instance = NewWebServer()
		instance.Init()
	})
	return instance
}
