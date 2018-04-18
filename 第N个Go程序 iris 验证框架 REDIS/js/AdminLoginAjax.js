$(function() {
    $("#btn").click(function() {
        //提交的参数，name是struts action中对应的接收变量
        var params = {
            Account : $("#Account").val(),
            Password : $("#Password").val()
        };
        $.post("/AdminLoginAjax",params,function (data) {
            if (data == " ") {
                $(location).attr('href', '/backend');
            } else {
                $("#text").html(data)
            }
        });
    });
});