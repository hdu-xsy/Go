package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"github.com/go-xorm/xorm"
	"strconv"
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
		// Write message back to the client message owner with: c.Emit("chat", msg)
		// Write message to all except this client with:
		c.To(websocket.Broadcast).Emit("chat", msg)
		var mesg  = showonlineuser()
		c.To(websocket.All).Emit("login",mesg)
	})
	c.OnDisconnect(func() {
		orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
		if err != nil {
			app.Logger().Fatalf("orm failed to initialized: %v", err)
		}
		iris.RegisterOnInterrupt(func(){
			orm.Close()
		})
		if err = orm.Sync2(new(OnlineUser));err !=nil {
			app.Logger().Fatalf("orm failed to initialized User table: %v", err)
		}
		olu := OnlineUser{Id:uuu}
		orm.Delete(&olu)
		var mesg = showonlineuser()
		c.To(websocket.All).Emit("login",mesg)
	})
}

func showonlineuser() string{
	orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	if err = orm.Sync2(new(OnlineUser));err !=nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	var olu []OnlineUser
	orm.Find(&olu)
	var mesg string
	for i,k := range olu {
		mesg = mesg + strconv.Itoa(i) + " " + k.Username + " " + k.Logintime.String()+"<br>"
	}
	return mesg
}