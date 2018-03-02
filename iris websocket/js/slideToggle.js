$(document).ready(function() {
    $("#s").click(function(){
        $("#p").slideToggle("slow");
        if($("#s").text() == "隐藏在线用户")
            $("#s").text("显示在线用户");
        else $("#s").text("隐藏在线用户");
    });
});