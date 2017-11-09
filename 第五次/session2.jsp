<%@ page language="java" contentType="text/html; charset=gb2312" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title></title>
</head>
<body>
<%
  int i;
  request.setCharacterEncoding("gb2312");
  String[] b = request.getParameterValues("obj");
  ArrayList objs = new ArrayList();
    for(i = 0;i<b.length;i++)
    {
      objs.add(b[i]);
    }
    session.setAttribute("objs",objs);
%>
<script type="text/javascript">
  window.setTimeout("window.location='session.jsp'",1000);
  </script>
</body>
</html>
