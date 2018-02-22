package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	//"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)
func checkError(err error) {
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(-1)
	}
}
type AdminUser struct {
	Id       int
	Account  string
	Password string
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	t, _ := template.ParseFiles("main/adminlogin.html")
	t.Execute(w,nil)
}
func adminlogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("main/adminlogin.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		var username = r.Form["admin_account"][0]
		var password = r.Form["admin_password"][0]
		fmt.Println(username)
		fmt.Println(password)
		o := orm.NewOrm()
		user := AdminUser{Account:username}
		err := o.Read(&user,"account")
		if err == orm.ErrNoRows {
			t, _ := template.ParseFiles("main/adminlogin.html")
			t.Execute(w,nil)
			fmt.Println("ErrNoRows")
		} else if user.Password != password {
			t, _ := template.ParseFiles("main/adminlogin.html")
			t.Execute(w,nil)
			fmt.Println("密码错误")
		} else {
			t, _ := template.ParseFiles("main/backend.html")
			t.Execute(w,nil)
		}
		fmt.Println("username:", r.Form["admin_account"])
		fmt.Println("password:", r.Form["admin_password"])
	}
}
func main() {
	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	http.HandleFunc("/adminlogin", adminlogin) // 设置访问的路由
	err := http.ListenAndServe(":4567", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:Xsydx886.@/javaweb?charset=utf8", 30)
	// register model
	orm.RegisterModel(new(AdminUser))
	// create table
	orm.RunSyncdb("default", false, true)
}