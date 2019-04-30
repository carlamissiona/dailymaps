package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
  "github.com/kataras/golog"
	"dailymaps/app/controllers"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-xorm/xorm"
	"time"
)


func GetAppRoutes() *iris.Application {
	app  := iris.New()

	sess := *sessions.New(sessions.Config{
	    Cookie: "dailymapsid",
	    Expires: time.Hour * 2,
		})


  app.Use(dbrunner)

	golog.Infof("Session controller %v", sess)
	app.RegisterView(iris.HTML("app/views", ".html").Layout("templates/layout.html"))


  home := mvc.New(app.Party("/")); home.Handle(new(controllers.HomeController))

	maps := mvc.New(app.Party("/maps")); maps.Handle(new(controllers.MapsController));	maps.Register( sess.Start, )

  rest := mvc.New(app.Party("/rest")); rest.Handle(new(controllers.ApiController))




	return app
}

func dbrunner( ctx iris.Context )  {

	sess := *sessions.New(sessions.Config{
	    Cookie: "dailymapsid",
	    Expires: time.Hour * 2,
		})

	mysqlgn, err := xorm.NewEngine("mysql", "root:@127.0.0.1:3036/DAILYMAPS_DATA")
	iris.RegisterOnInterrupt(func() {
		mysqlgn.Close()
	})
	golog.Infof("Session controller has error ? %v", err)

	// if there was an issue opening the connection, send a 500 error


	sess.Start(ctx).Set("dbservice", mysqlgn)

	str := spew.Sdump(sess)
	golog.Infof(str)

	golog.Infof("Session controller session %v", str)
	// move on now that it is exposed
	ctx.Next()

}
