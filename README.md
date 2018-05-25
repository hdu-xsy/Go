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
app.Favicon("[path]")	//图标
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
### 图片
- IRIS方法  
```
app.StaticWeb("/static","./static") //前者为访问目录 后者为文件目录

```
- Go方法  
```
<img  src=\"/static/IMG.jpg\">
http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
})
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
app.DoneGlobal(after)

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
## 路由状态
```
none := app.None("...", func(ctx iris.Context) {
	//[...]
})
app.Get("/change", func(ctx iris.Context) {

	if none.IsOnline() {
		none.Method = iris.MethodNone
	} else {
		none.Method = iris.MethodGet
	}

	// refresh re-builds the router at serve-time in order to be notified for its new routes.
	app.RefreshRouter()
})
```
## 重写Context
```
type MyContext struct {
	// Optional Part 1: embed (optional but required if you don't want to override all context's methods)
	context.Context // it's the context/context.go#context struct but you don't need to know it.
}

var _ context.Context = &MyContext{} // optionally: validate on compile-time if MyContext implements context.Context.

// The only one important if you will override the Context
// with an embedded context.Context inside it.
// Required in order to run the handlers via this "*MyContext".
func (ctx *MyContext) Do(handlers context.Handlers) {
	context.Do(ctx, handlers)
}

// The second one important if you will override the Context
// with an embedded context.Context inside it.
// Required in order to run the chain of handlers via this "*MyContext".
func (ctx *MyContext) Next() {
	context.Next(ctx)
}

// Override any context's method you want...
// [...]

func (ctx *MyContext) HTML(htmlContents string) (int, error) {
	//[...]
}

