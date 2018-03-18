package Util

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"../Entity"
)

func GetOnlineUser(app iris.Application) error{
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(Entity.OnlineUser))
	return err
}
func GetAdminUser(app iris.Application) error{
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(Entity.AdminUser))
	return err
}
func GetUserData(app iris.Application) error{
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(Entity.UserData))
	return err
}
func Getorm(app iris.Application) *xorm.Engine {
	orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	return orm
}
func GetMenu(app iris.Application) error {
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(Entity.Menu))
	return err
}
func GetArticle(app iris.Application) error {
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(Entity.Article))
	return err
}