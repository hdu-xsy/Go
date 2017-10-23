<%@ page language="java" contentType="text/html; charset=UTF-8" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>Test2</title>
</head>
<body>
  <p>服务端时间为:</p>
  <%
     Date date = new Date();
     out.print( "<p>" +date.toString()+"<p>");
  %>

  <p>客户端时间为:</p>
  <script language="javascript" type="text/javascript">
  function time()
  {
    var t= new Date();
    var year=t.getFullYear();
    var month=t.getMonth();
    var dates=t.getDate();
    var hours =t.getHours();
    var minutes =t.getMinutes();
    var seconds =t.getSeconds();
    document.getElementById("Systime").innerHTML=month+"."+dates+" "+hours+":"+minutes+":"+seconds+" "+"CST"+" "+year;
  }
  </script>

  <p id="Systime"></p>
  <body onload="time()">
</body>
</html>
