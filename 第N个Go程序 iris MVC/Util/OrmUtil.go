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
func GetOnlineUser(app iris.Application) error{
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(OnlineUser))
	return err
}
func GetAdminUser(app iris.Application) error{
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(AdminUser))
	return err
}
func GetUserData(app iris.Application) error{
	orm := Getorm(app)
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err := orm.Sync2(new(UserData))
	return err
}
func Getorm(app iris.Application) *xorm.Engine {
	orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	return orm
}