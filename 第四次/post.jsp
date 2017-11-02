<%@ page language="java" contentType="text/html; charset=gb2312" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>post.jsp</title>

</head>
<body>

  <%
      String name_post = request.getParameter("username");
      String password1_post = request.getParameter("password1");
      String mail_post = request.getParameter("mail");
      String from_post = request.getParameter("from");
      String birth_post = request.getParameter("birth");
      String tall_post = request.getParameter("tall");
      String phone_post = request.getParameter("phone");
      String love_post = request.getParameter("love");
      String sex_post = request.getParameter("sex");
      String info_post = request.getParameter("info");
  %>
  <table border="1px">
<tr>
<td>用户名：<%=name_post%></td>
</tr>
<tr>
<td>密码：<%=password1_post%></td>
</tr>
<tr>
<td>邮箱：<%=mail_post%></td>
</tr>
<tr>
<td>籍贯：<%=from_post%></td>
</tr>
<tr>
<td>生日：<%=birth_post%></td>
</tr>
<tr>
<td>身高：<%=tall_post%></td>
</tr>
<tr>
<td>手机：<%=phone_post%></td>
</tr>
<tr>
<td>爱好：<%=love_post%></td>
</tr>
<tr>
<td>性别：<%=sex_post%></td>
</tr>
<tr>
<td>介绍：<%=info_post%></td>
</tr>
</table>
</body>
</html>
