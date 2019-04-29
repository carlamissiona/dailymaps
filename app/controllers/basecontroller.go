package controllers
import (
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
)

type IUtil interface  {
  IsLogin() bool
  GetCredential(user string, pass string) bool
  Page404()   bool
  PageError() bool
  SetCookie(c *iris.Context) bool
}


type HttpUtil struct {
    username       string
    cookiename     string
    pgsession        *sessions.Sessions
}



func (s HttpUtil) IsLogin() bool {

    s := sess.Start(ctx)
		s.Set("dmlogin", "trueauth")
    return true
}

func DoLogout(ctx iris.Context) bool {
	session := sess.Start(ctx)

	// Revoke users authentication
	session.Set("authenticated", false)
}

func (s HttpUtil) ConfigCookie()  bool {

  s.sess := sessions.New(sessions.Config{
    Cookie: "dailymapsid",
    Expires: time.Hour * 2,
	})

}
func (s HttpUtil) GetCredential(user string, pass string)  bool {

    return true
}


func (s HttpUtil) Page404() bool {

    //s.cookiename =
    return true
}


func (s HttpUtil) PageError() bool {

    //s.cookiename =
    return true
}
