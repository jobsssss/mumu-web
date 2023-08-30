function SwapTab(name,title,content,Sub,cur){
  $(name+' '+title).click(function(){
    $(this).addClass(cur).siblings().removeClass(cur);
    $(content+" > "+Sub).eq($(name+' '+title).index(this)).fadeIn('fast').siblings().hide();
  });
}

function message(type,msg){
			var html ='<div id="toast"><div class="weui-mask_transparent"></div><div class="weui-toast">';
			if(type==1){
				 html +='<i class="weui-icon-success-no-circle weui-icon_toast"></i>';
			}else{
				 html +='<i class="weui-icon-info-circle weui-icon_toast"></i>';
			}
			html +='<p class="weui-toast__content">'+msg+'</p></div></div>';
			$('#toast').html(html).fadeIn(100);
			setTimeout(function () {
				$('body').find('#toast').fadeOut(100);
			}, 2000);
		}


$(function(){
	new SwapTab('.tabsbar','li','.tabs-content','.tabs-body','active')
	new SwapTab('.tabsbar','li','.tabs-content','.tabs-body','active')
	
	//充值
	$('.recharge-select li').click(function(){
		$(this).addClass('active').siblings('li').removeClass('active')
	})
	//导航
	$('.icon-menu').on('click', function() {
	  $('.dropnav').toggleClass('open');
		$('.icon-menu').toggleClass('icon-menu-close');
	})
	
	//女生男生
	$('.solinks a').click(function(){
		$(this).addClass('active').siblings('a').removeClass('active')
	})
	
	//加入书架

	
	//我的书架
	$('.shelve-admin').click(function(){
		$('#NormboxHd').hide();
		$('#ActionHd').show();
		$('.shelve-bar').show();
		$('.shelve-check').show();
	})
	$('.box-cancel').click(function(){
		$('#NormboxHd').show();
		$('#ActionHd').hide();
		$('.shelve-bar').hide();
		$('.shelve-check').hide();
	})
	$(document).on('click','.btn-select',function(){
		
	})
	$(".btn-select").click(function(){
		if($(this).hasClass('active')){
			$(this).text('全选')
			$(this).removeClass('active') 
			$("input[name='bookcaseid[]']").removeAttr("checked"); 
		}else{
			$(this).text('全不选')
			$(this).addClass('active')
			$("input[name='bookcaseid[]']").prop("checked","true");
		}
		var len = $(".bookshelves-list input[type='checkbox']:checked").length
		$('#delnum').text(len)
	})
	$('.shelve-check').click(function(){
		var len = $(".bookshelves-list input[type='checkbox']:checked").length
		$('#delnum').text(len)
		$(this).parents('li').addClass('lichecked')
	})
	//打赏
	$('.gift-list li').click(function(){
		$(this).addClass('active').siblings('li').removeClass('active')
	})
	
});










