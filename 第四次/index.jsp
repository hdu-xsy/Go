<%@ page language="java" contentType="text/html; charset=gb2312" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>�û�ע��</title>
<script type="text/javascript">

  function fname(){
    var a = document.getElementById("name").value;
    var b = document.getElementById("M4");
    var c = document.getElementById("M5");
    if(a == ""){
      b.innerHTML = "����Ϊ��";
    }
    if(a != ""){
      if (!('A'<=a.substring(0,1) && a.substring(0,1) <='z')) {
        c.innerHTML="����Ӣ�Ŀ�ͷ";
      }
      else {
        c.innerHTML="";
      }
      if(a.length>10 || a.length<6){
        b.innerHTML = "����Ϊ6-10λ";
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
			b.innerHTML = "����Ϊ��";
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
			b.innerHTML = "����Ϊ��";
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
			b.innerHTML = "����Ϊ��";
		}
		if(a != ""){
			b.innerHTML = "";
			if(a != document.getElementById("password1").value){
				b.innerHTML = "�������벻һ��";
			}
		}
	}

  function fphone(){
    var a = document.getElementById("phone").value;
    var b = document.getElementById("M6");
    if(a == ""){
      b.innerHTML = "����Ϊ��";
    }
    if(a != ""){
      if(a.length != 11){
        b.innerHTML = "����Ϊ11λ";
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
    if (confirm("�ύ��?")) {
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
      <label>�û���:<input type="text" id="name" name = "username" onClick="fname()" onKeyUp="fname()"><span id="M4" name="M4"></span><span id="M5" name="M5"></span></label>
      <br>
  		<label>����:<input name="password1" type="password" id="password1" onClick="fpassword1()" onKeyUp="fpassword1()"><span id="M2" name="M2"></span></label>
  		<br>
  		<label>ȷ������:<input type="password" id="password2" onClick="fpassword2()" onKeyUp="fpassword2()"><span id="M3" name="M3"></span></label>
  		<br>
      <label>����:<input name="mail" type="text" id="mail" onClick="email()" onKeyUp="email()"><span id="M1" name="M1"></span></label>
      <br>
  		<label>����:
        <select name="from">
          <option value="BeiJing">����</option>
          <option value="ShangHai">�Ϻ�</option>
          <option value="GuangZhou">����</option>
        </select>
      </label>
  		<br>
      <label>��������:<input name="birth" type="text"></label>
      <br>
      <label>���:<input name="tall" type="text"></label>
      <br>
      <label>�ֻ�:<input name="phone" id="phone" onclick="fphone()" onkeyup="fphone()" type="text"><span id="M6" name="M6"></span></label>
      <br>
  		<label>����:<input name="love" type="text"></label>
  		<br>
  		<label>�Ա�:<input name="sex" type="radio" name="sex" value="male">��
  		<input type="radio" name="sex" value="female">Ů</label>
  		<br>
  		���˽���<textarea name="info" rows="10" cols="30" name="self"></textarea>
  		<br>
  		<label><input type="submit" name="submit" value="ȷ��" onclick="return validate()"></label>
  		<label><input type="reset" name="reset"></label>
  	</form>

</body>
</html>
