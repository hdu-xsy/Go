package Entity

import (
	"time"
	"github.com/kataras/iris/sessions"
)

type AdminUser struct {
	Id       int64`xorm:"pk"`
	Account  string`xorm:"unique"`
	Password string
}
type UserData struct {
	Id       int64`xorm:"pk"`
	Username string`xorm:"unique"`
	Password string
}
type OnlineUser struct {
	Id        int64
	Username  string`xorm:"unique"`
	Logintime time.Time`xorm:"created"`
}
type Article struct {
	Id        int64`xorm:"pk"`
	User      int64`xorm:"unique"`
	Time	  time.Time`xorm:"created"`
	Title	  string`xorm:"unique"`
	Menu	  string`xorm:"unique"`
	Classify  string`xorm:"unique"`
	Content   string`xorm:"varchar(5000)"`
}
type Menu struct {
	Id		  int64`xorm:"pk"`
	Name	  string`xorm:"unique"`
}
var (
	CookieNameForSessionID = "mycookiesessionnameid"
	Sess                   = sessions.New(sessions.Config{Cookie: CookieNameForSessionID})
)