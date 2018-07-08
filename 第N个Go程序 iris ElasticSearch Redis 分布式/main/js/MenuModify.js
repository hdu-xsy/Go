function mvalidate(v) {
    var selected = 0;
    var a = document.getElementsByName("select");
    var s = $('input:radio[name="select"]:checked').val();
    var form = document.getElementById("form");
    //$('input[name="username'+s+'"]').attr("name","username");
    //$('input[name="password'+s+'"]').attr("name","password");
    //$('input[name="username"]').attr("id","username");
    //$('input[name="password"]').attr("id","password");
    var x = $("input[name='name"+s+"']").val();
    var z = $("input[name='id"+s+"']").val();
    $("#Id").val(z);
    $("#Name").val(x);
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
            document.form.action = "/menudelete";
            document.forms.from.submit();
        }else if(v == 1) {
            document.form.action = "/menumodify";
            document.forms.from.submit();
        }
    }
    else {
        return false;
    }
}