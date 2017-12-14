<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>
<html>
<head>
<title>无标题文档</title>
<style type="text/css">
<!--
.STYLE1 {font-size: 12px}
.STYLE3 {color: #707070; font-size: 12px; }
.STYLE5 {color: #0a6e0c; font-size: 12px; }
body {
	margin-top: 0px;
	margin-bottom: 0px;
}
.STYLE7 {font-size: 12}
-->
</style>
</head>

<body>
<table width="100%" border="0" cellspacing="0" cellpadding="0">
  <tr>
    <td height="30"><table width="100%" border="0" cellspacing="0" cellpadding="0">
      <tr>
        <td>&nbsp;</td>
        <td style="padding-right:10px;"><div align="right">
          <table border="0" align="right" cellpadding="0" cellspacing="0">
            <tr>
              <td width="60"><table width="87%" border="0" cellpadding="0" cellspacing="0">
                  <tr>
                    <td class="STYLE1"><div align="center">
                        <input type="checkbox" name="checkbox62" value="checkbox" />
                    </div></td>
                    <td class="STYLE1"><div align="center">全选</div></td>
                  </tr>
              </table></td>
              <td width="60"><table width="90%" border="0" cellpadding="0" cellspacing="0">
                  <tr>
                    <td class="STYLE1"><div align="center"><img src="images/001.gif" width="14" height="14" /></div></td>
                    <td class="STYLE1"><div align="center"><a href="insert.jsp">新增</a></div></td>
                  </tr>
              </table></td>
              <td width="60"><table width="90%" border="0" cellpadding="0" cellspacing="0">
                  <tr>
                    <td class="STYLE1"><div align="center"><img src="images/114.gif" width="14" height="14" /></div></td>
                    <td class="STYLE1"><div align="center"><a href="motify.jsp">修改</a></div></td>
                  </tr>
              </table></td>
              <td width="52"><table width="88%" border="0" cellpadding="0" cellspacing="0">
                  <tr>
                    <td class="STYLE1"><div align="center"><img src="images/083.gif" width="14" height="14" /></div></td>
                    <td class="STYLE1"><div align="center"><a href="delete.jsp">删除</a></div></td>
                  </tr>
              </table></td>
            </tr>
          </table>
        </div></td>
      </tr>
    </table></td>
  </tr>

  <%
      int count = 0;
       Class.forName("com.mysql.jdbc.Driver");
       Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
       Statement stat=conn.createStatement();
       String sql;
       sql = "SELECT * FROM log order by opertime desc";
       ResultSet rs = stat.executeQuery(sql);
   %>

  <tr>
    <td><table width="100%" border="0" cellpadding="0" cellspacing="1" bgcolor="#c9c9c9">
      <tr>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><strong><span class="STYLE1">操作用户</span></strong></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><strong><span class="STYLE1">操作时间</span></strong></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><strong><span class="STYLE1">操作名称</span></strong></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><strong><span class="STYLE1">操作对象</span></strong></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><strong><span class="STYLE1">明细</span></strong></div></td>
      </tr>

      <%
      while(rs.next() && count <9){
      String rsoperuser = rs.getString("operuser");
      String rsoperdate = rs.getString("opertime");
      String rsopername = rs.getString("opername");
      String rsoperobject = rs.getString("operobject");
      count ++;
       %>
      <tr>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><span class="STYLE3"><%=rsoperuser %></span></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><span class="STYLE3"><%=rsoperdate %></span></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><span class="STYLE3"><%=rsopername %></span></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center"><span class="STYLE3"><%=rsoperobject %></span></div></td>
        <td height="22" bgcolor="#FFFFFF"><div align="center" class="STYLE5">明细</div></td>
      </tr>
      <%
      }
      stat.close();
      conn.close();
      %>
    </table></td>
  </tr>
  <tr>
    <td height="35"><table width="100%" border="0" cellspacing="0" cellpadding="0">
      <tr>
        <td width="25%" height="29" nowrap="nowrap"><table width="342" border="0" cellspacing="0" cellpadding="0">
          <tr>
            <td width="44%" class="STYLE1">当前页：1/2页 每页
              <input name="textfield2" type="text" class="STYLE1" style="height:14px; width:25px;" value="15" size="5" />            </td>
            <td width="14%" class="STYLE1"><img src="images/sz.gif" width="43" height="20" /></td>
            <td width="42%" class="STYLE1"><span class="STYLE7">数据总量 15 </span></td>
          </tr>
        </table></td>
        <td width="75%" valign="top" class="STYLE1"><div align="right">
            <table width="352" height="20" border="0" cellpadding="0" cellspacing="0">
              <tr>
                <td width="62" height="22" valign="middle"><div align="right"><img src="images/page_first_1.gif" width="48" height="20" /></div></td>
                <td width="50" height="22" valign="middle"><div align="right"><img src="images/page_back_1.gif" width="55" height="20" /></div></td>
                <td width="54" height="22" valign="middle"><div align="right"><img src="images/page_next.gif" width="58" height="20" /></div></td>
                <td width="49" height="22" valign="middle"><div align="right"><img src="images/page_last.gif" width="52" height="20" /></div></td>
                <td width="59" height="22" valign="middle"><div align="right">转到第</div></td>
                <td width="25" height="22" valign="middle"><span class="STYLE7">
                  <input name="textfield" type="text" class="STYLE1" style="height:10px; width:25px;" size="5" />
                </span></td>
                <td width="23" height="22" valign="middle">页</td>
                <td width="30" height="22" valign="middle"><img src="images/go.gif" width="26" height="20" /></td>
              </tr>
            </table>
        </div></td>
      </tr>
    </table></td>
  </tr>
</table>
</body>
</html>
