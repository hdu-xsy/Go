import java.io.IOException;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.Filter;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
public class filter implements Filter{
  public filter(){
	super();
  }
  private FilterConfig filterConfig;
  public void init(FilterConfig filterConfig)throws ServletException{
	this.filterConfig=filterConfig;
  }
  public void doFilter(ServletRequest req,ServletResponse res,FilterChain filterChain){
	try{
		HttpServletRequest request=(HttpServletRequest)req;
		HttpServletResponse response=(HttpServletResponse)res;
		HttpSession session=request.getSession();
		//code when login
		filterChain.doFilter(req,res);
		//code when logout
	}catch(IOException e){
		e.printStackTrace();
	}catch(ServletException e){
		e.printStackTrace();
	}
  }
  public void destroy(){
	  //put destroy string in log
	  //put your code
  }
}
