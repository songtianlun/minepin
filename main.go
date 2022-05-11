package main

import (
	"minepin/com/cfg"
	"minepin/com/http"
	"minepin/com/utils"
)

func main() {
	initCfg()
	initLog()
	initHandle()

	utils.P("ChitChat", utils.Version(), "started at", cfg.GetString("Address"))
	http.Run(cfg.GetString("Address"))
}
