package Controller

import (
	"github.com/kataras/iris/websocket"
	"github.com/kataras/iris/mvc"
	"../Util"
	"../Entity"
	"fmt"
)

func ConfigureMVC(m *mvc.Application) {
	ws := websocket.New(websocket.Config{})
	m.Router.Any("/iris-ws", websocket.ClientHandler())

	// This will bind the result of ws.Upgrade which is a websocket.Connection
	// to the controller(s) served by the `m.Handle`.
	m.Register(ws.Upgrade)
	m.Handle(new(WebSocketController))
}
type WebSocketController struct {
	Conn websocket.Connection
}

func (c *WebSocketController) On(msg string) {
	c.Conn.To(websocket.All).Emit("chat", msg)
	fmt.Println(msg)
	var mesg  = showonlineuser()
	c.Conn.To(websocket.All).Emit("login",mesg)
}

func (c *WebSocketController) OnDisconnect() {
	orm := Util.Getorm(*app)
	olu := Entity.OnlineUser{Uid:uid}
	orm.Delete(&olu)
	var mesg = showonlineuser()
	c.Conn.To(websocket.All).Emit("login",mesg)
}
func (c *WebSocketController) Get( /* websocket.Connection could be lived here as well, it doesn't matter */ ) {
	c.Conn.On("chat", c.On)
	c.Conn.OnDisconnect(c.OnDisconnect)
	c.Conn.Wait()
}

func showonlineuser() string{
	orm := Util.Getorm(*app)
	var olu []Entity.OnlineUser
	orm.Find(&olu)
	var mesg string
	for _,k := range olu {
		mesg = mesg + k.Username + "   " + k.Logintime.Format("2006-01-02 15:04:05")+"<br>"
	}
	return mesg
}