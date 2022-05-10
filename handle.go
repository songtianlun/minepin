package main

import (
	"fmt"
	"minepin/com/cfg"
	"minepin/com/http"
)

func initHandle() {
	// static file
	http.RegisterFile("/static/", cfg.GetString("Static"), true)

	// index
	http.RegisterHandle("/", index)

	// error
	http.RegisterHandle("/err", err)

	// defined in route_auth.go
	http.RegisterHandle("/login", login)
	http.RegisterHandle("/logout", logout)
	http.RegisterHandle("/signup", signup)
	http.RegisterHandle("/signup_account", signupAccount)
	http.RegisterHandle("/authenticate", authenticate)

	// defined in route_thread.go
	http.RegisterHandle("/thread/new", newThread)
	http.RegisterHandle("/thread/create", createThread)
	http.RegisterHandle("/thread/post", postThread)
	http.RegisterHandle("/thread/read", readThread)
}

func initCfg() {
	// 首先完成配置项的注册
	cfg.RegisterCfg("test", "test", "string")
	cfg.RegisterCfg("Address", "0.0.0.0:8080", "string")
	cfg.RegisterCfg("ReadTimeout", 10, "int")
	cfg.RegisterCfg("WriteTimeout", 600, "int")
	cfg.RegisterCfg("Static", "public", "string")

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
