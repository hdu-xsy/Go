import java.io.IOException;
import java.io.PrintWriter;
import java.util.*;
import java.sql.*;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class paper_login_confirm extends HttpServlet {
    private static final long serialVersionUID = 1L;

    public paper_login_confirm() {
        super();
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        response.setContentType("text/html;charset=UTF-8");
	      PrintWriter out = response.getWriter();
        String username = request.getParameter("textfield");
        String password = request.getParameter("textfield2");
        try {
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javawebpaper","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql = "SELECT * FROM users WHERE username = '" + username + "';";
          ResultSet rs = stat.executeQuery(sql);
          while(rs.next())
          {
            String rsusername = rs.getString("username");
            String rspwd = rs.getString("pwd");
            if(rsusername.equals(username) && rspwd.equals(password))
            {
                response.sendRedirect("./paper/admin/web/main.html");
            }
          }
          stat.close();
          conn.close();
          response.setHeader("refresh","0;URL=./paper/admin/login/login.html");
        }
        catch(Exception e) {
            e.printStackTrace();
        }
    }

    public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        doGet(request, response);
    }
}
