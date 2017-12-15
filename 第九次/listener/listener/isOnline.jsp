<%@ page language="java" import="java.util.*" pageEncoding="gbk"%>  

<html>  
  <head>  
    <title></title>  
  </head>  
    
  <body>  
    <%  
  
session=request.getSession();  
  
session.setAttribute("userName",request.getParameter("uName"));  
  
response.sendRedirect("showOnline.jsp");  
%>  
  </body>  
</html>  