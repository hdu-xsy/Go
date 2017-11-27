import java.io.*;
import java.util.*;
import java.sql.*;
import javax.servlet.*;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;


 
public class GBKFilter implements Filter {
 
@Override
public void destroy() {

}
 
@Override
public void doFilter(ServletRequest req, ServletResponse resp,
FilterChain chain) throws IOException, ServletException {

HttpServletRequest request=(HttpServletRequest)req;
HttpServletResponse response=(HttpServletResponse)resp;
 
request.setCharacterEncoding("GBK");
response.setContentType("text/html;charset=GBK");
 
chain.doFilter(req, resp);
 
}
 
@Override
public void init(FilterConfig filterConfig) throws ServletException {
// TODO Auto-generated method stub
 
}
}