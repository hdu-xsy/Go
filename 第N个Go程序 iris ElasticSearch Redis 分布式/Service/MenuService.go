package Service

import (
	"github.com/kataras/iris"
	"../Entity"
	"strconv"
)


//栏目页面
type MenuService struct {

}
type Nav struct {
	Class	string
	Page	int
	Num		int
	Id		int64
}
func (s *MenuService)Get(ctx iris.Context) {
	h1 := redisdao.Get("h1")
	var auth string
	if userauth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !userauth { auth = "false" } else { auth = "true" }
	username := Entity.Sess.Start(ctx).GetString("Username")
	ctx.ViewData("username",username)

	menulist := menudao.GetAll()

	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,menu := menudao.Get(Entity.Menu{Id:id})

	articlelist := articledao.FindAll(strconv.FormatInt(id,10))

	m := articledao.Count()
	var suc,max int
	page64,_ := strconv.ParseInt(ctx.Params().Get("page"),10,64)
	page := int(page64)
	var c1,c2 string
	if page == 1 {
		c1 = "disabled"
	}else if page == max {
		c2 = "disabled"
	}
	if len(articlelist)%20 == 0 {
		max = len(articlelist)/20
	} else {
		max = len(articlelist)/20 + 1
	}
	if page>max || page == 0 {
		ctx.Redirect("/404")
		return
	}
	if page * 20 >= len(articlelist) {
		suc = len(articlelist)
	} else {
		suc = page * 20
	}
	for i,a := range articlelist {
		id,_ := strconv.ParseInt(a.Menu,10,64)
		_,_,menu := menudao.Get(Entity.Menu{Id:id})
		articlelist[i].Menu = menu.Name
	}

	nav := make([]Nav,5)
	for k := range nav {
		if max >5 {
			nav[k].Page = page + k
		} else {
			nav[k].Page = k + 1
		}
		flag := 0
		if (page == k+1 || page > 5) && flag == 0 {
			nav[k].Class = "active"
			flag = 1
		}
		if k+1 > max {
			nav[k].Class = "disabled"
		}
		nav[k].Id = id
		nav[k].Num = k+1
	}

	ctx.ViewData("h1",h1)
	ctx.ViewData("menuList",menulist)
	ctx.ViewData("auth",auth)
	ctx.ViewData("username",username)
	ctx.ViewData("menuName",menu.Name)
	ctx.ViewData("articleSum",len(articlelist))
	ctx.ViewData("articleList",articlelist[(page-1)*20:suc])
	ctx.ViewData("page",page)
	ctx.ViewData("id",id)
	ctx.ViewData("c1",c1)
	ctx.ViewData("c2",c2)
	ctx.ViewData("pagedec",page-1)
	ctx.ViewData("pageplus",page+1)
	ctx.ViewData("nav",nav)
	ctx.ViewData("classify",m)
	ctx.View("menu.html")
}


//分类页面
type ClassifySercice struct {

}
func (s *ClassifySercice) BeginRequest(ctx iris.Context) {
	c := ctx.Params().Get("Classify")
	var articleList []Entity.Article
	articleList = articledao.GetClassify()
	for _,v := range articleList {
		if v.Classify == c {
			return
		}
	}
	ctx.Redirect("/404")
}
func (s *ClassifySercice) EndRequest(ctx iris.Context) {}
func (s *ClassifySercice) Get(ctx iris.Context) {
	h1 := redisdao.Get("h1")
	var auth string
	if userauth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !userauth { auth = "false" } else { auth = "true" }
	username := Entity.Sess.Start(ctx).GetString("Username")

	menulist := menudao.GetAll()

	classify := ctx.Params().Get("Classify")
	var articleList []Entity.Article
	articleList = articledao.FindByClassify(classify)

	m := articledao.Count()
	var suc,max int
	page,_ := strconv.Atoi(ctx.Params().Get("page"))
	var c1,c2 string
	if page == 1 {
		c1 = "disabled"
	}else if page == max {
		c2 = "disabled"
	}
	if len(articleList)%20 == 0 {
		max = len(articleList)/20
	} else {
		max = len(articleList)/20 + 1
	}
	if page>max || page == 0 {
		ctx.Redirect("/404")
		return
	}
	if page * 20 >= len(articleList) {
		suc = len(articleList)
	} else {
		suc = page * 20
	}
	for i,a := range articleList {
		id,_ := strconv.ParseInt(a.Menu,10,64)
		_,_,menu := menudao.Get(Entity.Menu{Id:id})
		articleList[i].Menu = menu.Name
	}

	nav := make([]Nav,5)
	for k := range nav {
		if max >5 {
			nav[k].Page = page + k
		} else {
			nav[k].Page = k + 1
		}
		flag := 0
		if (page == k+1 || page > 5) && flag == 0 {
			nav[k].Class = "active"
			flag = 1
		}
		if k+1 > max {
			nav[k].Class = "disabled"
		}
		nav[k].Num = k+1
	}

	ctx.ViewData("h1",h1)
	ctx.ViewData("menuList",menulist)
	ctx.ViewData("auth",auth)
	ctx.ViewData("username",username)
	ctx.ViewData("classify",classify)
	ctx.ViewData("articleSum",len(articleList))
	ctx.ViewData("articleList",articleList[(page-1)*20:suc])
	ctx.ViewData("nav",nav)
	ctx.ViewData("c1",c1)
	ctx.ViewData("pagedec",page-1)
	ctx.ViewData("c2",c2)
	ctx.ViewData("pageplus",page+1)
	ctx.ViewData("classifyList",m)
	ctx.View("classify.html")
}