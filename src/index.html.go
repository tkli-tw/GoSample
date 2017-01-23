//如果，要將 .html 存成 字串時，建議用這樣的方式，
//方便管理，一目了然。
//而且也很方便使用，
//要測試維護網頁時，就去掉檔名的.go，編輯測試好，再加上 .go (以及 上面這個 字串的宣告) 即可
//
//用這個的好處是，
//前端的程式及網頁內容 也可以整合到 伺服器程式檔裡面，比較不怕被別人隨意修改。
//使用者(這裡指伺服器程式使用者->委託撰寫這個程式的人)如果要改網頁內容時，就變成也要經過 系統開發者一手，可以賺點外快了(呵呵)

package main

var index string = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>雲端範例系</title>


<script type="text/javascript" src="js/AJAX.js"></script>
<script type="text/javascript" src="js/htmlLoader.js"></script>

<script>
//--------------------------------------------------------------------
// 檢查是否已登入
// 已登入則自動導向至 home.html
//--------------------------------------------------------------------
function checkLogin(){
    AJAXGet("checkLogin", checkLoginHandler);
}

function checkLoginHandler(xhttp){
     var data = JSON.parse(xhttp.responseText);
     
     if (data.success == "ok"){
         //alert("登入成功!!!");
         window.location.href = "main.html";
     }
}

//--------------------------------------------------------------------
// 初始化頁面
//--------------------------------------------------------------------
function initPage(){
    checkLogin();
    htmlLoader();
    
    document.getElementById("uid").focus();
}


//--------------------------------------------------------------------
// 帳號檢查
//--------------------------------------------------------------------
function loginVerify(){
    var uid,upass;
    uid = document.getElementById('uid').value;
    upass = document.getElementById('upass').value;
    
    AJAXPost("loginVerify", "uid="+ uid + "&upass=" + upass, loginVerifyHandler);
    //AJAXGet("loginVerify.json?uid="+ uid + "&upass=" + upass, ajaxHandler1);    //用來做 前端程式檢測用
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

    
</script>

<style>


</style>

</head>
  
<body onload = "initPage();">

<header load-html='header.html'></header>

<section id='content' style='border:1px solid black;font-family:Microsoft JhengHei;width:924px;height:550px;margin: auto;border-radius: 3px;'>
    
    <div style='position:absolute;'>
    
    <article id='introduce' style='position:absolute;height:575px;width:550px;left:50px;top:75px;'>

      <p style="font-size:24pt">歡迎使用<BR>健行科技大學資訊工程系<BR>雲端範例系統</p>
      
      <p style="font-size:14pt">使用 GO 語言</p>  

    </article>  

    <article id='loginUI' style="position:absolute;height:175px;width:270px;background-color:#FFFFFF;left:575px;top:150px;" id="right" >

        <form id="form1" name="form1" method="post" action="login.php">
          <fieldset style="width:260px;border-radius:8px;">
          	<legend>使用SIP系統之帳號與密碼登入</legend>
        帳號：
        <input type="text" name="uid" id="uid" />
        </label>
            <p>密碼：
              <label>
              <input type="password" name="upass" id="upass" />
              </label>
            <p>
              <label>
              <input type="button" name="submit" id="submit" value="登入" onclick='loginVerify();' />
              </label>
            </p>
        </fieldset>
        </form>
    </article>
    
    </div>

</section>

 
<footer load-html='footer.html'></footer>

</body>  
</html>
`
