## Go IRIS框架学习笔记
### 生成网页与注册路由
- main函数中
```Golang
app.RegisterView(iris.HTML("html",".html").Reload(true))
app.Get("/",index)  //路由
app.Run(iris.Addr(":80"))  //端口
```
- 路由函数中
```
if err := ctx.View("index.html");err!=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
}
```
### 读取POST请求和ORM基本操作
```
type 表名 struct {
	Id       int
	字段名   string
	Password string
}
// 用XORM生成的表的Id字段会自动增长
// 下面例子假定表名为AdminUser
func Funcname(ctx iris.Context) {
	user := AdminUser{}
	//  读取请求
	err := ctx.ReadForm(&user)
	if err !=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		ctx.Redirect("/404")
	}
	//  ORM
	orm,err := xorm.NewEngine("mysql", "用户名:密码@/库名?charset=utf8")
	//  不同数据库操作不同 具体参考文档
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
	adminuser := AdminUser{Account: user.Account}
	//查找操作
	if ok, _ := orm.Get(&adminuser); ok { //orm.Id(xxx).Get(&) 或 orm.Where("属性=?",value).Get(&)
	} else {
	}
	//插入操作
	adminuser.Account = "AAAAAAA"
	orm.Insert(adminuser)
	//修改操作
	orm.Id(xx).Update(adminuser)
	//删除操作
	orm.delete(adminuser)
	//详细操作日后再来学习
}
```
## Ajax
- 前端页面
```
<script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
<script type="text/javascript">
    $(function() {
        $("#按钮Id").click(function() {
            //提交的参数
            var params = {
                Account : $("#input的Id").val(),
                Password : $("#input的Id").val()
            };
            $.post("/路由",params,function (参数) {
                要执行的func(参数);
        });
    });
</script>
```
- Go
```
//接收post
user := AdminUser{}
err := ctx.ReadForm(&user)
if err !=nil {
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.WriteString(err.Error())
	ctx.Redirect("/404")
}
//要执行的操作func(user)
ctx.WriteString(返回值)
return
```
