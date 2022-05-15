package main

import (
	"minepin/com/cfg"
	"minepin/com/utils"
	"minepin/com/web"
	"strconv"
)

func main() {
	initCfg()
	initLog()
	initDB()
	initHandle()

	Addr := ":" + strconv.FormatInt(cfg.GetInt64("Port"), 10)
	utils.P("ChitChat", utils.Version(), "started at: ", Addr)
	web.Run(Addr)
}
