package main

import (
	"minepin/com/cfg"
	"minepin/com/utils"
	"minepin/com/web"
)

func main() {
	initCfg()
	initLog()
	initDB()
	initHandle()

	utils.P("ChitChat", utils.Version(), "started at", cfg.GetString("Address"))
	web.Run(cfg.GetString("Address"))
}
