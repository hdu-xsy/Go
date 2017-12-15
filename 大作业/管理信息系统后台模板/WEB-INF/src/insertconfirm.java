import java.io.IOException;
import java.io.PrintWriter;
import java.util.*;
import java.sql.*;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class insertconfirm extends HttpServlet {
    private static final long serialVersionUID = 1L;

    public insertconfirm() {
        super();
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        response.setContentType("text/html;charset=UTF-8");
	      PrintWriter out = response.getWriter();
        String user = request.getParameter("user");
        String time = request.getParameter("time");
        String name = request.getParameter("name");
        String object = request.getParameter("object");
        try {
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql = "INSERT INTO log(operuser,opertime,opername,operobject) VALUES('"+user+"','"+time+"','"+name+"','"+object+"')";
          int i = stat.executeUpdate(sql);
          stat.close();
          conn.close();
          response.setHeader("refresh","0;URL=./html/web/right.jsp");
        }
        catch(Exception e) {
            e.printStackTrace();
        }
    }

    public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        doGet(request, response);
    }
}
