package main

import (
	"minepin/com/cfg"
	"minepin/com/log"
	"minepin/com/utils"
	"minepin/com/web"
	"strconv"
)

func main() {
	if runCLI() {
		return
	}
	initCfg()
	initLog()
	initDB()
	initHandle()

	Addr := ":" + strconv.FormatInt(cfg.GetInt64("Port"), 10)
	log.Info("Minepin "+utils.Version()+" started at ", Addr)
	web.Run(Addr)
}
