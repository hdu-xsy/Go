package DAO

import(
	"github.com/kataras/iris"
	"../Entity"
	"../Util"
)
var app = iris.New()
var orm = Util.Getorm(*app)
var UserDataErr = Util.GetUserData(*app)
var AdminUserErr = Util.GetAdminUser(*app)
var OnlineUserErr = Util.GetOnlineUser(*app)
var ArticleErr = Util.GetArticle(*app)
var MenuErr = Util.GetMenu(*app)
var CommentErr = Util.GetComment(*app)
var FileErr = Util.GetFile(*app)

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
	Count() map[string]int64
	FindByClassify(classify string) []Entity.Article
	GetClassify() []Entity.Article
}
type MenuDAOInterface interface {
	Get(menu Entity.Menu) (bool,error,Entity.Menu)
	GetAll() []Entity.Menu
	Insert(menu Entity.Menu) (int64,error)
	Delete(menu Entity.Menu) (int64,error)
	Motify(menu Entity.Menu) (int64,error)
}
type CommentDAOInterface interface{
	FindAll(article string) []Entity.Comment
	Insert(comment Entity.Comment) (int64,error)
	OrderByTime() []Entity.Comment
}
type FileDAOInterface interface {
	GetAll() []Entity.File
	Insert(file Entity.File) (int64,error)
	GetName(id int64) string
}
type RedisDAOInterface interface {
	Set(key string,value string) interface{}
	Get(key string) string
}
