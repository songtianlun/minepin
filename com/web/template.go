package web

import (
	"embed"
	"net/http"
)

type Tpler interface {
	RegisterTplEmbedFs(efs *embed.FS)
	GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string)
}
