function mvalidate() {
    var selected = 0;
    var a = document.getElementsByName("select");
    var s = $('input:radio[name="select"]:checked').val();
    //$('input[name="username'+s+'"]').attr("name","username");
    //$('input[name="password'+s+'"]').attr("name","password");
    //$('input[name="username"]').attr("id","username");
    //$('input[name="password"]').attr("id","password");
    var x = $("input[name='username"+s+"']").val();
    var y = $("input[name='password"+s+"']").val();
    $("#Username").val(x);
    $("#Password").val(y);
    for(var i=0;i<a.length;i++){
        if(a.item(i).checked==true){
            selected = 1;
            break;
        }
    }
    if(selected == 0){
        alert("请选择记录");
        return false;
    }
    if (confirm("提交?")){
        return true;
    }
    else {
        return false;
    }
}
