var scheme = document.location.protocol == "https:" ? "wss" : "ws";
var port = document.location.port ? (":" + document.location.port) : "";
// see app.Get("/echo", ws.Handler()) on main.go
var wsURL = scheme + "://" + document.location.hostname + port+"/echo";
var input = document.getElementById("msg");
var output = document.getElementById("message");
var useroutput = document.getElementById("login")
// Ws comes from the auto-served '/iris-ws.js'
var socket = new Ws(wsURL)
socket.OnConnect(function () {
    output.innerHTML += "已连接服务器<br>";
    socket.Emit("chat",{{.Username}}+"连接服务器")
});
socket.OnDisconnect(function () {
    output.innerHTML += "无法连接服务器\n";
});
socket.OnConnect(function () {
    socket.Emit("login","aaaaaaaa")
})
socket.On("login",function (msg) {
    useroutput.innerHTML = msg
})
socket.onclose(function () {
    socket.Emit("login","bbbbbbbbbbb")
})
// read events from the server
socket.On("chat", function (msg) {
    addMessage(msg+"<br/>");
});
function send() {
    addMessage("Me: " + input.value); // write ourselves
    socket.Emit("chat", {{.Username}}+input.value);// send chat event data to the websocket server
    input.value = ""; // clear the input
}
function addMessage(msg) {
    output.innerHTML += msg + "<br>";
}