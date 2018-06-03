需要安装Redis Golang Mysql ElasticSearch

I.在Docker下部署(推荐) 按步骤执行即可(不带序号的不用执行) 若不在Docker下部署(直接在机器内部署) 则 跳到第85行(不在Docker下部署)
------------Docker--------------
1 yum install docker
2 sudo systemctl start docker
3 sudo systemctl enable docker
4 sudo service docker start
5 vim /etc/docker/daemon.json
改为
{
  "registry-mirrors": ["http://hub-mirror.c.163.com"]
}
保存:wq
6 mkdir /usr/docker
7 setenforce 0
【重启机器后如何重启容器:】
setenforce 0
docker ps -a
第一栏就是ID
docker start ID
【删除容器】
docker rm ID
【删除镜像】
docker images
docker rmi ID

------------mysql----------------
8 mkdir /usr/docker/mysql
9 cd /usr/docker/mysql
10 docker run --name mymysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Xsydx886. -d mysql
11 上一步MYSQL_ROOT_PASSWORD后为Mysql密码 若修改 则须打开工程的main/UtilOrmUtil.go Getorm函数的NewEngine方法内root:后改为该密码
12 docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)
13 找到mymysql的ip地址 打开工程的main/UtilOrmUtil.go Getorm函数的NewEngine方法内ip改为该ip 端口不变
【进入容器】 
14 docker exec -it mymysql bash
15 mysql -u root -p
12 你的密码
16 create database Goweb
17 exit
18 exit
------------redis----------------
19 mkdir /usr/docker/redis
20 mkdir /usr/docker/redis/data
21 cd /usr/docker/redis
22 docker run --name myredis -p 6379:6379 -v $PWD/data:/data  -d redis redis-server --appendonly yes
23 docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)
24 找到myredis的ip地址 打开工程的 main/DAO/RedisDAO.go 找到redis.Dial("tcp", "172.17.0.2:6379")中ip改为你的myredis的ip
【连接redis】
25 docker run --rm -it --link myredis:redis-cli redis bash
26 redis-cli -h redis -p 6379
27 set h1 任意一个单词如 dasfa
28 set ac 任意一句话如fadsfasdfasdfsadfasdfaf
29 exit
30 exit

---------elasticsearch-------------
31 docker run --name myes -p 9200:9200 -p 9300:9300 -d elasticsearch

------kibana(es可视化管理工具 可不装)--------------
32 docker run --name mykibana --link=myes:elasticsearch -p 5601:5601 -d kibana

---------------golang--------------
33 进入工程的main目录
34 docker build -t goweb .
35 docker run -it --name mygoweb -p 80:80  --link=myredis:redis --link=mymysql:mysql -v /home/kannaduki/GOPATH/WorkSpace/GoWeb/main/:/go/src/goweb/ -w /go/src/goweb goweb 【-v 参数后：前的部分为工程的main路径 需要修改】
36 ctrl + c
37 docker exec -it mymysql bash
38 mysql -u root -p
39 你的密码
40 use Goweb
41 alter table article change content content text;
42 insert into admin_user(account,password) value('admin','admin');
43 exit
44 exit
45 docker ps -a
46 找到Goweb前的id 执行docker start ID
47 localhost:80
47 必须先注册一个用户 才能进入后台 进去后先创建几个栏目 然后再创建一篇文章 然后就可以正常使用了
48 关于聊天室 因为限制了一个设备只登陆一个帐号 不能一个浏览器开两个页面登两个号 为了方便测试 session劫持的部分被我注释了 可以用两个浏览器登两个号测试


II.不在Docker下部署(麻烦 不推荐)

------------redis----------------
解压并进入redis目录
mkdir /usr/local/redis
mv src/redis-server /usr/local/redis/
mv src/redis-cli /usr/local/redis/
vim redis.conf
daemonize 后的no改为yes 保存退出:wq
/usr/local/redis/redis-server redis.conf
/usr/local/redis/redis-cli
set h1 大标题名(随便写一个)
set ac 公告内容(随便写一个)
exit

---------- bash_profile(跳过 供下面修改时参考)----------
# .bash_profile

# Get the aliases and functions
if [ -f ~/.bashrc ]; then
        . ~/.bashrc
fi

# User specific environment and startup programs

#GOLANG
export GOROOT_BOOTSTRAP=/home/kannaduki/GOROOT_BOOTSTRAP/go
export GOROOT=/home/kannaduki/go
export GOPATH=/home/kannaduki/GOPATH

#JAVA
export JAVA_HOME=/usr/java/jkd-10.0.1
export JRE_HOME=/usr/java/jkd-10.0.1/jre
export CLASS_PATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JRE_HOME/lib

PATH=$PATH:$HOME/bin:/usr/lib/ccache/bin:$GOROOT/bin:$JAVA_HOME/bin:$JRE_HOME/bin

export PATH


------------------------Go--------------------------------
vim  ~/.bash_profile
最后添加
###############################################################
export GOROOT_BOOTSTRAP=/home/.../(解压后的GOROOT_BOOTSTRAP目录)
###############################################################
保存 :wq
source ~/.bash_profile
cd (解压后的go目录)/go/src/
./make.bash
vim  ~/.bash_profile
最后添加
######################################################################
#GO
export GOROOT=/home/.../go (解压的go目录)
export GOPATH=/home/.../GOPATH (自己新建的一个GOPATH文件夹 用来存放框架和工作目录)
PATH=$PATH:$HOME/bin:$GOROOT/bin (如果已经有PATH了 在后面添加 :$GOROOT/bin)
export PATH
#######################################################################
保存 :wq
source ~/.bash_profile
输入go回车出现提示 则安装成功
go get github.com/kataras/iris
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
go get github.com/gomodule/redigo/redis
go get github.com/shiyanhui/hero

---------------Jdk环境(elasticsearch 和 Kibana 需要用到)----------------------
mkdir /usr/java
解压到jdk到 /usr/java
vim  ~/.bash_profile
最后添加
########################################################################
#JAVA
export JAVA_HOME=/usr/java/jkd-10.0.1
export JRE_HOME=/usr/java/jkd-10.0.1/jre
export CLASS_PATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JRE_HOME/lib
修改PATH
PATH=$PATH:$HOME/bin:/usr/lib/ccache/bin:$GOROOT/bin:$JAVA_HOME/bin:$JRE_HOME/bin
export PATH
#########################################################################
保存 :wq
source ~/.bash_profile

-----------------------mysql--------------------------------
wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm
sudo rpm -ivh mysql-community-release-el7-5.noarch.rpm
sudo yum install mysql-server
sudo chown -R root:root /var/lib/mysql
service mysqld restart
mysql -u root
use mysql;
//重置密码
update user set password=password('123456') where user='root';
exit;
service mysqld restart
mysql -u root -p
你的密码
CREATE DATABASE Goweb
exit 
打开项目的/Util/OrmUtil.go Getorm函数的NewEngine方法的root后的文字改为你的密码

--------------------- elasticsearch ---------------------------------
解压后 cd bin 
./elasticsearch即可打开
Kibana同理


----------------------------------- 运行项目 -------------------------------------------------
cd main
go build
运行
nohup ./main &
端口为80 localhost:80
关闭
netstat -tunlp|grep 80
kill -9 进程号
因为ORM框架有点问题 别的表的生成都没有问题 就需要手动修改一个列属性
mysql -u root -p
你的密码
use Goweb
ALTER TABLE article CHANGE content content text;
insert into admin_user(account,password) value('admin','admin');
exit
然后就可以正常使用了 需要注册一个用户 进入后台后先添加几个文章和栏目
