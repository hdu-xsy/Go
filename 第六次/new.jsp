<%@ page language="java" contentType="text/html; charset=GBK" pageEncoding="GBK"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>

<html>
<head>
<title>JDBC</title>
<script type="text/javascript">
   function validate() {
      var a = document.getElementById("id").value;
      if(a == "")
      {
        alert("id 不能为空");
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
      <form action="newConfirm.jsp">
        id<lable><input type="text" id ="id" name="id"></lable><br>
        link<lable><input type="text" id ="link" name="link"></lable><br>
        pwd<lable><input type="text" id ="pwd" name="pwd"></lable><br>
        roles<lable><input type="text" id ="roles" name="roles"></lable><br>
        uid<lable><input type="text" id ="uid" name="uid"></lable><br>
        username<lable><input type="text" id ="username" name="username"></lable><br>
        <label><input type="submit" name="submit" value="确认" onclick="return validate()"></label>
      	<label><input type="reset" name="reset"></label>
      </form>
</body>
</html>
