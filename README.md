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
## 中间件生态系统
```
app.Use(func(ctx iris.Context){
	// ... any code here

	ctx.Next() //为了继续下一个处理程序，
	//如果缺少那么链处理程序中的下一个将不会被执行，
	//对于认证中间件很有用
})

//全局
//在任何路线之后或之前，将中间件预先添加到所有路线
app.UseGlobal(handler1, handler2, handler3)

//每个路由
app.Post("/login", authenticationHandler, loginPageHandler)

//每个派对（路线组）
users := app.Party("/users", usersMiddleware)
users.Get("/", usersIndex)

//每个子域
mysubdomain := app.Party("mysubdomain.", firstMiddleware)
mysubdomain.Use(secondMiddleware)
mysubdomain.Get("/", mysubdomainIndex)

//每个通配符，动态子域
dynamicSub := app.Party(".*", firstMiddleware, secondMiddleware)
dynamicSub.Get("/", func(ctx iris.Context){
	ctx.Writef("Hello from subdomain: "+ ctx.Subdomain())
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
- hero
    - 官方文档 https://github.com/shiyanhui/hero/blob/master/README_CN.md
    - iris中的实现
        1. HTML中编写模板（语法见下）
        2. 命令行中进入%GOPATH%/bin 执行hero -source="模板所在文件夹"
        3. main函数中
        
        ```
        //对于<%: func FuncName(变量名 []string, buffer *bytes.Buffer) %>
        app.Get("/路由", func(ctx iris.Context) {
    		var 变量名 = []string{
    		    //变量值
    		    //..
    		}
    		// Had better use buffer sync.Pool.
    		// Hero exports GetBuffer and PutBuffer for this.
    		//
    		// buffer := hero.GetBuffer()
    		// defer hero.PutBuffer(buffer)
    		buffer := new(bytes.Buffer)
    		template.FuncName(变量名, buffer)
    		if _, err := ctx.Write(buffer.Bytes()); err != nil {
    			log.Printf("ERR: %s\n", err)
    		}
		})
		
        ```
        ```
        //对于<%: func FuncName(变量名 []string, w io.Writer) (int, error)%>
        app.Get("/路由", func(ctx iris.Context) {
    		var 变量名 = []string{
                //
    		}
    		// using an io.Writer for automatic buffer management (i.e. hero built-in buffer pool),
    		// iris context implements the io.Writer by its ResponseWriter
    		// which is an enhanced version of the standard http.ResponseWriter
    		// but still 100% compatible.
    		template.UserListToWriter(FuncName, ctx)
    	})
        ```
    - 以下内容转自官方文档
    - Hero总共有九种语句，他们分别是：
    - 函数定义语句 `<%: func define %>`
      - 该语句定义了该模板所对应的函数，如果一个模板中没有函数定义语句，那么最终结果不会生成对应的函数。
      - 该函数最后一个参数必须为`*bytes.Buffer`或者`io.Writer`, hero会自动识别该参数的名字，并把把结果写到该参数里。
      - 例:
        - `<%: func UserList(userList []string, buffer *bytes.Buffer) %>`
        - `<%: func UserList(userList []string, w io.Writer) %>`
        - `<%: func UserList(userList []string, w io.Writer) (int, error) %>`
    
    - 模板继承语句 `<%~ "parent template" %>`
      - 该语句声明要继承的模板。
      - 例: `<%~ "index.html" >`
    
    - 模板include语句 `<%+ "sub template" %>`
      - 该语句把要include的模板加载进该模板，工作原理和`C++`中的`#include`有点类似。
      - 例: `<%+ "user.html" >`
    
    - 包导入语句 `<%! go code %>`
      - 该语句用来声明所有在函数外的代码，包括依赖包导入、全局变量、const等。
    
      - 该语句不会被子模板所继承
    
      - 例:
    
        ```go
        <%!
        	import (
              	"fmt"
            	"strings"
            )
    
        	var a int
    
        	const b = "hello, world"
    
        	func Add(a, b int) int {
            	return a + b
        	}
    
        	type S struct {
            	Name string
        	}
    
        	func (s S) String() string {
            	return s.Name
        	}
        %>
        ```
    
    - 块语句 `<%@ blockName { %> <% } %>`
    
      - 块语句是用来在子模板中重写父模中的同名块，进而实现模板的继承。
    
      - 例:
    
        ```html
        <!DOCTYPE html>
        <html>
            <head>
                <meta charset="utf-8">
            </head>
    
            <body>
                <%@ body { %>
                <% } %>
            </body>
        </html>
        ```
    
    - Go代码语句 `<% go code %>`
    
      - 该语句定义了函数内部的代码部分。
    
      - 例:
    
        ```go
        <% for _, user := userList { %>
            <% if user != "Alice" { %>
            	<%= user %>
            <% } %>
        <% } %>
    
        <%
        	a, b := 1, 2
        	c := Add(a, b)
        %>
        ```
    
    - 原生值语句 `<%==[t] variable %>`
    
      - 该语句把变量转换为string。
    
      - `t`是变量的类型，hero会自动根据`t`来选择转换函数。`t`的待选值有:
        - `b`: bool
        - `i`: int, int8, int16, int32, int64
        - `u`: byte, uint, uint8, uint16, uint32, uint64
        - `f`: float32, float64
        - `s`: string
        - `bs`: []byte
        - `v`: interface
    
        注意：
        - 如果`t`没有设置，那么`t`默认为`s`.
        - 最好不要使用`v`，因为其对应的转换函数为`fmt.Sprintf("%v", variable)`，该函数很慢。
    
      - 例:
    
        ```go
        <%== "hello" %>
        <%==i 34  %>
        <%==u Add(a, b) %>
        <%==s user.Name %>
        ```
    
    - 转义值语句 `<%= statement %>`
    
      - 该语句把变量转换为string后，又通过`html.EscapesString`记性转义。
      - `t`跟上面原生值语句中的`t`一样。
      - 例:
    
        ```go
        <%= a %>
        <%= a + b %>
        <%= Add(a, b) %>
        <%= user.Name %>
        ```
    - 注释语句 `<%# note %>`
      - 该语句注释相关模板，注释不会被生成到go代码里边去。
      - 例: `<# 这是一个注释 >`.

