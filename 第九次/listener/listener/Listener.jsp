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
            �û�����  
            <input type="text" name="uName" />  
            <input type="submit" value="����">  
    </body>  
</html>  