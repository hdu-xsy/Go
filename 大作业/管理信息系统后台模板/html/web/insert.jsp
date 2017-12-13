<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*"%>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>无标题文档</title>
<style type="text/css">
<!--
.STYLE1 {font-size: 12px}
.STYLE3 {color: #707070; font-size: 12px; }
.STYLE5 {color: #0a6e0c; font-size: 12px; }
body {
	margin-top: 0px;
	margin-bottom: 0px;
}
.STYLE7 {font-size: 12}
-->
</style>
<script type="text/javascript">
   function validate() {
      var a = document.getElementById("user").value;
      var b = document.getElementById("time").value;
      var c = document.getElementById("name").value;
      var d = document.getElementById("object").value;
      if(a == "")
      {
        alert("操作用户 不能为空");
        return false;
      }
      if(b == "")
      {
        alert("操作时间 不能为空");
        return false;
      }
      if(c == "")
      {
        alert("操作名称 不能为空");
        return false;
      }
      if(d == "")
      {
        alert("操作对象 不能为空");
        return false;
      }
      if (confirm("提交表单?"))
      {
        return true;
      }
      else {
        return false;
      }
   }
	</script>

</head>
<body>
<p>服务端时间为:</p>
<%
Date date = new Date();
out.print( "<p>" +date.toString()+"<p>");
%>
      <form action="insertconfirm.jsp">
        操作用户<lable><input type="text" id ="user" name="user"></lable><br>
        操作时间<lable><input type="text" id ="time" name="time"></lable><br>
        操作名称<lable><input type="text" id ="name" name="name"></lable><br>
        操作对象<lable><input type="text" id ="object" name="object"></lable><br>
        <label><input type="submit" name="submit" value="确认" onclick="return validate()"></label>
      	<label><input type="reset" name="reset"></label>
      </form>
</body>
</html>
