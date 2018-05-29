package Article

import (
	"github.com/shiyanhui/hero"
	"io"
	"../../Entity"
	"strconv"
	"github.com/kataras/iris"
	"../../DAO"
)
func ContextWriter(entity Entity.Entity,pre Entity.Article,suc Entity.Article,ctx iris.Context, w io.Writer) (int, error){
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>index</title>
    <!-- Bootstrap -->
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <style type="text/css">
        body{ font-family: Microsoft YaHei,'宋体' , Tahoma, Helvetica, Arial, "\5b8b\4f53", sans-serif;}
    </style>
    <script src="https://cdn.bootcss.com/markdown.js/0.6.0-beta1/markdown.min.js"></script>
</head>
<body>
`)
	_buffer.WriteString(`<nav id="navbar-example" class="navbar navbar-default navbar-static navbar-fixed-top" role="navigation">
    <div class="container-fluid">
        <div class="navbar-header">
            <button class="navbar-toggle" type="button" data-toggle="collapse"
                    data-target=".bs-js-navbar-scrollspy">
                <span class="sr-only">切换导航</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/">Web开发学习笔记</a>
        </div>
        <div class="collapse navbar-collapse bs-js-navbar-scrollspy">
            <ul class="nav navbar-nav">`)
	for i,m := range entity.MenuList {
		_buffer.WriteString(`<li><a href="/menu/`+strconv.Itoa(i+1)+`">`+m.Name+`</a></li>`)
	}
	_buffer.WriteString(`
            </ul>
            <form class="navbar-form navbar-left">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Search" value="golang是最好的语言">
                </div>
                <button type="submit" class="btn btn-default">Submit</button>
            </form>
            <ul class="nav navbar-nav navbar-right">
				<li><a href="/adminlogin">后台</a></li>`)
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !auth {
		_buffer.WriteString(`<li><a href="/register">注册</a></li>`+`<li><a href="/login">登录</a></li>`)
	}else {
		_buffer.WriteString(`<li><a href="/user">欢迎你&nbsp;:&nbsp;`+Entity.Sess.Start(ctx).GetString("Username")+`&nbsp;&nbsp;&nbsp;更多</a></li>`)
	}
	_buffer.WriteString(`
            </ul>
        </div>
    </div>
</nav>
<div class="cow">
    <div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
    <div data-spy="scroll" data-target="#navbar-example" data-offset="0"class="col-md-8 col-lg-8 col-sm-10 col-xs-10">
        <div>
			<ol class="breadcrumb">
			  <li><a href="/">主页</a></li>
			  <li><a href="/menu/`)
	_buffer.WriteString(strconv.FormatInt(entity.Menu.Id,10))
	_buffer.WriteString(`/1">`)
	_buffer.WriteString(entity.Menu.Name)
	_buffer.WriteString(`</a></li>
			  <li class="active">`)
    _buffer.WriteString(entity.Article.Title)
    _buffer.WriteString(`</li>
			</ol>
			<hr/>
			<h5>分类:`)
    _buffer.WriteString(entity.Article.Classify)
	_buffer.WriteString(`<p class="text-right">上一篇: <a href="/article/`)
	_buffer.WriteString(strconv.FormatInt(entity.Article.Id-1,10)+`">`)
	_buffer.WriteString(pre.Title)
	_buffer.WriteString(`</a></p><p class="text-right">下一篇: <a href="/article/`)
	_buffer.WriteString(strconv.FormatInt(entity.Article.Id+1,10)+`">`)
	_buffer.WriteString(suc.Title)
	_buffer.WriteString(`</a></p></h5><br/>`)
	_buffer.WriteString("<h3>"+entity.Article.Title+"</h3><br/><h5>作者:")
	_buffer.WriteString(entity.UserData.Username)
	_buffer.WriteString("</h5><h5>最后修改日期:")
	_buffer.WriteString(entity.Article.Time.Format("2006-01-02 15:04:05"))
	_buffer.WriteString("</h5><br/>")
	_buffer.WriteString("<div>"+string(entity.Article.Content)+"</div>")
	_buffer.WriteString(`
        </div>
	<hr/>
	<h3>留言&nbsp;共&nbsp;`)
	_buffer.WriteString(strconv.Itoa(len(entity.CommentList)))
	_buffer.WriteString(`&nbsp;条)</h3>
	<hr/>`)
	user := Entity.UserData{}
	for _,com := range entity.CommentList {
		var userdatadao DAO.UserDataDAOInterface = new(DAO.UserData)
		_,_,user = userdatadao.Get(Entity.UserData{Id:com.User})
		_buffer.WriteString("<h4>"+strconv.FormatInt(com.Floor,10)+` 楼 `)
		_buffer.WriteString(user.Username)
		_buffer.WriteString(` 说:</h4>
		<p>`+com.Content+`</p>
		<p class="text-right">日期:`+com.Time.Format("2006-01-02 15:04:05")+` | </p><a href="#">回复</a>
	<hr dashed>`)
	}
	_buffer.WriteString(`<hr>
	<br><br>
	<h3>发表观点</h3>
	<form  method="post" name="form" id="form">
		<textarea class="form-control" rows="3" id="Comment" name="Comment"></textarea>
		<input type="text" id ="Article" name="Article" style="display:none" value="`+strconv.FormatInt(entity.Article.Id,10)+`"></input>
		<input type="text" id ="Floor" name="Floor" style="display:none" value="`+strconv.Itoa(len(entity.CommentList)+1)+`"></input>
	</form>
	<button class="btn btn-default" name="btn" id="btn">提交</button>
	<p id="text" name="text"></p>
    </div>
    <div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
</div>`)

	_buffer.WriteString(`
<script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
<script type="text/javascript" src="../js/CommentAjax.js"></script>
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}