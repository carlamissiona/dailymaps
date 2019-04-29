package controllers


import (
  "github.com/kataras/golog"
	"github.com/kataras/iris/mvc"
  

)

type ApiController struct {
		HttpUtil

}

func(c *HomeController) Get() mvc.Result {
	// set the model and render the view template.
  golog.Info("controller")
	return mvc.View {
		Name: "home.html",

	}

}
