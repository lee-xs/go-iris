package main

import (
	"go-iris/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main(){
	app := iris.New()

	//　从　"./views"　目录下加载扩展名是".html" 　的所有模板，
	//　并使用标准的　`html/template`　 包进行解析。
	app.RegisterView(iris.HTML("view", ".html"))

	//静态资源访问
	app.HandleDir("/static", "static")

	mvc.Configure(app.Party("/"), func(a *mvc.Application) {
		a.Party("/").Handle(new(controllers.RouteController))
		a.Party("/user").Handle(new(controllers.UserController))
		a.Party("/hello").Handle(new(controllers.HelloController))
	})




	app.Run(iris.Addr(":2922"))
}
