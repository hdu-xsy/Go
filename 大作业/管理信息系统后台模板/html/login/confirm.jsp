<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>

<html>
<head>
<title>confirm</title>
</head>
<body>
     <%
          String username = request.getParameter("textfield");
          String password = request.getParameter("textfield2");
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql = "SELECT * FROM users WHERE username = '" + username + "';";
          out.println(sql);
          ResultSet rs = stat.executeQuery(sql);
          while(rs.next())
          {
            String rsusername = rs.getString("username");
            String rspwd = rs.getString("pwd");
            if(rsusername.equals(username) && rspwd.equals(password))
            {
                response.sendRedirect("../web/main.html");
            }
          }
          stat.close();
          conn.close();
      %>
      <script language="javascript" type="text/javascript">
      window.location.href='login.html';
      setTimeout("javascript:location.href='login.html'", 0);
      </script>

</body>
</html>
