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
	SessionId string
}
type OnlineUser struct {
	Uid       int64
	Username  string`xorm:"unique"`
	Logintime time.Time`xorm:"created"`
}
type Article struct {
	Id        int64`pk`
	User      int64
	Time	  time.Time`xorm:"created"`
	Title	  string
	Menu	  string
	Classify  string
	Content   string
}
type Menu struct {
	Id		  int64`xorm:"pk"`
	Name	  string`xorm:"unique"`
}
type Comment struct {
	Id		  int64`pk`
	Article   int64
	Floor     int64
	User      int64
	Time      time.Time`xorm:"created"`
	Content   string
}
type Entity struct {
	AdminUser 		AdminUser
	AdminUserList   []AdminUser
	UserData  		UserData
	UserDataList	[]UserData
	OnlineUser  	OnlineUser
	OnlineUserList  []OnlineUser
	Article 		Article
	ArticleList		[]Article
	Menu 			Menu
	MenuList		[]Menu
	Comment 		Comment
	CommentList		[]Comment
}
type File struct {
	Id			int64`pk`
	Name 		string
}
var (
	CookieNameForSessionID = "mycookiesessionnameid"
	Sess                   = sessions.New(sessions.Config{Cookie: CookieNameForSessionID})
)