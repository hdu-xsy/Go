<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*,java.text.*"%>
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
<script type="text/javascript">
   function validate() {
      if(document.form1.motify.value == "")
      {
        alert("请选择数据");
        return false;
      }
      if (confirm("提交表单?"))
      {
        return true;
      }
      else {
        return false;
      }
   }
	</script>
</head>
<body>
<%
		int count = 0;
		 Class.forName("com.mysql.jdbc.Driver");
		 Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
		 Statement stat=conn.createStatement();
		 String sql;
		 sql = "SELECT * FROM log order by opertime desc";
		 ResultSet rs = stat.executeQuery(sql);
 %>
<form action="../../motifyconfirm" name="form1" id="form1">
<tr>
	<td><table width="100%" border="0" cellpadding="0" cellspacing="1" bgcolor="#c9c9c9">
		<tr>
		  <td height="22" bgcolor="#FFFFFF"><div align="center"><strong><span class="STYLE1">修改</span></strong></div></td>
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
		  <td height="22" bgcolor="#FFFFFF"><div align="center"><span class="STYLE3"><input type="radio" name="motify" value="<%= count %>"></span></div></td>
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
    java.util.Date date = new java.util.Date();
    out.print( "<p>" +date.toString()+"<p>");
    DateFormat format1 = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
		%>
	</table></td>
</tr>
操作用户<lable><input type="text" id ="user" name="user"></lable><br>
操作时间<lable><input type="text" id ="time" name="time" VALUE="<%= format1.format(date) %>"></lable><br>
操作名称<lable><input type="text" id ="name" name="name"></lable><br>
操作对象<lable><input type="text" id ="object" name="object"></lable><br>
<input type="submit" name="submit" value="修改" onclick="return validate()">
<input type="reset" name="reset">
</form>
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
</body>
</html>
