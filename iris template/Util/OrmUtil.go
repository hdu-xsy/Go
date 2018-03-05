package Util

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"time"
)

type OnlineUser struct {
	Id        int64
	Username  string`unique`
	Logintime time.Time`created`
}
type AdminUser struct {
	Id       int64`pk`
	Account  string`unique`
	Password string
}
type UserData struct {
	Id       int64`pk`
	Username string`unique`
	Password string
}
func GetOnlineUser(app iris.Application) *xorm.Engine{
	orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	if err = orm.Sync2(new(OnlineUser));err !=nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	return orm
}
func GetAdminUser(app iris.Application) * xorm.Engine {
	orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err = orm.Sync2(new(AdminUser))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	return orm
}
func GetUserData(app iris.Application) *xorm.Engine {
	orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	if err = orm.Sync2(new(UserData));err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	return orm
}