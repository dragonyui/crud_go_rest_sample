package controllers

import (
	"github.com/dragonyui/router"
	"github.com/kataras/iris/context"
	"github.com/dragonyui/crud_go_rest_sample/services"

	"github.com/dragonyui/config"
	"github.com/dragonyui/crud_go_rest_sample/models"
	"github.com/kataras/iris"
)

func init()  {
	product:=router.CreateNewControllerInstance("product","product")
	product.Get("/{id:int64}",ProductGet)
	product.Put("/",ProductInsert)
	product.Post("/",ProductUpdate)
	product.Get("/all",ProductGetAll)
	product.Delete("/{id:int64}",ProductDel)

}
func ProductGet(ctx context.Context){
	id,err:=ctx.Params().GetInt64("id")
	if config.FancyHandleError(err){
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	svc:=services.ProductDao{}
	err=svc.Get(id)
	if err != nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	ctx.JSON(svc.Product)
}

func ProductDel(ctx context.Context){
	id,err:=ctx.Params().GetInt64("id")
	if config.FancyHandleError(err){
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	svc:=services.ProductDao{models.Product{Id:id}}
	err=svc.Delete()
	if err != nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	ctx.StatusCode(iris.StatusAccepted)
}
func ProductInsert(ctx context.Context)  {
	var product models.Product
	err:=ctx.ReadJSON(&product)
	if config.FancyHandleError(err){
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	svc:=services.ProductDao{product}

	err=svc.Insert()
	if err != nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	ctx.StatusCode(iris.StatusCreated)
}

func ProductUpdate(ctx context.Context)  {
	var product models.Product
	err:=ctx.ReadJSON(&product)
	if config.FancyHandleError(err){
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	svc:=services.ProductDao{product}
	err=svc.Update()

	if err != nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	ctx.StatusCode(iris.StatusAccepted)
}
func ProductGetAll(ctx context.Context){
	svc:=services.ProductDao{}
	all,err:=svc.GetAll()
	if err != nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(router.GetErrorResponse(err.Error()))
		return
	}
	ctx.JSON(all)
}
