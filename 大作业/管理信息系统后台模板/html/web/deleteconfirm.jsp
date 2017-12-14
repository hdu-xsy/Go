<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>
<html>
<head>
<title>无标题文档</title>
</head>
<body>
    <%
      String sel = request.getParameter("delete");
      int selint = Integer.valueOf(sel).intValue();
      int count = 1;
      Class.forName("com.mysql.jdbc.Driver");
      Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
      Statement stat=conn.createStatement();
      String sql = "SELECT * FROM log order by opertime desc;";
      ResultSet rs = stat.executeQuery(sql);
      while(rs.next())
      {
        if(count == selint)
        {
            String rsoperuser = rs.getString("operuser");
            String rsopertime = rs.getString("opertime");
            String rsopername = rs.getString("opername");
            String rsoperobject = rs.getString("operobject");
            String deletesql = "DELETE FROM log WHERE operuser = '"+rsoperuser+"' AND opertime = '"+rsopertime+"' AND opername = '"+rsopername+"' AND operobject = '"+rsoperobject+"';";
            int i = stat.executeUpdate(deletesql);
            break;
        }
        count ++;
      }
      stat.close();
      conn.close();
    %>
    <script language="javascript" type="text/javascript">
    window.location.href='right.jsp';
    setTimeout("javascript:location.href='right.jsp'", 0);
    </script>

</body>
</html>
