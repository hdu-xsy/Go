<%@ page language="java" contentType="text/html; charset=UTF-8" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>Test1</title>
</head>
<body>
  <script language="javascript" type="text/javascript">
  function time()
  {
    var t = new Date(<%=new Date().getTime()%>);
    var day = t.getDay();
    var hours =t.getHours();
    var minutes =t.getMinutes();
    if(hours<=12)
    {
      document.getElementById("Systime").innerHTML="今天是星期"+day+"上午"+hours+"点"+minutes+"分";
    }
    else
    {
      document.getElementById("Systime").innerHTML="今天是星期"+day+"下午"+hours+"点"+minutes+"分";
    }
  }
  </script>

  <p id="Systime"></p>
  <body onload="time()">


</body>
</html>
