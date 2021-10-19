var DMcurrentFirstDate = new Date();
var DMformatDate = function(date){
 
  var year = date.getFullYear()+'-';
  var month = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
  var day = (date.getDate() < 10 ? '0'+date.getDate(): date.getDate()) + ' ';
  var week = ['星期天','星期一','星期二','星期三','星期四','星期五','星期六'][date.getDay()];
   
  return [year+month+day,week];
};
var DMaddDate= function(date,n){       
  date.setDate(date.getDate()+n);        
  return date;
};
var DMsetDate = function(date){

  var week = date.getDay()-1;
  date = DMaddDate(date,week*-1);
  DMcurrentFirstDate = new Date(date);
  var DMdatearr = [];
  var DMweekarr = [];
  for(var i = 0;i<7;i++){
    let formatdatearr = DMformatDate(i==0 ? date : DMaddDate(date,1));
    DMdatearr.push(formatdatearr[0]);
    DMweekarr.push(formatdatearr[1]);
     
  }
  return [DMdatearr,DMweekarr];
  
};

function DMlastweek(){
    return DMsetDate(DMaddDate(DMcurrentFirstDate,-7)); 
}

function DMcurrentweek(){
    return DMsetDate(DMaddDate(new Date(),0)); 
}
function DMnextweek(){

    return DMsetDate(DMaddDate(DMcurrentFirstDate,+7)); 
}

