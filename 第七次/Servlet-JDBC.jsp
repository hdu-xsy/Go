<%@ page language="java" contentType="text/html; charset=GBK" pageEncoding="GBK"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>

<html>
<head>
<title>JDBC</title>
</head>
<body>
  <form action="JDBC">
    id<lable><input type="text" id ="id" name="id"></lable><br>
    link<lable><input type="text" id ="link" name="link"></lable><br>
    pwd<lable><input type="text" id ="pwd" name="pwd"></lable><br>
    roles<lable><input type="text" id ="roles" name="roles"></lable><br>
    uid<lable><input type="text" id ="uid" name="uid"></lable><br>
    username<lable><input type="text" id ="username" name="username"></lable><br>
    <label><input type="submit" name="submit" value="添加"></label>
  	<label><input type="reset" name="reset"></label>
  </form>
     <%

          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/mydatabase","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql;
          sql = "SELECT * FROM students";
          ResultSet rs = stat.executeQuery(sql);
          out.println("从数据库查询出的数据为"+"<br>"+"<br>");
          %>
          <table border="1">
            <tr>
                <td>id</td>
                <td>link</td>
                <td>pwd</td>
                <td>roles</td>
                <td>uid</td>
                <td>username</td>
            </tr>
          <%while(rs.next()){
            String rsid = rs.getString("id");
            String rslink = rs.getString("link");
            String rspwd = rs.getString("pwd");
            String rsroles = rs.getString("roles");
            String rsuid = rs.getString("uid");
            String rsusername = rs.getString("username");
            %>
            <tr>
                <td><%=rsid %></td>
                <td><%=rslink%></td>
                <td><%=rspwd%></td>
                <td><%=rsroles%></td>
                <td><%=rsuid%></td>
                <td><%=rsusername%></td>
            </tr>
          <% }
          stat.close();
          conn.close();
          %>


</table>
</body>
</html>
