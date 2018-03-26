
// Code generated by hero.
// source: C:\Users\hduoct\Desktop\GoWeb\index\body.html
// DO NOT EDIT!
package index
import (
	"github.com/shiyanhui/hero"
	"io"
	"../Entity"
	"strconv"
	"github.com/kataras/iris"
)
func ListWriter(entity Entity.Entity, ctx iris.Context,w io.Writer) (int, error){
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
	_buffer.WriteString(`
<nav id="navbar-example" class="navbar navbar-default navbar-static" role="navigation">
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
            <ul class="nav navbar-nav">
                <li><a href="/menu/1">网站更新日志</a></li>
                <li><a href="/menu/2">GO语言学习</a></li>
                <li><a href="/menu/3">Python语言学习</a></li>
                <li><a href="/menu/4">Java学习</a></li>
                <li><a href="/menu/5">JS学习</a></li>
                <li><a href="/menu/6">PHP学习</a></li>
                <li><a href="/menu/7">日记/感想</a></li>
                <li><a href="/menu/8">后端知识学习</a></li>
                <li><a href="/menu/9">计算机基础</a></li>
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
<div class="row">
    <div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
    <div class="col-md-5 col-lg-5 col-sm-10 col-xs-10">
        <div>
            <h3>最近的20篇文章</h3>
            <hr/>`)
            	_buffer.WriteString(`<br/><div class="list-group">`)
	for _, menu := range entity.ArticleList {
		_buffer.WriteString(`<a href="/article/`)
	hero.EscapeHTML(strconv.FormatInt(menu.Id,10),_buffer)
	_buffer.WriteString(`" class="list-group-item">
`)
	hero.EscapeHTML(menu.Time.Format("2006-01-02 15:04:05")+" : "+menu.Title, _buffer)
	_buffer.WriteString(`
</a>`)
}
_buffer.WriteString(`
            </div>
        </div>
    </div>
    <div class="col-md-3 col-lg-3 hidden-sm hidden-xs">
		<div><h3>最近留言</h3>
		<div class="list-group">`)
		for _, com := range entity.CommentList {
			_buffer.WriteString(`<a href="/article/`)
			hero.EscapeHTML(strconv.FormatInt(com.Article,10),_buffer)
			_buffer.WriteString(`" class="list-group-item">
		`)
			var con string
			if len(com.Content)>30 {con = com.Content[0:30]+"..."}else{con = com.Content}
			hero.EscapeHTML(com.Time.Format("2006-01-02 15:04:05")+" : "+con, _buffer)
			_buffer.WriteString(`
		</a>`)
		}
		_buffer.WriteString(`
				</div>
		</div><hr/>
        <div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">公告</h3>
			</div>
			<div class="panel-body">
				欢迎帮忙找BUG
			</div>
		</div>
    </div>
    <div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
</div>
`)

	_buffer.WriteString(`
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}