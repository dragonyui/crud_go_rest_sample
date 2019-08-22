package main

import (
	"log"
	_ "github.com/dragonyui/crud_go_rest_sample/controllers"
	"github.com/kataras/iris"
	"github.com/dragonyui/router"
	"github.com/dragonyui/config"
)

func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app:=iris.New()
	router.ProcessController(app)
	app.Run(iris.Addr(":"+config.GetCurrentEnvValueFromConfig("app.port")))
}
