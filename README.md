## Go IRIS框架学习笔记

> 当前网上关于Golang的IRIS框架的教程极极极极少，有参考价值的也就官方文档，可惜官网文档*除 了 代 码 什 么 都 没 有*，而且 还 *没 有 中 文 版*，竟 然 还 是 *按 照 首 字 母 排 得 序*，但是比起热门的GIN、BEEGO、马卡龙框架，IRIS还是有自己的特点的，于是在此整理方便学习和以后使用（以官方文档为主）  []内内容为备注(类型) [...]为方法体

### 安装iris和导入iris
```
命令行:go get github.com/kataras/iris
```
```
import (
	"github.com/kataras/iris"
)
```
### 生成网页与注册路由的一般方法
- main函数中
```Golang
app.RegisterView(iris.HTML("/html[directory]",".html[extension]").Reload(true))
app.StaticWeb("/js[request path]", "./js[system path]") // serve our custom javascript code
app.Get("/",index[funcName])  //路由
app.Run(iris.Addr(":80"))  //端口
```
- 路由函数中
```
func funcName(ctx iris.Context){
    if err := ctx.View("index.html");err!=nil {
    		//错误处理[...]
    }
}
```
- 错误页面
```
app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context){
		ctx.View("404.html")
})
```
### 读取POST请求和ORM基本操作
```
type 表名 struct {
	Id       int64`pk`
	字段名   string`unique`
	Password string`notnull`
}
// 生成的表会自动转化格式 如 AdminUser → admin_user
// 下面例子假定表名为admin_user
func Funcname(ctx iris.Context) {
	user := AdminUser{}
	//  读取请求
	err := ctx.ReadForm(&user)
	if err !=nil {
		//[..]
	}
	//  ORM
	//  mysql
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
	//查找全部
	var users []AdminUser
	orm.Find(&users)
	//详细操作日后再来学习
    //官方文档http://www.xorm.io/docs/
}
```
## 带参数的路径
```
app.Get("/{xxx}/{namedRoute}", func(ctx iris.Context) {
        xxx := ctx.Params().Get("xxx")
		routeName := ctx.Params().Get("namedRoute")
		//[...]
})
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
### session
```
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
session := sess.Start(ctx)  //ctx iris.Context
session.Set（"key", value)
session.Get("key")  //若使用Get方法需要强转 指定：GetString GetInt GetBoolean GetAll GetFloat ...
session.Delete("key")   //删除
session.Clear()     //清空
session.Destroy()   //销毁
```
### WebSokect
- iris的websocket
    - html
    ```
    <form name = "form1" method="post">
        输入聊天信息:<input id="msg" name="msg" type="text" size="40" class="form-control">
    </form>
    <button name="websocket" id="websocket" class="btn btn-default" onclick="send()">发送</button>
    <script src="/iris-ws.js"></script>
    <div id="message" name="message"></div>
    
    ```
    - js
    ```
    var scheme = document.location.protocol == "https:" ? "wss" : "ws";
    var port = document.location.port ? (":" + document.location.port) : "";
    // see app.Get("/echo", ws.Handler()) on main.go
    var wsURL = scheme + "://" + document.location.hostname + port+"/echo";
    var input = document.getElementById("msg");
    var output = document.getElementById("message");
    // Ws comes from the auto-served '/iris-ws.js'
    var socket = new Ws(wsURL)
    socket.OnConnect(function () {
        //连接服务器[...]
    });
    socket.OnDisconnect(function () {
        //无法连接服务器[...]
    });
    // read events from the server
    socket.On("消息名称", function (msg) {
        //读取事件[...]
        addMessage(msg+"<br/>");
    });
    function send() {
        addMessage("Me: " + input.value); // write ourselves
        socket.Emit("消息名称", "消息内容");// send chat event data to the websocket server
        input.value = ""; // clear the input
    }
    function addMessage(msg) {
        output.innerHTML += msg + "<br>";
    }
    ```
    - Go
    ```
    package main

    import (
    	"github.com/kataras/iris"
    	"github.com/kataras/iris/websocket"
    	"github.com/go-xorm/xorm"
    )
    
    func setupWebsocket(app *iris.Application) {
    	// create our echo websocket server
    	ws := websocket.New(websocket.Config{
    		ReadBufferSize:  1024,
    		WriteBufferSize: 1024,
    	})
    	ws.OnConnection(handleConnection)
    	// register the server on an endpoint.
    	app.Get("/echo", ws.Handler())
    	// see html script tags, this path is used.
    	app.Any("/iris-ws.js", func(ctx iris.Context) {
    		ctx.Write(websocket.ClientSource)
    	})
    }
    
    func handleConnection(c websocket.Connection) {
    	c.On("chat", func(msg string) {
    		// fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
    		// Write message back to the client message owner with: c.Emit("消息名称", msg)
    		// 返回所有人 websocket.All
    		// Write message to all except this client with:
    		c.To(websocket.Broadcast).Emit("消息名称", msg)
    	})
    	c.OnDisconnect(func() {
    	    //断开连接操作[...]
    	})
    }
    func main() {
        setupWebsocket(app)
    }
    ```
- 官方的websocket
- 第三方websocket
### 模板
- 上下文视图数据
```
app.Get("/路由", func(ctx iris.Context) {
		ctx.ViewData("Key", "Value")
		if err := ctx.View("index.html"); err != nil {
			//[...]
		}
})
//HTML中使用{{.Key}}获取
```
- 嵌入模板
```
//layout页面中加入{{ yield }}使用模板
//HTML中使用{{funcName 参数}}使用函数 如下例中可使用 {{ greet "string" }}
//{{ render "路径" }} 把别的html文件的内容拿过来
func main() {
	app := iris.New()
	//全局模板 所有页面适用
	tmpl := iris.HTML("./templates[html文件夹路径]", ".html")
	tmpl.Layout("layouts/layout.html[layout文件路径]")
	tmpl.AddFunc("greet", func(s string) string {
		//[..]
	})
	tmpl.Binary(Asset, AssetNames) // <-- IMPORTANT
	app.RegisterView(tmpl)
	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// 不使用全局模板
	app.Get("/nolayout", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// 局部模板 不受全局模板影响
	my := app.Party("/my").Layout("layouts/mylayout.html") { 
		my.Get("/", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
		my.Get("/other", func(ctx iris.Context) {
			ctx.View("page2.html")
		})
	}
	app.Run(iris.Addr(":8080"))
}
```
### mvc
### 认证框架
### 文件上传
### 测试
