package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"dailymaps/app/controllers"
	"github.com/kataras/golog"
) 


func GetAppRoutes() *iris.Application {
	app  := iris.New()
	app.RegisterView(iris.HTML("app/views", ".html")).Layout("shared/layout.html"))
  golog.Info("hi get app")

  // HomeController no session
  home := mvc.New(app.Party("/")); home.Handle(new(controllers.HomeController))

  // MapsController has session
  maps := mvc.New(app.Party("/maps")); maps.Handle(new(controllers.MapsController))

	return app
}
