<%@ page language="java" contentType="text/html; charset=gb2312" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>用户注册</title>
</head>
<body>
  <p>各种肉大甩卖一律8元</p>
  <form action="session4.jsp">
  <input type="checkbox" name="obj" value="猪肉">猪肉<br>
  <input type="checkbox" name="obj" value="牛肉">牛肉<br>
  <input type="checkbox" name="obj" value="羊肉">羊肉<br>
  <input type="submit" name="submit" value="确认">
  <input type="reset" name="reset">
  </form>
  <a href="session.jsp">买点别的</a>
  <a href="shopcar.jsp">查看购物车</a>
</body>
</html>
