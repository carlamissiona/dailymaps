package controllers


import (
  "github.com/kataras/golog"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
  _ "fmt"

  _ "github.com/go-sql-driver/mysql"

)

type MapsController struct {
    HttpUtil
    isCacheable bool
    //  Service services.MapsHelper

}

func(c *MapsController) Get() mvc.Result {
	// set the model and render the view template.

 c.cookiename = "hihi"
 c.isCacheable = false

 golog.Infof("Maps controller %v",)
 golog.Infof("Maps controller  db orm %v",c.dbhelper)

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

c.IsLogin()
	return mvc.View {
		Name: "maps2.html",

	}



}


//***
//  /maps                - display available maps
//  /maps/user/{id}      - id cookie get maps by user id   isloggedin
//  /
//
//****
