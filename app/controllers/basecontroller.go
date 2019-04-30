package controllers
import (
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
	"github.com/kataras/golog"
	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris/mvc"
	"github.com/go-xorm/xorm"
 _"github.com/go-sql-driver/mysql"

)

type IUtil interface  {
  IsLogin() bool
  PullCredential(user string, pass string) bool
  Page404()   bool
  PageError() mvc.View
  SetCookie() bool
	ConfigSession() bool
}


type HttpUtil struct {
    username       string
    cookiename     string
    pgsess         *sessions.Session
		contx					 *iris.Context
		dbhelper		   *xorm.Engine
}



func (s *HttpUtil) IsLogin() bool {

		str := spew.Sdump(s.contx)
		str2 := spew.Sdump(s.pgsess)
		golog.Info("Object httputil %v",str)
		golog.Info("Object httputil %v",str2)

    return true
}

//
// func (s *HttpUtil) DoLogout(ctx iris.Context) bool {
//
// 	if auth, _ := s.pgsess.GetBoolean("dmlogin"); !auth {
// 		ctx.StatusCode(iris.StatusForbidden)
// 		return true
// 	}
// 	// do logout
// 	s.pgsess.Delete("dmlogin");
// 	return true
//
//
// }

// func  (s *HttpUtil) ConfigSession(ctx iris.Context)  bool {
//
// 	  s.pgsess = sessions.New(sessions.Config{
// 	    Cookie: "dailymapsid",
// 	    Expires: time.Hour * 2,
// 		})
//
// 		str := spew.Sdump(s)
// 		golog.Info("Object httputil %v",str)
//
// 		s.pgsess.Start(ctx).Set("dmlogin", "true_auth")
// 		return false
//
// }

func (s *HttpUtil) PullCredential(user string, pass string)  bool {

    return true
}


func (s *HttpUtil) Page404() bool {

    return true
}


func (s *HttpUtil) PageError() mvc.View {

	return mvc.View {
		Name: "error.html",

	}
}
