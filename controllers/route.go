package controllers

import "github.com/kataras/iris"

type RouteController struct {}

func (r *RouteController) Get(ctx iris.Context) {
	ctx.Redirect("/index")
}

func (r *RouteController) GetIndex(ctx iris.Context) {
	ctx.ViewData("msg", "主页")
	ctx.View("index.html")
}

func (r *RouteController) GetRegister(ctx iris.Context) {
	ctx.View("register.html")
}

func (r *RouteController) GetLogin(ctx iris.Context) {
	ctx.View("login.html")
}