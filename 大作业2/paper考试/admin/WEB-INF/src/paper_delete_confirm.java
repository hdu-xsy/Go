import java.io.IOException;
import java.io.PrintWriter;
import java.util.*;
import java.sql.*;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class paper_delete_confirm extends HttpServlet {
    private static final long serialVersionUID = 1L;

    public paper_delete_confirm() {
        super();
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        response.setContentType("text/html;charset=UTF-8");
        String sel = request.getParameter("id");
        int selint = Integer.valueOf(sel).intValue();
        int count = 1;
        try {
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javawebpaper","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql = "SELECT * FROM users order by Id;";
          ResultSet rs = stat.executeQuery(sql);
          while(rs.next())
          {
            if(count == selint)
            {
                String rsid = rs.getString("Id");
                String deletesql = "DELETE FROM users WHERE id = '"+rsid+"';";
                int i = stat.executeUpdate(deletesql);
                break;
            }
            count ++;
          }
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
