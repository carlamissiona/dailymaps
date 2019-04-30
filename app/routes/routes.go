package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
  "github.com/kataras/golog"
	"dailymaps/app/controllers"
	"time"
)


func GetAppRoutes() *iris.Application {
	app  := iris.New()



  sess := *sessions.New(sessions.Config{
	    Cookie: "dailymapsid",
	    Expires: time.Hour * 2,
		})

	golog.Infof("Session controller %v", sess)
	app.RegisterView(iris.HTML("app/views", ".html").Layout("templates/layout.html"))


  home := mvc.New(app.Party("/")); home.Handle(new(controllers.HomeController))

  maps := mvc.New(app.Party("/maps")); maps.Handle(new(controllers.MapsController))

 // maps.Get("/", func(ctx iris.Context) {
 //    // 		ctx.View("page1.html")
 //    // 	})



  rest := mvc.New(app.Party("/rest")); rest.Handle(new(controllers.ApiController))




	return app
}
