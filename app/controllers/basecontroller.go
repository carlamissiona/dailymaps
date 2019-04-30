package controllers
import (
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
	"github.com/kataras/golog"
	"github.com/davecgh/go-spew/spew"
)

type IUtil interface  {
  IsLogin() bool
  GetCredential(user string, pass string) bool
  Page404()   bool
  PageError() bool
  SetCookie() bool
	ConfigSession() bool
}


type HttpUtil struct {
    username       string
    cookiename     string
    pgsess         *sessions.Session
		contx					 *iris.Context
}



func (s *HttpUtil) IsLogin() bool {

		str := spew.Sdump(s.contx)
		golog.Info("Object httputil %v",str)

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

func (s *HttpUtil) GetCredential(user string, pass string)  bool {

    return true
}


func (s *HttpUtil) Page404() bool {

    return true
}


func (s *HttpUtil) PageError() bool {

    return true
}
