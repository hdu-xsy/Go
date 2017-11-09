<%@ page language="java" contentType="text/html; charset=gb2312" %>
<%@ page import="java.io .*,java.util.*" %>
<html>
<head>
<title>用户注册</title>
</head>
<body>
  <p>各大球大甩卖一律8元</p>
  <form action="session2.jsp">
  <input type="checkbox" name="obj" value="basketball">篮球<br>
  <input type="checkbox" name="obj" value="football">足球<br>
  <input type="checkbox" name="obj" value="volleyball">排球<br>
  <input type="submit" name="submit" value="确认">
  <input type="reset" name="reset">
  </form>
  <a href="session3.jsp">买点别的</a>
  <a href="shopcar.jsp">查看购物车</a>
</body>
</html>
