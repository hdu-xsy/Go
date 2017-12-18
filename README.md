# JavaWeb大作业题目

------


### 1.制作官方主页的首页，也就是静态页面中的index.htm，其他不需要你制作。
> * - [x] 把静态页面index.htm改成index.jsp。
> * - [x] 静态页面中公司简介div中，图片不变，文字介绍需要从数据库中读取该栏目的栏目描述。
> * - [x] 静态页面中公司新闻div中，图片不变，文字介绍需要从数据库中读取该栏目的栏目描述。然后显示4条（实际可能多于4条）对应此栏目的新闻记录（显示标题即可）
> * - [x] 静态页面中的“行业新闻”div，图片不变，标题改成“技术服务”，然后从数据库读取该栏目的文字介绍以及显示4条相应的新闻记录。
> * - [x] 联系我们的内容，保持不变，不需要做。
> * - [x] 静态页面中的“最新产品“div改成”产品介绍“div，然后从数据库中读取相应的新闻记录，图片是对应的图片路径，产品名称对应的是新闻标题。

---
### 2.制作网站后台，关于网站后台的静态页面在admin文件夹下。
> * - [ ] 首先通过admin/login/login.html，当用户点击登录以后，读取数据库的users表，如果用户名密码正确，则跳到admin/web/main.html，否则返回login.html。
> * - [ ] 修改left中的栏目，修改成“用户管理“、”栏目管理“、“新闻管理”，三个功能点。
> * - [ ] 点击用户管理的时候，中间的表格从数据库中读取信息，显示用户记录，然后完成“编辑“跟“删除”功能。

注意：也可以采用admin1文件夹的静态页面模板，功能如上，此模板中登录界面是index.html，主功能界面是main.html。

---
