import java.io.IOException;
import java.io.PrintWriter;
import java.util.*;
import java.sql.*;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class paper_motify_confirm extends HttpServlet {
    private static final long serialVersionUID = 1L;

    public paper_motify_confirm() {
        super();
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        response.setContentType("text/html;charset=UTF-8");
	      PrintWriter out = response.getWriter();
        String id = request.getParameter("checkbox");
        String user = request.getParameter("username");
        String pwd = request.getParameter("pwd");
        try {
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javawebpaper","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String updatesql = "update users set username = '"+user+"',pwd = '"+pwd+"' where Id = "+id+";";
          int i = stat.executeUpdate(updatesql);
          stat.close();
          conn.close();
          response.setHeader("refresh","0;URL=./paper/admin/web/tab/tab.jsp");
        }
        catch(Exception e) {
            e.printStackTrace();
        }
    }

    public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        doGet(request, response);
    }
}
