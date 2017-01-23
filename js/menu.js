
function logout(){
  AJAXGet('logout', logoutHandler);
}

function logoutHandler(){
  alert('已登出!!');
  window.location.href = 'index.html';
  
}

window.onscroll = function(){
  var menuTop = document.getElementsByTagName("body")[0].scrollTop;
  
  //console.log(menuTop);
  
  if (menuTop >= 160) {
    document.getElementById("menu").style.position = 'fixed';
    document.getElementById("menu").style.top = '0px';
  }
  else {
    document.getElementById("menu").style.position = 'absolute';
    document.getElementById("menu").style.top = '170px';
  }

}