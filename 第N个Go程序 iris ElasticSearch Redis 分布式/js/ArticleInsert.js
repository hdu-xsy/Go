function validate(){
    var a = document.getElementById("Title").value;
    var b = document.getElementById("Content").value;
    if(a == ""){
        alert("标题不能为空");
        return false;
    }
    if(b == ""){
        alert("内容不能为空");
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
            Title : $("#Title").val(),
            Menu : $("#Menu").val(),
            Content : $("#Content").val()
        };
        $.post("/artinsert",params,function (data) {
            if (data == " ") {
                $(location).attr('href', '/backend/1');
            } else {
                $("#text").html(data)
            }
        });
    });
});