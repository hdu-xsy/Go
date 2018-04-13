$(function() {
    $("#btn").click(function() {
        var params = {
            Comment : $("#Comment").val(),
            Article : $("#Article").val(),
            Floor : $("#Floor").val(),
        };
        $.post("/Comment",params,function (data) {
            if (data == " ") {
                window.location.reload();
            } else {
                $("#text").html(data)
            }
        });
    });
});