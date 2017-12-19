<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>无标题文档</title>
<style type="text/css">
<!--
body {
	margin-left: 0px;
	margin-top: 0px;
	margin-right: 0px;
	margin-bottom: 0px;
}
.STYLE1 {font-size: 12px}
.STYLE4 {
	font-size: 12px;
	color: #1F4A65;
	font-weight: bold;
}

a:link {
	font-size: 12px;
	color: #06482a;
	text-decoration: none;

}
a:visited {
	font-size: 12px;
	color: #06482a;
	text-decoration: none;
}
a:hover {
	font-size: 12px;
	color: #FF0000;
	text-decoration: underline;
}
a:active {
	font-size: 12px;
	color: #FF0000;
	text-decoration: none;
}

-->
</style>
<script type="text/javascript">
   function validate() {
      if(document.form1.checkbox.value == "")
      {
        alert("请选择数据");
        return false;
      }
			if(document.form1.username.value == "")
			{
				alert("请输入用户名");
				return false;
			}
			if(document.form1.pwd.value == "")
			{
				alert("请输入密码");
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
<table width="100%" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td height="30"><table width="100%" border="0" cellspacing="0" cellpadding="0">
      <tr>
        <td width="15" height="30"><img src="images/tab_03.gif" width="15" height="30" /></td>
        <td background="images/tab_05.gif"><img src="images/311.gif" width="16" height="16" /> <span class="STYLE4">服务器进程配置列表</span></td>
        <td width="14"><img src="images/tab_07.gif" width="14" height="30" /></td>
      </tr>
    </table></td>
  </tr>
  <tr>
    <td><table width="100%" border="0" cellspacing="0" cellpadding="0">
      <tr>
        <td width="9" background="images/tab_12.gif">&nbsp;</td>
        <td bgcolor="e5f1d6"><table width="99%" border="0" align="center" cellpadding="0" cellspacing="1" bgcolor="#CECECE">
          <tr>
            <td width="6%" height="26" background="images/tab_14.gif" class="STYLE1"><div align="center" class="STYLE2 STYLE1">选择</div></td>
            <td width="8%" height="18" background="images/tab_14.gif" class="STYLE1"><div align="center" class="STYLE2 STYLE1">编号</div></td>
            <td width="24%" height="18" background="images/tab_14.gif" class="STYLE1"><div align="center" class="STYLE2 STYLE1">用户名</div></td>
            <td width="10%" height="18" background="images/tab_14.gif" class="STYLE1"><div align="center" class="STYLE2 STYLE1">密码</div></td>
            <td width="7%" height="18" background="images/tab_14.gif" class="STYLE1"><div align="center" class="STYLE2">编辑</div></td>
            <td width="7%" height="18" background="images/tab_14.gif" class="STYLE1"><div align="center" class="STYLE2">删除</div></td>
          </tr>
					<%
							 int count = 1;
               String sel = request.getParameter("id");
               int selint = Integer.valueOf(sel).intValue();
							 Class.forName("com.mysql.jdbc.Driver");
							 Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javawebpaper","root","Xsydx886.");
							 Statement stat=conn.createStatement();
							 String sql;
							 sql = "SELECT * FROM users order by id";
							 ResultSet rs = stat.executeQuery(sql);
					 %>
					 <%
					 String rsid="";
					 String rsusername="";
					 String rspwd="";
           while(rs.next())
           {
             if(count == selint)
             {
                 rsid = rs.getString("Id");
                 rsusername = rs.getString("username");
                 rspwd = rs.getString("pwd");
                 break;
             }
             count ++;
           }
           stat.close();
           conn.close();
						%>
            <form action="../../../../paper_motify_confirm" name="form1" id="form1">
						<tr>
	            <td height="18" bgcolor="#FFFFFF"><div align="center" class="STYLE1">
	              <input name="checkbox" type="checkbox" class="STYLE2" value="<%=rsid %>" checked="checked" />
	            </div></td>
	            <td height="18" bgcolor="#FFFFFF" class="STYLE2"><div align="center" class="STYLE2 STYLE1"><%=rsid %></div></td>
	            <td height="18" bgcolor="#FFFFFF"><div align="center" class="STYLE2 STYLE1"><input type="text" id ="username" name="username" VALUE="<%=rsusername %>"></div></td>
	            <td height="18" bgcolor="#FFFFFF"><div align="center" class="STYLE2 STYLE1"><input type="text" id ="pwd" name="pwd" VALUE="<%=rspwd %>"></div></td>
	            <td height="18" bgcolor="#FFFFFF"><div align="center"><img src="images/037.gif" width="9" height="9" /><span class="STYLE1"> [</span>编辑<span class="STYLE1">]</span></div></td>
	            <td height="18" bgcolor="#FFFFFF"><div align="center"><span class="STYLE2"><img src="images/010.gif" width="9" height="9" /> </span><span class="STYLE1">[</span>删除<span class="STYLE1">]</span></div></td>
	          </tr>
						<input type="submit" name="submit" value="修改" onclick="return validate()">
						<input type="reset" name="reset">
						</form>
        </table></td>
        <td width="9" background="images/tab_16.gif">&nbsp;</td>
      </tr>
    </table></td>
  </tr>
  <tr>
    <td height="29"><table width="100%" border="0" cellspacing="0" cellpadding="0">
      <tr>
        <td width="15" height="29"><img src="images/tab_20.gif" width="15" height="29" /></td>
        <td background="images/tab_21.gif"><table width="100%" border="0" cellspacing="0" cellpadding="0">
          <tr>
            <td width="40%"><div align="left"><span class="STYLE1">共120条纪录，当前第1/10页，每页10条纪录</span></div></td>
            <td width="60%" class="STYLE1">&nbsp;</td>
          </tr>
        </table></td>
        <td width="14"><img src="images/tab_22.gif" width="14" height="29" /></td>
      </tr>
    </table></td>
  </tr>
</table>
</body>
</html>
