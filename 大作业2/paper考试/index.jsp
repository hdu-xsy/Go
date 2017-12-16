<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ page import="java.io.*,java.util.*,java.sql.*"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<link href="css/chinaese_main.css" type="text/css" rel="stylesheet"/>
<title>杭州言汇科技有限公司</title>
<script src="js/jquery-1.3.2.js" type="text/javascript"></script>
<script src="js/globle.js" type="text/javascript"></script>
</head>

<body>
<div class="w-950">
  <div class="toper">
    <div class="logo"><img src="images/login.png" width="205" height="59" /></div>
    <div class="help_div">
      <div class="lang"><a href="#" class="chinaese">中文</a><a href="#" class="english">英语</a><a href="#" class="hanyu">韩语</a><br class="clear"/>
      </div>
      <div class="search">
        <input name="" type="text" value="站内搜索 / 输入内容后回车" class="search_txt" onclick="this.style.color='#333';this.value='';" onblur="this.style.color='#818181';this.value='站内搜索 / 输入内容后回车';"/>
      </div>
    </div>
    <br class="clear"/>
  </div>
  <div class="menu">
    <div class="menu_l"></div>
    <div class="menu_r"></div>
    <div class="menu_main">
      <div class="menu_top">
        <div class="menu_nav">
          <p class="menu_head"><a href="index.html">言汇首页</a></p>
        </div>
        <div class="menu_nav more_nav">
          <p class="menu_head"><span class="child_nav">公司介绍</span></p>
          <div class="menu_body"><a href="about.html">言汇简介</a> <a href="about.html">言汇宗旨</a> <a href="about.html">发展历程</a> <a href="about.html">企业文化</a> <a href="about.html">团队介绍</a></div>
        </div>
        <div class="menu_nav more_nav">
          <p class="menu_head"><span class="child_nav">公司新闻</span></p>
          <div class="menu_body"><a href="news.html">公司新闻</a><a href="news.html">行业动态</a></div>
        </div>
        <div class="menu_nav more_nav">
          <p class="menu_head"><span class="child_nav">产品介绍</span></p>
          <div class="menu_body"><a href="product.html">产品介绍</a><a href="product.html">产品介绍</a></div>
        </div>
        <div class="menu_nav more_nav">
          <p class="menu_head"><span class="child_nav">技术服务</span></p>
          <div class="menu_body"><a href="support.html">技术服务</a><a href="support.html">技术服务</a></div>
        </div>
        <div class="menu_nav more_nav">
          <p class="menu_head"><span class="child_nav">商业伙伴</span></p>
          <div class="menu_body"><a href="command.html">商业伙伴</a><a href="command.html">商业伙伴</a></div>
        </div>
        <div class="menu_nav more_nav">
          <p class="menu_head"><span class="child_nav">交流招聘</span></p>
          <div class="menu_body"><a href="job.html">交流招聘</a><a href="job.html">交流招聘</a></div>
        </div>
        <div class="menu_nav more_nav">
          <p class="menu_head"><span class="child_nav">合作频道</span></p>
          <div class="menu_body"><a href="cooperation.html">商务合作</a><a href="survey.html">网上调查</a><a href="question.html">在线解答</a></div>
        </div>
        <div class="menu_nav">
          <p class="menu_head"><a href="contant.html">联系我们</a></p>
        </div>
      </div>
    </div>
  </div>
  <div class="banner"><img src="images/banner.jpg" width="950" height="236" /></div>
  <div class="main m-b-10">
    <div class="main_left">
      <div class="about_us m-b-10">
        <h2>公司简介<a href="#" class="more">MORE</a></h2>
        <div class="about_us_contain"> <img src="images/about_img.png" width="113" height="101" /><span class="blod">杭州言汇科技有限公司</span>
        <%
             Class.forName("com.mysql.jdbc.Driver");
             Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/Javawebpaper","root","Xsydx886.");
             Statement stat=conn.createStatement();
             String sql;
             sql = "SELECT * FROM News where father = 1;";
             ResultSet rs = stat.executeQuery(sql);
            while(rs.next()){
              String rsnews = rs.getString("ncontent");
              out.println(rsnews);
              break;
            }
            %>
      </div>
      </div>
      <div class="lay_div">
        <h2>联系我们</h2>
        <div class="lay_contain">
          <dl class="contain_dl">
            <dt><img src="images/contant_img.png" width="113" height="101" /></dt>
            <dd class="title">杭州言汇科技有限公司</dd>
            <dd>电话：86-571-85227120</dd>
            <dd>传真：86-571-88073759</dd>
          </dl>
          <ul class="contain_ul">
            <li>地址：杭州市莫干路1418-19-2楼</li>
            <li>邮编：310011</li>
            <li>邮箱：yh@yeahoo.cn/greatabcd@gmail.com</li>
          </ul>
        </div>
      </div>
    </div>
    <div class="main_right">
      <div class="main_right_top m-b-10">
        <div class="lay_div news m-r-10">
          <h2>公司新闻<a href="#" class="more">MORE</a></h2>
          <div class="lay_contain">
            <dl class="right_dl">
              <dt><img src="images/news_img.png" width="113" height="101" /></dt>
              <dd class="title">
                <%
                sql = "SELECT * FROM News where father = 2;";
                ResultSet rs2 = stat.executeQuery(sql);
                while(rs2.next()){
                 String rstitle = rs2.getString("ntitle");
                 out.print(rstitle);
                 break;
               }
                %>
              </dd>
              <dd>
                <%
                  String rsnews = rs2.getString("ncontent");
                  out.print(rsnews);
                %>
              </dd>
              <dd>[<a href="#">查看</a>]</dd>
            </dl>
            <ul class="right_ul">
              <%
                int count = 0;
                while(rs2.next() && count<4)
                {
                  String rstitle = rs2.getString("ntitle");
                  count ++;
              %>
              <li><a href="#"><%=rstitle %></a></li>
              <%
                }
              %>
            </ul>
          </div>
        </div>
        <div class="lay_div news">
          <h2>行业新闻<a href="#" class="more">MORE</a></h2>
          <div class="lay_contain">
            <dl class="right_dl">
              <dt><img src="images/hy_news_img.png" width="113" height="101" /></dt>
              <dd class="title">消费者对医疗保健服务的需求
              </dd>
              <dd>根据医疗保健市场的调据医疗保健市场的调研出版...</dd>
              <dd>[<a href="#">查看</a>]</dd>
            </dl>
            <ul class="right_ul">
              <li><a href="#">低碳水化合物饮食可能增加心脑血管疾病的风险</a></li>
              <li><a href="#">名人代言医疗广告明年起将禁播</a>，</li>
              <li><a href="#">外贸入侵国内低端医疗器械企业前景堪忧</a></li>
              <li><a href="#">外贸入侵国内低端医疗器械企业前景堪忧</a></li>
            </ul>
          </div>
        </div>
        <br class="clear"/>
      </div>
      <%
      stat.close();
      conn.close();
      %>
      <div class="lay_div">
        <h2>最新产品<a href="#" class="more">MORE</a></h2>
        <div class="lay_contain product">
       	  <div class="move_left"><input name="" type="button"  class="move_left_btn" value=""/></div>
            <div class="move_right"><input name="" type="button"  class="move_right_btn" value=""/></div>
            <div class="porduct_list">
       	  <ul class="product_ul">
                	<li class="poduct_img"><a href="#"><img src="images/product/index_01.png" width="133" height="106" /></a></li>
                    <li><a href="#">产品名称</a></li>
              </ul>
          <ul class="product_ul">
                	<li class="poduct_img"><a href="#"><img src="images/product/index_01.png" width="133" height="106" /></a></li>
                    <li><a href="#">产品名称</a></li>
              </ul>
          <ul class="product_ul">
                	<li class="poduct_img"><a href="#"><img src="images/product/index_01.png" width="133" height="106" /></a></li>
                    <li><a href="#">产品名称</a></li>
              </ul>
          <ul class="product_ul">
                	<li class="poduct_img"><a href="#"><img src="images/product/index_01.png" width="133" height="106" /></a></li>
                    <li><a href="#">产品名称</a></li>
              </ul>
            </div>
        </div>
      </div>
    </div>
    <br class="clear"/>
  </div>
  <div class="bottom">
    <p><a href="#">首页</a> | <a href="#">联系我们</a></p>
    <p>©2011 杭州言汇信息科技有限公司 All rights reserved.</p>
  </div>
</div>
</body>
</html>
