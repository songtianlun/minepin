package web

import (
	"context"
	"embed"
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"io/fs"
	"minepin/com/log"
	"net/http"
	"strings"
	"time"
)

type chiWeb struct {
	s      *http.Server
	r      chi.Router
	tplEFS *embed.FS
}

func (cw *chiWeb) Run(addr string) error {
	cw.s = &http.Server{
		Addr:    addr,
		Handler: cw.r,
	}
	return cw.s.ListenAndServe()
}

func (cw *chiWeb) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if cw.s != nil {
		cw.s.SetKeepAlivesEnabled(false)
		return cw.s.Shutdown(ctx)
	}

	return fmt.Errorf("web server is not running")
}

func (cw *chiWeb) Init() {
	cw.r = chi.NewRouter()
}

func (cw *chiWeb) RegisterEmbedHandleWithStripPrefix(path string, preFix string, fs *embed.FS) {
	cw.r.Handle(path,
		http.StripPrefix(preFix,
			http.FileServer(http.FS(fs)),
		),
	)
}

func (cw *chiWeb) RegisterEmbedWithSub(path, stripPrefix, subPath string, efs *embed.FS) {
	nfs, err := fs.Sub(efs, subPath)
	if err != nil {
		panic("register embed with sub error: " + err.Error())
	}

	cw.r.Get(path, func(writer http.ResponseWriter, request *http.Request) {
		if stripPrefix != "" {
			http.StripPrefix(stripPrefix, http.FileServer(http.FS(nfs))).ServeHTTP(writer, request)
		} else {
			http.FileServer(http.FS(nfs)).ServeHTTP(writer, request)
		}
	})
}

func (cw *chiWeb) RegisterGet(path string, handler http.HandlerFunc) {
	cw.r.Get(path, handler)
}

func (cw *chiWeb) RegisterEmbedHandleFs(path string, fs *embed.FS) {
	if path != "/" && path[len(path)-1] != '/' {
		cw.r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	cw.r.Get(path,
		func(w http.ResponseWriter, r *http.Request) {
			rctx := chi.RouteContext(r.Context())
			pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
			fs := http.StripPrefix(pathPrefix, http.FileServer(http.FS(fs)))
			fs.ServeHTTP(w, r)
		},
	)
}

func (cw *chiWeb) RegisterTplEmbedFs(efs *embed.FS) {
	cw.tplEFS = efs
}

func (cw *chiWeb) RegisterGlobelMiddleware(middlewares ...func(http.Handler) http.Handler) {
	cw.r.Use(middlewares...)
}

func (cw *chiWeb) RegisterGroup(fn func(r chi.Router)) {
	cw.r.Group(fn)
}

func (cw *chiWeb) RegisterHandle(pattern string, handler http.HandlerFunc) {
	cw.r.Handle(pattern, handler)
}

func (cw *chiWeb) RegisterHandleFuncWithMethod(method string, pattern string, handler http.HandlerFunc) {
	cw.r.MethodFunc(method, pattern, handler)
}

func (cw *chiWeb) RegisterHandleWithMiddlewares(pattern string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	cw.r.With(middlewares...).Handle(pattern, handler)
}

func (cw *chiWeb) GetRouters() (routes map[string]string) {
	_ = chi.Walk(cw.r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		// log.Infof("%s %s", method, route)
		if routes == nil {
			routes = make(map[string]string)
		}
		routes[route] = method
		return nil
	})
	return
}

func (cw *chiWeb) GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	if cw.tplEFS == nil {
		log.Error("template embed fs is nil")
		return
	}

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.tmpl", file))
	}

	templates := template.Must(template.ParseFS(cw.tplEFS, files...))
	err := templates.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		log.Errorf("Generate HTML error: %v", err.Error())
	}
}
