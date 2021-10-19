
let dmbasepathurl = 'http://'+window.location.host;

let DMingAjax = {
	//ajax请求数据
	method: function(obj) {

		
		var dmdic = obj.dict;
		let basepath = obj.url;
		$.ajax({
			type: "POST",
			url: dmbasepathurl+basepath,
			dataType: "json",
			data: dmdic,
			timeout: 20000,
			beforeSend:function(){
                layer.load(2);
            },
			success: function(data) {
				if (data.status=='notlogin') {
					if (window != top) {
						top.location.href = '/admin/login/index'; 
					}
					else
					{
						window.location.href = '/admin/login/index';
					}

					
					
				}
				else
				{
					obj.success(data);
				}
				
			},
			error: function(error) {
				obj.fail(error);
			},
            complete:function(){
                layer.closeAll('loading');
            }
		});
	}
}
//格式化填充数据000000001
function dmPrefixInteger(num, length) {
	if (num>99999999) {
	  return num;
	}
	return (Array(length).join('0') + num).slice(-length);
}
function dmzhengzemoney(num,iszero){
	//iszero 0包括0 1不包括0 
	var reg =/^(([1-9]\d*)|\d)(\.\d{1,2})?$/;
	if (reg.test(num)){
	  if (iszero&&parseFloat(num)==0) {
	     return false;
	  }
	  return true;

	}
	else{
	  return false;
	}
}
function dmzhengzezhengshu(num){
	//iszero 0包括0 1不包括0 
	var reg =/(^[1-9]\d*$)/;
	if(reg.test(num)){
	    return true;
	}
	else{
	  return false;
	}
}
function dmGetLocalTime(nS) {     
    var date = new Date(nS*1000);
    Y = date.getFullYear() + '-';
    M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
    D = (date.getDate() < 10 ? '0'+date.getDate(): date.getDate()) + ' ';
    h = (date.getHours() < 10 ? '0'+date.getHours(): date.getHours())+ ':';
    m = (date.getMinutes() < 10 ? '0'+date.getMinutes(): date.getMinutes())+ ':';
    s = (date.getSeconds() < 10 ? '0'+date.getSeconds(): date.getSeconds()); 
    return Y+M+D+h+m+s;
}
function dmGetLocaldate(nS) {     
	var date = new Date(nS*1000);
	Y = date.getFullYear() + '-';
	M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
	D = date.getDate()<10?'0'+date.getDate():date.getDate() + ' ';

	return Y+M+D;

}
function dmGetLocaldateweek(nS){
	var date = new Date(nS*1000);
	var week = ['星期天','星期一','星期二','星期三','星期四','星期五','星期六'][date.getDay()];
	return week;
}

function dmGetLocaldatetime(nS){
	var date = new Date(nS*1000);
   
    h = (date.getHours() < 10 ? '0'+date.getHours(): date.getHours())+ ':';
    m = (date.getMinutes() < 10 ? '0'+date.getMinutes(): date.getMinutes());
   
    return h+m;
}
function dmgetLocalTimeadd(nS) {     
	var date = new Date(nS);
	Y = date.getFullYear() + '-';
	M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
	D = date.getDate()<10?'0'+date.getDate():date.getDate() + ' ';

	return Y+M+D;

}      
function dmgetLocalTimedateadd() {     
    var date = new Date();
    Y = date.getFullYear() + '-';
    M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
    D = date.getDate() + ' ';
    h = (date.getHours() < 10 ? '0'+date.getHours(): date.getHours())+ ':';
    m = (date.getMinutes() < 10 ? '0'+date.getMinutes(): date.getMinutes())+ ':';
    s = (date.getSeconds() < 10 ? '0'+date.getSeconds(): date.getSeconds()); 
    return Y+M+D+h+m+s;

}      
function dmGetcurrentdate(day){
	var today = new Date();
	var targetday_milliseconds=today.getTime() + 1000*60*60*24*day;
	today.setTime(targetday_milliseconds); //注意，这行是关键代码
	var tYear = today.getFullYear();
	var tMonth = today.getMonth();
	var tDate = today.getDate();
	tMonth = doHandleMonth(tMonth + 1);
	tDate = doHandleMonth(tDate);
	return tYear+"-"+tMonth+"-"+tDate;
}

function doHandleMonth(month){
  var m = month;
  if(month.toString().length == 1){
   m = "0" + month;
  }
  return m;
}