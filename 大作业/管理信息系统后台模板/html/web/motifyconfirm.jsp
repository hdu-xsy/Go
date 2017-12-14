<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>
<html>
<head>
<title>无标题文档</title>
</head>
<body>
    <%
      String sel = request.getParameter("motify");
      String user = request.getParameter("user");
      String time = request.getParameter("time");
      String name = request.getParameter("name");
      String object = request.getParameter("object");
      int selint = Integer.valueOf(sel).intValue();
      int count = 1;
      Class.forName("com.mysql.jdbc.Driver");
      Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
      Statement stat=conn.createStatement();
      String sql = "SELECT * FROM log order by opertime desc;";
      ResultSet rs = stat.executeQuery(sql);
      String rsoperuser = "";
      String wheresql = "";
      while(rs.next())
      {
        if(count == selint)
        {
            rsoperuser = rs.getString("operuser");
            String rsopertime = rs.getString("opertime");
            String rsopername = rs.getString("opername");
            String rsoperobject = rs.getString("operobject");
            wheresql = "WHERE operuser = '"+rsoperuser+"' AND opertime = '"+rsopertime+"' AND opername = '"+rsopername+"' AND operobject = '"+rsoperobject+"';";
            break;
        }
        count ++;
      }
      String updatesql = "update log ";
      String setsql = "set ";
      if(user != "")
      {
        setsql = setsql + "operuser = '"+user+"' ";
      }
      else
      {
        setsql = setsql + "operuser = '"+rsoperuser+"' ";
      }
      if(time != "")
      {
        setsql = setsql + ",opertime = '"+time+"' ";
      }
      if(name != "")
      {
        setsql = setsql + ",opername = '"+name+"' ";
      }
      if(object != "")
      {
        setsql = setsql + ",operobject = '"+object+"' ";
      }
      int i = stat.executeUpdate(updatesql+setsql+wheresql);
      stat.close();
      conn.close();
    %>
    <script language="javascript" type="text/javascript">
    setTimeout("javascript:location.href='right.jsp'", 0);
    </script>

</body>
</html>
