
// Code generated by hero.
// source: C:\Users\hduoct\Desktop\GoWeb\Article\body.html
// DO NOT EDIT!
package Article

import (
	"github.com/shiyanhui/hero"
	"io"
	"../Entity"
	"strconv"
)
func ContextWriter(Content Entity.Article,prearticle Entity.Article,username string,menu Entity.Menu, w io.Writer) (int, error){
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
	_buffer.WriteString(`<nav id="navbar-example" class="navbar navbar-default navbar-static" role="navigation">
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
                <li><a href="/menu/9">关于二次元</a></li>
            </ul>
            <form class="navbar-form navbar-left">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Search" value="golang是最好的语言">
                </div>
                <button type="submit" class="btn btn-default">Submit</button>
            </form>
            <ul class="nav navbar-nav navbar-right">
				<li><a href="/adminlogin">后台</a></li>
                <li><a href="/register">注册</a></li>
                <li><a href="/login">登录</a></li>
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
	_buffer.WriteString(strconv.FormatInt(menu.Id,10))
	_buffer.WriteString(`">`)
	_buffer.WriteString(menu.Name)
	_buffer.WriteString(`</a></li>
			  <li class="active">`)
    _buffer.WriteString(Content.Title)
    _buffer.WriteString(`</li>
			</ol>
			<hr/>
			<h4>分类:`)
    _buffer.WriteString(Content.Classify)
	_buffer.WriteString(`<p class="text-right">上一篇: <a href="/article/`)
	_buffer.WriteString(strconv.FormatInt(Content.Id-1,10))
	_buffer.WriteString(`">`)
	_buffer.WriteString(prearticle.Title)
	_buffer.WriteString(`</a></p></h4><br/>`)
	_buffer.WriteString("<h3>"+Content.Title+"</h3><br/><h5>作者:")
	_buffer.WriteString(username)
	_buffer.WriteString("</h5><h5>最后修改日期:")
	_buffer.WriteString(Content.Time.Format("2006-01-02 15:04:05"))
	_buffer.WriteString("</h5><br/>")
	_buffer.WriteString("<div>"+string(Content.Content)+"</div>")
	_buffer.WriteString(`
        </div>
	<hr/>
	<h3>留言(共 施工中 条):施工中</h3>
    </div>
    <div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
</div>`)

	_buffer.WriteString(`
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}