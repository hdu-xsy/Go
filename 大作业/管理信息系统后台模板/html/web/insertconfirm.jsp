<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>

<html>
<head>
<title>insertconfirm</title>
</head>
<body>
     <%
          String user = request.getParameter("user");
          String time = request.getParameter("time");
          String name = request.getParameter("name");
          String object = request.getParameter("object");
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql = "INSERT INTO log(operuser,opertime,opername,operobject) VALUES('"+user+"','"+time+"','"+name+"','"+object+"')";
          int i = stat.executeUpdate(sql);
          stat.close();
          conn.close();
      %>
      <script language="javascript" type="text/javascript">
      window.location.href='right.jsp';
      setTimeout("javascript:location.href='right.jsp'", 0);
      </script>

</body>
</html>
