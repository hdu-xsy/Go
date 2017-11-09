<%@ page language="java" contentType="text/html; charset=gb2312" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>用户注册</title>
</head>
<body>
  <p>您选择的结果是</p>
<%
  ArrayList objout = (ArrayList)session.getAttribute("objs");
    if(objout !=null )
  for(int i=0;i<objout.size();i++)
  {
    String object = (String)objout.get(i);
    out.println(object + "<br>");
  }
  ArrayList objout2 = (ArrayList)session.getAttribute("objs2");
  if(objout2 != null)
    for(int i=0;i<objout2.size();i++)
  {
    String object2 = (String)objout2.get(i);
    out.println(object2 + "<br>");
  }
%>
</body>
</html>
