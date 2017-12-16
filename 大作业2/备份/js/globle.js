// JavaScript Document


//menu
$(document).ready(function(){   
    $(".menu_head").click(function(){   
        $(this).parent().find(".menu_body").slideDown('fast').show();    
        $(this).parent().hover(function(){
										
			},function(){   
				$(this).parent().find(".menu_body").slideUp('slow');   
			});   
		}).hover(function(){   
				$(this).css({color:"#F4AE34"});   
		   }, function(){   
				$(this).css({color:"#fff"});   
    });
	
	//left nav
	$(".left_menu h2").parent().find(".left_menu_body").hide();
	$(".left_menu h2").toggle(function(){   
        $(this).parent().find(".left_menu_body").slideDown('fast');
		$(this).addClass("open");
    },function(){
		$(this).parent().find(".left_menu_body").slideUp('fast');
		$(this).removeClass("open");
	}); 
	
});  