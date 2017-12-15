import java.util.ArrayList;  
import javax.servlet.ServletContext;  
import javax.servlet.http.HttpSessionAttributeListener;  
import javax.servlet.http.HttpSessionBindingEvent;  
import javax.servlet.http.HttpSessionEvent;  
import javax.servlet.http.HttpSessionListener;  
  
public class onlineListener implements HttpSessionListener,  
        HttpSessionAttributeListener {  
    ServletContext sc;  
    ArrayList list = new ArrayList();  
    public void sessionCreated(HttpSessionEvent se) {  
        sc = se.getSession().getServletContext();  ;  
    }  
  
    public void sessionDestroyed(HttpSessionEvent se) {  
        if (!list.isEmpty()) {  
            list.remove((String) se.getSession().getAttribute("userName"));  
            sc.setAttribute("list", list);  
        }  
    }  
   
    public void attributeAdded(HttpSessionBindingEvent sbe) {  
        list.add((String) sbe.getValue());  
        System.out.println(sbe.getValue());  
        sc.setAttribute("list", list);  
    }  
    
    public void attributeRemoved(HttpSessionBindingEvent arg0) {  
          
    }  
  
    public void attributeReplaced(HttpSessionBindingEvent arg0) {  
    }  
}  