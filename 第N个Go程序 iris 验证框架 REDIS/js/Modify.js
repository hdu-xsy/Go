function mvalidate(v) {
    var selected = 0;
    var a = document.getElementsByName("select");
    var s = $('input:radio[name="select"]:checked').val();
    var form = document.getElementById("form");
    //$('input[name="username'+s+'"]').attr("name","username");
    //$('input[name="password'+s+'"]').attr("name","password");
    //$('input[name="username"]').attr("id","username");
    //$('input[name="password"]').attr("id","password");
    var x = $("input[name='username"+s+"']").val();
    var y = $("input[name='password"+s+"']").val();
    var z = $("input[name='userid"+s+"']").val();
    $("#Id").val(z);
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
        //return true;
        if(v == 0) {
            document.form.action = "/delete";
            document.forms.from.submit();
        }else if(v == 1) {
            document.form.action = "/modify";
            document.forms.from.submit();
        }
    }
    else {
        return false;
    }
}
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