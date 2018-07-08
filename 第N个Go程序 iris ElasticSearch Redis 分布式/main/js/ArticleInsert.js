function validate(){
    var a = document.getElementById("Title").value;
    var b = document.getElementById("Content").value;
    var c = document.getElementById("Classify").value;
    if(a == ""){
        alert("标题不能为空");
        return false;
    }
    if(b == ""){
        alert("内容不能为空");
        return false;
    }
    if(c == ""){
        alert("分类不能为空");
        return false;
    }
    if (confirm("提交?")){
        return true;
    }else{
        return false;
    }
}

