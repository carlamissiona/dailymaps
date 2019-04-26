package controllers


import (
  "github.com/kataras/golog"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
  _ "fmt"
)

type MapsController struct {
		Helpers  BaseController

}

func(c *MapsController) Get() mvc.Result {
	// set the model and render the view template.

  golog.Infof("Maps controller %v",true)
	return mvc.View {
		Name: "maps.html",
    Data: iris.Map{"Title": "Maps List", "isLogin":"true" },
	}



}