func main() {
	app := iris.New()
	app.ContextPool.Attach(func() context.Context {
		return &MyContext{
			Context: context.NewContext(app),
		}
	})
	app.Handle("GET", "/", recordWhichContextJustForProofOfConcept, func(ctx context.Context) {
		// use the context's overridden HTML method.
		ctx.HTML([...])
	})
	app.Run(iris.Addr(":8080"))
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
### Writer
- write-gzip
```
ctx.GzipResponseWriter().WriteString("Hello World!")

or

ctx.WriteGzip([]byte("Hello World!"))
ctx.Header("X-Custom","Headers can be set here after WriteGzip as well, because the data are kept before sent to the client when using the context's GzipResponseWriter and ResponseRecorder.")
```
- Binary
```
ctx.Binary([]byte("Some binary data here."))
```
- Text
```
ctx.Text("Plain text here")
```
- Json
```
ctx.JSON(map[string]string{"hello": "json"}) // or myjsonStruct{hello:"json}
```
- Jsonp
```
ctx.JSONP(map[string]string{"hello": "jsonp"}, context.JSONP{Callback: "callbackName"})
```
- XML
```
ctx.XML(ExampleXML{One: "hello", Two: "xml"}) // or iris.Map{"One":"hello"...}
```
- Markdown
```
ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
```
### 事务
```
//子域也适用于所有可用的路由器，就像其他功能一样。

app.Get("/", func(ctx context.Context) {
	ctx.BeginTransaction(func(t *context.Transaction) {
		//选项步骤：如果为真，那么如果此事务失败，则不会执行下一个事务
		// t.SetScope(context.RequestTransactionScope)
		//可选步骤：
		//在这里创建一个新的自定义类型的错误以跟踪状态码和原因消息
		err := context.NewTransactionErrResult()

		//我们应该使用t.Context，如果我们想回滚这个函数clojure中的任何错误。
		t.Context().Text("Blablabla this should not be sent to the client because we will fill the err with a message and status")

		//在这里虚拟出一个假的错误，以举例说明
		fail := true
		if fail {
			err.StatusCode = iris.StatusInternalServerError
			//注意：如果是空的原因，那么默认或自定义http错误将被触发（如ctx.FireStatusCode）
			err.Reason = "Error: Virtual failure!!"
		}

		//选项步骤：
		//但如果我们想在事务失败时将错误消息回发给客户端，那么它很有用。
		//如果原因为空，则交易成功完成，
		//否则我们会回滚整个回复作者的主体，
		//标题和Cookie，状态码和所有内容都存在于此事务中
		t.Complete(err)
	})

	ctx.BeginTransaction(func(t *context.Transaction) {
		t.Context().HTML("<h1>This will sent at all cases because it lives on different transaction and it doesn't fails</h1>")
		// *如果我们没有任何'throw error'逻辑，那么不需要
	})

	// OPTIONALLY取决于用法：
	//在任何情况下，上下文事务中发生的事情都会发送给客户端
	ctx.HTML("<h1>Let's add a second html message to the response, " +
		"if the transaction was failed and it was request scoped then this message would " +
		"not been shown. But it has a transient scope(default) so, it is visible as expected!</h1>")
})
```
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
- html
```
<form enctype="multipart/form-data"
	action="http://127.0.0.1:8080/upload" method="POST">
	<input type="file" name="uploadfile" /> <input type="hidden"
		name="token" value="{{.}}" /> <input type="submit" value="upload" />
</form>
```
- token
```
// Serve the upload_form.html to the client.
app.Get("/upload", func(ctx iris.Context) {
	// create a token (optionally).
	now := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(now, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	// render the form with the token for any use you'd like.
	// ctx.ViewData("", token)
	// or add second argument to the `View` method.
	// Token will be passed as {{.}} in the template.
	ctx.View("upload_form.html", token)
})
```
- upload
```
app.Post("/upload", func(ctx iris.Context) {
	// iris.LimitRequestBodySize(32 <<20) as middleware to a route
	// or use ctx.SetMaxRequestBodySize(32 << 20)
	// to limit the whole request body size,
	//
	// or let the configuration option at app.Run for global setting
	// for POST/PUT methods, including uploads of course.

	// Get the file from the request.
	file, info, err := ctx.FormFile("uploadfile")

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile("./路径/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)
})
// start the server at http://localhost:8080 with post limit at 32 MB.
app.Run(iris.Addr(":8080"), iris.WithPostMaxMemory(32<<20))
```
- 修改名称
```
func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=	time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}
```
### 文件下载
```
app.Get("/", func(ctx iris.Context) {
		file := "[path]"
		ctx.SendFile(file, "name")
})
```
### 测试
- GDB调试
- 测试用例
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
- 压力测试（基准测试）
> 压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似
首字母不能是小写字母func BenchmarkXXX(b *testing.B) { ... }
需要带上参数 -test.bench语法: - test.bench="test_name_regex" ,例如 go test -test.bench=".*"
在压力测试用例中,请记得在循环体内使用 testing.B.N ,以使测试可以正常的运行
文件名也必须以 _test.go 结尾

```
func Benchmark_Division(b *testing.B) {
    //[..]
}
func Benchmark_TimeConsumingFunction(b *testing.B) {
    b.StopTimer() // 调用该函数停止压力测试的时间计数
    // 做一些初始化的工作 , 例如读取文件数据 , 数据库连接之类的 ,
    b.StartTimer() // 重新开始时间
    // [..]
}
```
- 样本测试
> 以Example开头  
函数体末尾添加注释比较输出内容是否与预期相符,第一行注释以Output:开头
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
### file-logger
```
// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func main() {
	f := newLogFile()
	defer f.Close()
	// attach the file as logger, remember, iris' app logger is just an io.Writer.
	app.Logger().SetOutput(newLogFile())
	app.Get("/", func(ctx iris.Context) {
		// for the sake of simplicity, in order see the logs at the ./_today_.txt
		ctx.Application().Logger().Info("Request path: " + ctx.Path())
	})
	if err := app.Run(iris.Addr(":8080"), iris.WithoutBanner); err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warn("Shutdown with error: " + err.Error())
		}
	}
}
```
### Redis
- iris自带的redis
`不过感觉不是很好用`
```
// replace with your running redis' server settings:
db := redis.New(service.Config{
	Network:     service.DefaultRedisNetwork,
	Addr:        service.DefaultRedisAddr,
	Password:    "",
	Database:    "",
	MaxIdle:     0,
	MaxActive:   0,
	IdleTimeout: service.DefaultRedisIdleTimeout,
	Prefix:      ""}) // optionally configure the bridge between your redis server

// close connection when control+C/cmd+C
iris.RegisterOnInterrupt(func() {
	db.Close()
})
sess := sessions.New(sessions.Config{Cookie: "sessionscookieid", Expires: 45 * time.Minute})
//
// IMPORTANT:
//
sess.UseDatabase(db)
//set
s.Set("key", value)
// get a specific key, as string, if no found returns just an empty string
name := sess.Start(ctx).GetString("key")
// delete a specific key
sess.Start(ctx).Delete("bbb")
// removes all entries
sess.Start(ctx).Clear()
//destroy, removes the entire session data and cookie
sess.Destroy(ctx)
//update
sess.ShiftExpiration(ctx)

```
- Redigo
`比IRIS自带的好用多了`  
    - 读写
```
c, err := redis.Dial("tcp", "127.0.0.1:6379")//, options)  
if err != nil {  
	fmt.Println(err)  
	return  
}  
defer c.Close()  
//执行命令使用的Do函数，和直接执行redis-cli命令差不多  
v, err := c.Do("SET", "test", "redisgo")  
if err != nil {  
	fmt.Println(err)  
	return  
}  
fmt.Println(v)  
v, err = redis.String(c.Do("GET", "test"))  
if err != nil {  
	fmt.Println(err)  
	return  
}  
fmt.Println(v)  
```  
    - 设置过期
```
_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
```  
    - 是否存在Key
```
is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey1"))
```  
    - 删除键
```
_, err = c.Do("DEL", "mykey")
```
    - 列表操作
```
    _, err = c.Do("lpush", "runoobkey", "redis")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    _, err = c.Do("lpush", "runoobkey", "mongodb")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }
    _, err = c.Do("lpush", "runoobkey", "mysql")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))

    for _, v := range values {
        fmt.Println(string(v.([]byte)))
    }
```  
    - 管道
```
c.Send("SET", "foo", "bar")
c.Send("GET", "foo")
c.Flush()
c.Receive() // reply from SET
v, err = c.Receive() // reply from GET
```
    - 连接池
```
连接池的结构
type Pool struct {
    //Dial 是创建链接的方法
    Dial func() (Conn, error)

    //TestOnBorrow 是一个测试链接可用性的方法
    TestOnBorrow func(c Conn, t time.Time) error

    // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
    MaxIdle int

    // 最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
    MaxActive int

    //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
    IdleTimeout time.Duration

    // 当链接数达到最大后是否阻塞，如果不的话，达到最大后返回错误
    Wait bool

}
```
```
使用连接池
//声明一些全局变量
var (
    pool          *redis.Pool
    redisServer   = flag.String("redisServer", ":6379", "")
    redisPassword = flag.String("redisPassword", "123456", "")
)
//初始化一个pool
func newPool(server, password string) *redis.Pool {
    return &redis.Pool{
        MaxIdle:     3,
        MaxActive:   5,
        IdleTimeout: 240 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            if _, err := c.Do("AUTH", password); err != nil {
                c.Close()
                return nil, err
            }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < time.Minute {
                return nil
            }
            _, err := c.Do("PING")
            return err
        },
    }
}

func main() {
    flag.Parse()
    pool = newPool(*redisServer, *redisPassword)
    conn := pool.Get()
    defer conn.Close()
}
```
### API Doc
```
import (
	"github.com/kataras/iris"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
)
```
```
yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
	On:       true,
	DocTitle: "Iris",
	DocPath:  "apidoc.html",
	BaseUrls: map[string]string{"Production": "", "Staging": ""},
})
app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.
```
### i18n
```
ini文件中为: hi = 您好，%s

func newApp() *iris.Application {
	app := iris.New()
	globalLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "./locales/locale_en-US.ini",
			"el-GR": "./locales/locale_el-GR.ini",
			"zh-CN": "./locales/locale_zh-CN.ini"}})
	app.Use(globalLocale)
	app.Get("/", func(ctx iris.Context) {
		// it tries to find the language by:
		// ctx.Values().GetString("language")
		// if that was empty then
		// it tries to find from the URLParameter setted on the configuration
		// if not found then
		// it tries to find the language by the "language" cookie
		// if didn't found then it it set to the Default setted on the configuration

		// hi is the key, 'iris' is the %s on the .ini file
		// the second parameter is optional

		// hi := ctx.Translate("hi", "iris")
		// or:
		hi := i18n.Translate(ctx, "hi", "iris")

		language := ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())
		// return is form of 'en-US'

		// The first succeed language found saved at the cookie with name ("language"),
		//  you can change that by changing the value of the:  iris.TranslateLanguageContextKey
		ctx.Writef("From the language %s translated output: %s", language, hi)
	})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

```
### PPROF
`pprof是golang程序一个性能分析的工具，可以查看堆栈、cpu信息等。`
```
func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	app.Any("/debug/pprof/{action:path}", pprof.New())
	//                              ___________
	app.Run(iris.Addr(":8080"))
}
```
### GO-BINDATA
```
该软件包可将任何文件转换为可管理的Go源代码。用于将二进制数据嵌入到go程序中。文件数据在转换为原始字节片段之前可以选择进行gzip压缩。

它在go-bindata子目录中提供了一个命令行工具。该工具提供了一组命令行选项，用于自定义正在生成的输出。

安装

要安装库和命令行程序，请使用以下命令：

go get -u github.com/shuLhan/go-bindata/...
用法

转换是在一组或多组文件上完成的。它们都嵌入到一个新的Go源文件中，以及一个目录和一个资产功能，该功能允许基于其名称快速访问资产。

最简单的调用在当前工作目录中生成一个bindata.go文件。它包含数据目录中的所有资产。

$ go-bindata data/
要递归地包含所有输入子目录，请使用为Go导入路径定义的elipsis后缀。否则，它只会考虑输入目录本身的资产。

$ go-bindata data/...
要指定正在生成的输出文件的名称，我们使用以下内容：

$ go-bindata -o myfile.go data/
如果需要，可以指定多个输入目录。

$ go-bindata dir1/... /path/to/dir2/... dir3
以下段落详细介绍了一些可以提供给go-bindata的命令行选项。有关testdata / in中资源的各种输出示例，请参阅testdata / out目录。每个示例使用不同的命令行选项。

要忽略文件，请使用-ignore传入正则表达式，例如：

$ go-bindata -ignore=\\.gitignore data/...
Accessing an asset

To access asset data, we use the Asset(string) ([]byte, error) function which is included in the generated output.

data, err := Asset("pub/style/foo.css")
if err != nil {
	// Asset was not found.
}

// use asset data
调试与发布构建

使用-debug标志调用程序时，生成的代码实际上不包含资产数据。相反，它会生成函数存根，用于从磁盘上的原始文件加载数据。资产API在调试版本和发布版本之间保持一致，因此您的代码不必更改。

当您期望资产经常更改时，这在开发过程中很有用。使用这些资产的主机应用程序在这两种情况下都使用相同的API，并且不必关心实际数据来自何处。

一个例子是一个带有一些嵌入式静态Web内容的Go web服务器，如HTML，JS和CSS文件。在开发它的时候，你不想重建整个服务器，并在你每次改变一些javascript时重启它。您只想构建并启动服务器一次。然后只需在浏览器中按刷新即可查看这些更改。用调试标志嵌入资产可以让你做到这一点。当您完成开发并准备部署时，只需重新调用go-bindata而不使用-debug标志。它现在将嵌入最新版本的资产。
```
