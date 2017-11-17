<%@ page language="java" contentType="text/html; charset=gb2312" pageEncoding="gb2312"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>

<html>
<head>
<title>JDBC</title>
</head>
<body>
     <%
          int Chk = 0;
          String Sid = request.getParameter("id");
          String Slink = request.getParameter("link");
          String Spwd = request.getParameter("pwd");
          String Sroles = request.getParameter("roles");
          String Suid = request.getParameter("uid");
          String Susername = request.getParameter("username");
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/mydatabase","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql = "INSERT INTO students(id,link,pwd,roles,uid,username) VALUES('"+Sid+"','"+Slink+"','"+Spwd+"','"+Sroles+"','"+Suid+"','"+Susername+"')";
          out.println("SQL语句为:"+sql+"<br>");
          String sqls = "SELECT * FROM students WHERE id = " + Sid;
          ResultSet rs = stat.executeQuery(sqls);
          while(rs.next())
          {
            String rsid = rs.getString("id");
            if(rsid.equals(Sid))
            {
              out.println("违反主键规则");
              Chk =1;
            }
          }
          if(Chk ==0)
          {
            int i = stat.executeUpdate(sql);
            out.println("新建学生成功 成功添加"+i+"行");
          }
          stat.close();
          conn.close();
      %>
      <script type="text/javascript">
        window.setTimeout("window.location='list.jsp'",5000);
      </script>

</body>
</html>
