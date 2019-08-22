package controllers

import (
	"github.com/dragonyui/router"
	"github.com/kataras/iris/context"
)

func init()  {
	index:=router.CreateNewControllerInstance("index","")
	index.Get("/",Index)
}

type AppCredit struct {
	AppName string `json:"app_name"`
	Creator string `json:"creator"`
}
func Index(ctx context.Context)  {
	ctx.JSON(AppCredit{AppName:"CRUD GO REST EXAMPLE",Creator:"Ricky Oktavianus Lazuardy"})
}
