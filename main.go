package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris"
  "dailymaps/app/routes"
	"github.com/kataras/golog"
)


func startApp() *iris.Application {
	app := iris.New()
  app = routes.GetAppRoutes()
	app.StaticWeb("/templates", "./app/views/templates")

	str := spew.Sdump(app)

	golog.Infof(str)


	return app
}

func main() {

	app := startApp()

  golog.Info(" hi hiwwww hi app ")
  golog.Infof("%d" ,app)

  app.Run(iris.Addr(":8090"))

}
