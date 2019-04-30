package controllers


import (
  "github.com/kataras/golog"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
  _ "fmt"
)

type MapsController struct {
    HttpUtil
    isCacheable bool

}

func(c *MapsController) Get() mvc.Result {
	// set the model and render the view template.

 c.cookiename = "hihi"
 c.isCacheable = false

 golog.Infof("Maps controller %v",c.isCacheable)
 golog.Infof("Maps controller %v",c.cookiename)
 golog.Infof("Maps controller %v",c.IsLogin() )
 golog.Infof("Maps controller %v",c.cookiename)


	return mvc.View {
		Name: "maps.html",
    Data: iris.Map{"Title": "Maps List", "isLogin": c.cookiename },
	}


}
func(c *MapsController) GetNameUserBy(name string,id int) mvc.Result {
	// set the model and render the view template.
 c.cookiename = "hihi"
 c.isCacheable = false
 golog.Infof("Maps name%v",name) 
 // golog.Infof("Maps controller %v",c.cookiename)
 // golog.Infof("Maps controller %v",c.isCacheable)
 // golog.Infof("Maps controller %v",c.IsLogin() )
 // golog.Infof("Maps controller %v",c.cookiename)


	return mvc.View {
		Name: "maps2.html",

	}



}
