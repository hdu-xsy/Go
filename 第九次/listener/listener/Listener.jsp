<%@ page language="java" import="java.util.*" pageEncoding="gbk"%>  

<html>  
    <head>  
        <title></title>  
    </head>  
    <body>  
  
        <%  
            session = request.getSession(false);  
            if (session != null)  
                session.invalidate();  
        %>  
        <form action="isOnline.jsp" method="post">  
            用户名：  
            <input type="text" name="uName" />  
            <input type="submit" value="上线">  
    </body>  
</html>  