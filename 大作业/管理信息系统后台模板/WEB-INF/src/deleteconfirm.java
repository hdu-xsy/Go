import java.io.IOException;
import java.io.PrintWriter;
import java.util.*;
import java.sql.*;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class deleteconfirm extends HttpServlet {
    private static final long serialVersionUID = 1L;

    public deleteconfirm() {
        super();
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        response.setContentType("text/html;charset=UTF-8");
        String sel = request.getParameter("delete");
        int selint = Integer.valueOf(sel).intValue();
        int count = 1;
        try {
          Class.forName("com.mysql.jdbc.Driver");
          Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javaweb","root","Xsydx886.");
          Statement stat=conn.createStatement();
          String sql = "SELECT * FROM log order by opertime desc;";
          ResultSet rs = stat.executeQuery(sql);
          while(rs.next())
          {
            if(count == selint)
            {
                String rsoperuser = rs.getString("operuser");
                String rsopertime = rs.getString("opertime");
                String rsopername = rs.getString("opername");
                String rsoperobject = rs.getString("operobject");
                String deletesql = "DELETE FROM log WHERE operuser = '"+rsoperuser+"' AND opertime = '"+rsopertime+"' AND opername = '"+rsopername+"' AND operobject = '"+rsoperobject+"';";
                int i = stat.executeUpdate(deletesql);
                break;
            }
            count ++;
          }
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
