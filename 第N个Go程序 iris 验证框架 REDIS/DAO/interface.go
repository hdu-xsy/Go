package DAO

import(
	"github.com/kataras/iris"
	"../Entity"
	"../Util"
)
var app = iris.New()
var orm = Util.Getorm(*app)
var err1 = Util.GetUserData(*app)
var err2 = Util.GetAdminUser(*app)
var err3 = Util.GetOnlineUser(*app)
var err4 = Util.GetArticle(*app)
var err5 = Util.GetMenu(*app)

type UserDataDAOInterface interface {
	FindAll() []Entity.UserData
	Get(userdata Entity.UserData) (bool,error,Entity.UserData)
	Insert(Userdata Entity.UserData) (int64,error)
	Delete(userdata Entity.UserData) (int64,error)
	Motify(userdata Entity.UserData) (int64,error)
}

type AdminUserDAOInterface interface {
	Get(admin Entity.AdminUser) (bool,error,Entity.AdminUser)
}

type OnlineUserDAOInterface interface {
	Get(onlineuser Entity.OnlineUser) (bool,error,Entity.OnlineUser)
	Insert(onlineuser Entity.OnlineUser) error
	Delete(onlineuser Entity.OnlineUser) error
}
type ArticleDAOInterface interface {
	Get(article Entity.Article) (bool,error,Entity.Article)
	FindAll(id string) []Entity.Article
	Insert(article Entity.Article) (int64,error)
	OrderByTime() []Entity.Article
	FindAllA() []Entity.Article
	Update(article Entity.Article) (int64,error)
}
type MenuDAOInterface interface {
	Get(menu Entity.Menu) (bool,error,Entity.Menu)
}