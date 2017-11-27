
import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

public class filtertest extends HttpServlet {
 
  private String message;

  public void init() throws ServletException
  {
      message = "GBK乱码测试 若这句话无问号或者乱码 则过滤器正常";
  }

  public void doGet(HttpServletRequest request,
                    HttpServletResponse response)
            throws ServletException, IOException
  {
      
      response.setContentType("text/html");

     
      PrintWriter out = response.getWriter();
      out.println("<h1>" + message + "</h1>");
  }
  
  public void destroy()
  {
      // 什么也不做
  }
}