<%@ page language="java" import="java.util.*" pageEncoding="gbk"%>  

<html>  
    <head>  
        <title></title>  
    </head>  
  
    <body>  
        <%  
response.setHeader("Refresh","5");
ArrayList showList=(ArrayList)(getServletContext().getAttribute("list"));  
out.print("在线人数 "+showList.size()+"<br>");  
for(int i=0;i<showList.size();i++){  
out.print(showList.get(i)+"在线"+"<br>");  
}  
%>  
        <br>  
        <a href="index.jsp">退出</a>  
    </body>  
</html>  