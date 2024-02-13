

$("#send").click(function (){
  q = $("#querry").val().toLowerCase()
  min = $("#min").val()
  max = $("#max").val()
  
  if (q !== ""){
    $("tr").hide();
    $("td").filter(function() {
      return $(this).text().toLowerCase().indexOf(q) !== -1;
    }).parent().show();
  }else{
    $("tr").show();
  }
  if(min !== ""){
    $(".2").filter(function() {
      return parseInt($(this).text()) < parseInt(min);
    }).parent().hide();
  }
  if(max !== ""){
    $(".2").filter(function() {
      return parseInt($(this).text()) > parseInt(max);
    }).parent().hide();
  }
  
})