### mvc
- 基本
    - main函数中
    ```
    mvc.New(app.Party("/backend")).Handle(new(AdminLoginController))
    
    ```
    - Controller
    ```
    type UserLoginController struct {
	    Result string
    }
    func (c *UserLoginController) BeginRequest(ctx iris.Context) {
    	//[..]
    }
    func (c *UserLoginController) EndRequest(ctx iris.Context) {}
    func (c *UserLoginController) Get() mvc.View {
        //也可不使用返回值mvc.View使用hero模板的函数,Get(ctx iris.Context)
    	return mvc.View{
    		Name: "Name",
    		Data: iris.Map{
    			"Result": c.Result,
    		},
    	}
    }
    ```
    - HTML
    ```
    {{.Result}}
    ```
- BeforeActivation
```
func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // and any standard API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler", func (ctx iris.Context){})
}
```
> 通过`BeforeActivation`自定义事件回调，每个控制器将自定义控制器的结构方法注册为具有自定义路径的处理程序（即使使用正则表达式参数化路径）

### 认证框架
### 文件上传
### 测试
- GDB调试
- 测试用例
1. import "testing"包 文件以_test.go结尾 函数以Test开头
2. 案例
```
func Test_Name(t *testing.T) {
	t.Error("非预期的结果")
	t.Log("预期的结果")
	t.Fatal("致命错误")
}
```
3. go test -v Path
- 压力测试
- 基准测试
- 样本测试
- iris测试
```
func TestNewApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)
	e.GET("/").Expect().Status(httptest.StatusUnauthorized)
	e.GET("/admin").WithBasicAuth("myusername", "mypassword").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin myusername:mypassword")
	// with invalid basic auth
	e.GET("/admin/settings").WithBasicAuth("invalidusername", "invalidpassword").
		Expect().Status(httptest.StatusUnauthorized)
}
`go test -v`
```
### Redis
