package main

import (
	"fmt"
	"minepin/com/cfg"
	"minepin/com/http"
	"minepin/com/log"
	"minepin/handle"
)

func initHandle() {
	// static file
	http.RegisterFile("/static/", cfg.GetString("Static"), true)

	// index
	http.RegisterHandle("/", handle.Index)

	// error
	http.RegisterHandle("/err", handle.Err)

	// defined in route_auth.go
	http.RegisterHandle("/login", handle.Login)
	http.RegisterHandle("/logout", handle.Logout)
	http.RegisterHandle("/signup", handle.Signup)
	http.RegisterHandle("/signup_account", handle.SignupAccount)
	http.RegisterHandle("/authenticate", handle.Authenticate)

	// defined in route_thread.go
	http.RegisterHandle("/thread/new", handle.NewThread)
	http.RegisterHandle("/thread/create", handle.CreateThread)
	http.RegisterHandle("/thread/post", handle.PostThread)
	http.RegisterHandle("/thread/read", handle.ReadThread)
}

func initCfg() {
	// 首先完成配置项的注册
	cfg.RegisterCfg("test", "test", "string")
	cfg.RegisterCfg("Address", "0.0.0.0:8080", "string")
	cfg.RegisterCfg("ReadTimeout", 10, "int64")
	cfg.RegisterCfg("WriteTimeout", 600, "int64")
	cfg.RegisterCfg("Static", "public", "string")
	// log
	cfg.RegisterCfg("log.level", "info", "string")
	cfg.RegisterCfg("log.file_name", "log/minegin.log", "string")
	cfg.RegisterCfg("log.max_size_mb", 1, "int")
	cfg.RegisterCfg("log.max_file_num", 64, "int")
	cfg.RegisterCfg("log.max_file_day", 7, "int")
	cfg.RegisterCfg("log.compress", false, "bool")
	cfg.RegisterCfg("log.stdout", true, "bool")
	cfg.RegisterCfg("log.only_stdout", false, "bool")

	// 之后再进行初始化
	err := cfg.Init("")
	if err != nil {
		panic(fmt.Sprintf("init cfg failed: %s", err))
	}

	// 初始化结束后配置文件正常存取
	fmt.Printf("get cfg %s\n", cfg.Get("test", false))
	fmt.Printf("get cfg %v\n", cfg.Get("address", false))
	fmt.Printf("get cfg %v\n", cfg.Get("ReadTimeout", false))
	fmt.Printf("get cfg %v\n", cfg.Get("WriteTimeout", false))
	fmt.Printf("get cfg %v\n", cfg.Get("Static", false))
}

func initLog() {
	log.InitLogger(
		cfg.GetString("log.file_name"),
		cfg.GetString("log.level"),
		cfg.GetInt("log.max_size_mb"),
		cfg.GetInt("log.max_file_num"),
		cfg.GetInt("log.max_file_day"),
		cfg.GetBool("log.compress"))
}
