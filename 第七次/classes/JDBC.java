import java.io.IOException;
import java.io.PrintWriter;
import java.util.*;
import java.sql.*;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class JDBC extends HttpServlet {
    private static final long serialVersionUID = 1L;

    public JDBC() {
        super();
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        response.setContentType("text/html;charset=UTF-8");
	PrintWriter out = response.getWriter();
	int Chk = 0;
        String Sid = request.getParameter("id");
        String Slink = request.getParameter("link");
        String Spwd = request.getParameter("pwd");
        String Sroles = request.getParameter("roles");
        String Suid = request.getParameter("uid");
        String Susername = request.getParameter("username");
        try {
            Class.forName("com.mysql.jdbc.Driver");
            Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/mydatabase","root","Xsydx886.");
            Statement stat=conn.createStatement();
            String sql = "INSERT INTO students(id,link,pwd,roles,uid,username) VALUES('"+Sid+"','"+Slink+"','"+Spwd+"','"+Sroles+"','"+Suid+"','"+Susername+"')";
            out.println("SQL语句为:"+sql+"<br>");
            String sqls = "SELECT * FROM students WHERE id = " + Sid;
            ResultSet rs = stat.executeQuery(sqls);
            while(rs.next())
            {
              String rsid = rs.getString("id");
              if(rsid.equals(Sid))
              {
                out.println("违反主键规则");
                Chk =1;
              }
            }
            if(Chk ==0)
            {
              int i = stat.executeUpdate(sql);
              out.println("新建学生成功 成功添加"+i+"行");
            }
            stat.close();
            conn.close();
            response.setHeader("refresh","4;URL=http://hduoct.xyz/Servlet-JDBC.jsp");
        }
        catch(Exception e) {
            e.printStackTrace();
        }
    }

    public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        doGet(request, response);
    }
}
