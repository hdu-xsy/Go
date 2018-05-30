FROM golang
RUN go get github.com/kataras/iris
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/go-xorm/xorm
RUN go get github.com/gomodule/redigo/redis
RUN go get github.com/shiyanhui/hero
EXPOSE 80
CMD ["go","run","main.go"]

