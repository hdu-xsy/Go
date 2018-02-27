## Go IRIS框架学习笔记
### 生成网页与注册路由的一般方法
- main函数中
```Golang
app.RegisterView(iris.HTML("html",".html").Reload(true))
app.StaticWeb("/js", "./js") // serve our custom javascript code
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
	//查找全部
	var users []AdminUser
	orm.Find(&users)
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
### mvc
### 认证框架
### 文件上传
### 测试
