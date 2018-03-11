$(function() {
    $("#btn").click(function() {
        var params = {
            Username : $("#Username").val(),
            Password : $("#Password").val()
        };
        $.post("/UserLoginAjax",params,function (data) {
            if (data == " ") $(location).attr('href', '/chatform');
            $("#text").html(data)});
    });
});