package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
  _ "github.com/kataras/golog"
	"dailymaps/app/controllers"
)


func GetAppRoutes() *iris.Application {
	app  := iris.New()


	app.RegisterView(iris.HTML("app/views", ".html").Layout("templates/layout.html"))


  // HomeController no session
  home := mvc.New(app.Party("/")); home.Handle(new(controllers.HomeController))

  // MapsController has session
  maps := mvc.New(app.Party("/maps")); maps.Handle(new(controllers.MapsController))

  // MapsController has session
  maps := mvc.New(app.Party("/rest")); maps.Handle(new(controllers.RestController))




	return app
}
