<%@ page language="java" contentType="text/html; charset=UTF-8" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>Test1</title>
</head>
<body>
  <%
      Integer hitsCount =  (Integer)application.getAttribute("hitCounter");
      if( hitsCount ==null || hitsCount == 0 ){
         hitsCount = 1;
      }else{
         hitsCount += 1;
      }
      application.setAttribute("hitCounter", hitsCount);
  %>
  <p>访问量: <%= hitsCount%></p>
</body>
</html>
