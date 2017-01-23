//--------------------------------------------------------------------
// AJAX
// method: get
// 傳入 url, handler
// 回傳後 呼叫 handler(xhttp)
//--------------------------------------------------------------------
function AJAXGet(url, handler) {
  // 0. 檢查瀏覽器是否有支援AJAX
  if (window.XMLHttpRequest){}
  else{
    alert("不支援AJAX!!!請改用有支援AJAX的瀏覽器。");
  }

  // 1. 建立一個 非同步請求元件 XMLHttpRequest
  var xhttp = new XMLHttpRequest();
  
  // 2. 設定 事件處理函式 
  //xhttp.onreadystatechange = handler;
  xhttp.onreadystatechange = function(){
    //(一般簡單一點，都只設定 當 server 回應完請求之後。 也就 response 都接收完之後 的事件。)
    if (this.readyState == 4 && this.status == 200) {
      handler(this);
    }
  }
  
  // 3. 設定 URL請求的相關資訊
  //    true 為 非同步模式
  xhttp.open("GET", url, true);
  
  //4. 實際送出請求 (也就真正讓 AJAX元件 執行 網頁請求的任務)
  xhttp.send();
}


//--------------------------------------------------------------------
// AJAX
// method: post
//--------------------------------------------------------------------
function AJAXPost(url, data, handler) {
  // 0. 檢查瀏覽器是否有支援AJAX
  if (window.XMLHttpRequest){}
  else{
    alert("不支援AJAX!!!請改用有支援AJAX的瀏覽器。");
  }

  // 1. 建立一個 非同步請求元件 XMLHttpRequest
  var xhttp = new XMLHttpRequest();
  
  // 2. 設定 事件處理函式 
  //xhttp.onreadystatechange = handler;
  xhttp.onreadystatechange = function(){
    //(一般簡單一點，都只設定 當 server 回應完請求之後。 也就 response 都接收完之後 的事件。)
    if (this.readyState == 4 && this.status == 200) {
      handler(this);
    }
  }
  
  // 3. 設定 URL請求的相關資訊
  //    true 為 非同步模式
  xhttp.open("POST", url, true);
  
  //4. 實際送出請求 (也就真正讓 AJAX元件 執行 網頁請求的任務)
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  //xhttp.send("fname=Henry&lname=Ford");
  xhttp.send(data);
}


//--------------------------------------------------------------------
// 使用範例
// AJAXGet
//--------------------------------------------------------------------
/*
function send(){
    var uid,upass;
    uid = document.getElementById('uid').value;
    upass = document.getElementById('upass').value;
    
    AJAXGet("loginVerify.php?uid="+ uid + "&upass=" + upass, ajaxHandler1);
    //AJAXGet("loginVerify.json?uid="+ uid + "&upass=" + upass, ajaxHandler1);    //用來做 前端程式檢測用
}

function ajaxHandler1() {
     //document.getElementById("demo").innerHTML += "<br/>" + this.responseText;
     
     var data = JSON.parse(this.responseText);
     
     if (data.success == "ok"){
         alert("登入成功!!!");
         window.location.href = "main.html";
         
     }
     else {
         alert(data.msg);
     }
   
}
*/


//--------------------------------------------------------------------
// 使用範例
// AJAXPost
//--------------------------------------------------------------------
/*
function loginVerify(){
    var uid,upass;
    uid = document.getElementById('uid').value;
    upass = document.getElementById('upass').value;
    
    AJAXPost("loginVerify", "uid="+ uid + "&upass=" + upass, loginVerifyHandler);
}

function loginVerifyHandler(xhttp) {
    var data = JSON.parse(xhttp.responseText);
    
    if (data.success == "ok"){
       alert("登入成功!!!");
       window.location.href = "main.html";
    }
    else {
       alert(data.msg);
    }
}
*/
