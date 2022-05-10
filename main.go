package main

import (
	"minepin/com/cfg"
	"minepin/com/http"
)

func main() {
	initCfg()
	initHandle()

	p("ChitChat", version(), "started at", cfg.GetString("Address"))
	http.Run(cfg.GetString("Address"))
}
