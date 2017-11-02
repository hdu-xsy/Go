<%@ page language="java" contentType="text/html; charset=gb2312" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>用户注册</title>
<script type="text/javascript">

  function fname(){
    var a = document.getElementById("name").value;
    var b = document.getElementById("M4");
    var c = document.getElementById("M5");
    if(a == ""){
      b.innerHTML = "不能为空";
    }
    if(a != ""){
      if (!('A'<=a.substring(0,1) && a.substring(0,1) <='z')) {
        c.innerHTML="必须英文开头";
      }
      else {
        c.innerHTML="";
      }
      if(a.length>10 || a.length<6){
        b.innerHTML = "长度为6-10位";
      }
      else{
        b.innerHTML = "";
      }
    }
  }

	function email(){
	  var a = document.getElementById("mail").value;
		var b = document.getElementById("M1");
		if(a == ""){
			b.innerHTML = "不能为空";
      return false;
		}
    else{
      b.innerHTML ="";
      return true;
    }
	}

	function fpassword1(){
		var a = document.getElementById("password1").value;
		var b = document.getElementById("M2");
		if(a == ""){
			b.innerHTML = "不能为空";
      return false;
		}
		if(a != ""){
			b.innerHTML = "";
      return true;
	}
}

	function fpassword2(){
		var a = document.getElementById("password2").value;
		var b = document.getElementById("M3");
		if(a == ""){
			b.innerHTML = "不能为空";
		}
		if(a != ""){
			b.innerHTML = "";
			if(a != document.getElementById("password1").value){
				b.innerHTML = "两次密码不一致";
			}
		}
	}

  function fphone(){
    var a = document.getElementById("phone").value;
    var b = document.getElementById("M6");
    if(a == ""){
      b.innerHTML = "不能为空";
    }
    if(a != ""){
      if(a.length != 11){
        b.innerHTML = "长度为11位";
      }
      else{
        b.innerHTML = "";
        return true;
      }
    }
  }
   function validate() {
      var a = document.getElementById("M1");
      var b = document.getElementById("M2");
      var c = document.getElementById("M3");
      var d = document.getElementById("M4");
      var e = document.getElementById("M5");
      var f = document.getElementById("M6");
    if (confirm("提交表单?")) {
      if(fphone())
      {
        return true;
      }
      else {
        return false;
      }
    }
    else {
      return false;
    }
   }
	</script>
</head>
<body>
  <form action="post.jsp">
      <label>用户名:<input type="text" id="name" name = "username" onClick="fname()" onKeyUp="fname()"><span id="M4" name="M4"></span><span id="M5" name="M5"></span></label>
      <br>
  		<label>密码:<input name="password1" type="password" id="password1" onClick="fpassword1()" onKeyUp="fpassword1()"><span id="M2" name="M2"></span></label>
  		<br>
  		<label>确认密码:<input type="password" id="password2" onClick="fpassword2()" onKeyUp="fpassword2()"><span id="M3" name="M3"></span></label>
  		<br>
      <label>邮箱:<input name="mail" type="text" id="mail" onClick="email()" onKeyUp="email()"><span id="M1" name="M1"></span></label>
      <br>
  		<label>籍贯:
        <select name="from">
          <option value="BeiJing">北京</option>
          <option value="ShangHai">上海</option>
          <option value="GuangZhou">广州</option>
        </select>
      </label>
  		<br>
      <label>出生年月:<input name="birth" type="text"></label>
      <br>
      <label>身高:<input name="tall" type="text"></label>
      <br>
      <label>手机:<input name="phone" id="phone" onclick="fphone()" onkeyup="fphone()" type="text"><span id="M6" name="M6"></span></label>
      <br>
  		<label>爱好:<input name="love" type="text"></label>
  		<br>
  		<label>性别:<input name="sex" type="radio" name="sex" value="male">男
  		<input type="radio" name="sex" value="female">女</label>
  		<br>
  		个人介绍<textarea name="info" rows="10" cols="30" name="self"></textarea>
  		<br>
  		<label><input type="submit" name="submit" value="确认" onclick="return validate()"></label>
  		<label><input type="reset" name="reset"></label>
  	</form>

</body>
</html>
