$(function() {
    $("#btn").click(function() {
        //提交的参数，name是struts action中对应的接收变量
        var params = {
            Account : $("#Account").val(),
            Password : $("#Password").val()
        };
        $.post("/AdminLoginAjax",params,function (data) {
            if (data == "TRUE") $(location).attr('href', '/backend');
            $("#text").html(data)});
    });
});