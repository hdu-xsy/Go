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
<td>�û�����<%=name_post%></td>
</tr>
<tr>
<td>���룺<%=password1_post%></td>
</tr>
<tr>
<td>���䣺<%=mail_post%></td>
</tr>
<tr>
<td>���᣺<%=from_post%></td>
</tr>
<tr>
<td>���գ�<%=birth_post%></td>
</tr>
<tr>
<td>��ߣ�<%=tall_post%></td>
</tr>
<tr>
<td>�ֻ���<%=phone_post%></td>
</tr>
<tr>
<td>���ã�<%=love_post%></td>
</tr>
<tr>
<td>�Ա�<%=sex_post%></td>
</tr>
<tr>
<td>���ܣ�<%=info_post%></td>
</tr>
</table>
</body>
</html>
