package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-ignite/ignite/config"
	"github.com/go-ignite/ignite/controllers"
)

var (
	confPath    = flag.String("c", "", "config file")
	versionFlag = flag.Bool("v", false, "version")
	version     = "unknown"
)

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		return
	}
	config.Init(*confPath)

	initRouter()
}

func initRouter() {
	r := gin.Default()
	r.Static("/static", "./static")

	mainRouter := &controllers.MainRouter{}
	mainRouter.Initialize(r)
}
