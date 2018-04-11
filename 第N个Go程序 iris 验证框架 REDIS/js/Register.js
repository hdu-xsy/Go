function validate(){
    var a = document.getElementById("Username").value;
    var b = document.getElementById("Password").value;
    if(a == ""){
        alert("用户名不能为空");
        return false;
    }
    if(b == ""){
        alert("用户名不能为空");
        return false;
    }
    if (confirm("提交?")){
        return true;
    }else{
        return false;
    }
}
$(function() {
    $("#btn").click(function() {
        //提交的参数，name是struts action中对应的接收变量
        var params = {
            Username : $("#Username").val(),
            Password : $("#Password").val()
        };
        $.post("/Register",params,function (data) {
            if (data == " ") {
                $(location).attr('href', '/login');
            } else {
                $("#text").html(data)
            }
        });
    });
});
function rs() {
    document.getElementById("form").reset();
}