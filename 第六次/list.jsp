<%@ page language="java" contentType="text/html; charset=GBK" pageEncoding="GBK"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>

<html>
<head>
<title>JDBC</title>
</head>
<body>
  <form action="">
    id<lable><input type="text" id ="id" name="id"></lable><br>
    link<lable><input type="text" id ="link" name="link"></lable><br>
    pwd<lable><input type="text" id ="pwd" name="pwd"></lable><br>
    roles<lable><input type="text" id ="roles" name="roles"></lable><br>
    uid<lable><input type="text" id ="uid" name="uid"></lable><br>
    username<lable><input type="text" id ="username" name="username"></lable><br>
    <label><input type="submit" name="submit" value="查询"></label>
  	<label><input type="reset" name="reset"></label>
  </form>
     <%
          String Sid = request.getParameter("id");
          String Slink = request.getParameter("link");
          String Spwd = request.getParameter("pwd");
          String Sroles = request.getParameter("roles");
          String Suid = request.getParameter("uid");
          String Susername = request.getParameter("username");
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/mydatabase","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql;
          int Chk =0;
          if(Sid == null || Sid == "")
          {
            sql = "SELECT * FROM students";
          }

          else
          sql = "SELECT * FROM students WHERE ";
          if(Sid != null && Sid != ""){
            if(Chk == 1) sql+=" AND "; sql += "id = '"+Sid +"'";Chk =1;
          }
          if(Slink != null && Slink != "") {
            if(Chk == 1) sql+=" AND "; sql += "link = '"+Slink +"'";Chk =1;
          }
          if(Spwd != null && Spwd != "") {
            if(Chk == 1) sql+=" AND "; sql += "pwd = '"+Spwd +"'";Chk =1;
          }
          if(Sroles != null && Sroles != "") {
            if(Chk == 1) sql+=" AND "; sql += "roles = '"+Sroles +"'";Chk =1;
          }
          if(Suid != null && Suid != "") {
            if(Chk == 1) sql+=" AND "; sql += "uid = '"+Suid +"'";Chk =1;
          }
          if(Susername != null && Susername != "") {
            if(Chk == 1) sql+=" AND "; sql += "username = '"+Susername +"'";Chk =1;
          }
          out.println("sql代码为"+"<br>"+sql+"<br>");
          ResultSet rs = stat.executeQuery(sql);
          out.println("从数据库查询出的数据为"+"<br>");
          out.println("次序为id ，link ，pwd ，roles ，uid ，username"+"<br>");
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

      <form action="new.jsp">
        <label><input type="submit" name="submit" value="添加学生"></label>
      </form>

</table>
</body>
</html>
