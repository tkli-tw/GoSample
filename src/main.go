//------------------------------------------------------------------------
// 20170118 tkli
// 底下的參考連結 感覺蠻不錯的，不想找來找去的話，可直接參考這裡的 (雖然我是到處亂找的)
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md
//
// 執行前的注意事項：
// 1. 安裝 session 套件,
// 1.1 go get github.com/astaxie/session
// 1.2 go install github.com/astaxie/session
// 1.3 其他 session 注意事項，請自行參考 專案中的做法 和 註解。
//
// 2. 安裝 mysql 套件，
// 2.1 go get github.com/go-sql-driver/mysql
// 2.2 go install github.com/go-sql-driver/mysql
//
// 3. 啟動 mysql，
// 3.1 若是 一般系統，應該是 mysql -u <uer> -p
// 3.2 在 cloud9 裡頭是， mysql-ctl cli，不用密碼
//
// 4. 建立資料庫，資料表，以及塞入 使用者資料
// 4.1 請參見 sql.txt
//------------------------------------------------------------------------

package main

import (
	"fmt"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"io/ioutil"
	"net/http"
	"strings"
)

//------------------------------------------------------------------------
// 20170118 tkli
// 透過一個 陣列 將 有讀過的 html 內容都 快取在記憶體裡，
// 可減少 對 硬碟的 IO， (降低伺服器的 負載)
// 也可 提昇反應速度，
//
// 壞處是
// 如果 html 有變更時，需要重啟伺服器(也就是要重新執行 go)
//
// 未來可以嘗試修改成，輸出前先讀一下 檔案的最後更新時間，
// 不過，這樣也是會造成 IO 次數增加，
// 所以，可再評估，
//
// 事實上，如果有 html 變更了，就重新執行 伺服器，其實也未嘗不可。
// (不過，伺服器重新執行會牽涉到 如果有 使用者是在已登入的情況，就會變成被踢掉，要重登了)
//------------------------------------------------------------------------
var htmlCache = make(map[string]string)
var cacheEnable = false //*******************************要特別注意，可在開發期間把這個設定為 false，免得要一直重啟程式 20170121 tkli

//------------------------------------------------------------------------
// 使用 session 時要注意
// 1. 先安裝
// 1.1 go get github.com/astaxie/session
// 1.2 go install github.com/astaxie/session
//
// 2. 宣告底下的 globalSessions
// 2.1 然後，透過 呼叫 session.NewManager("memory", "gosessionid", 3600) 來初始化
//
// 3. 使用時，透過 sess := globalSessions.SessionStart(w, r) 取得 session 物件
//------------------------------------------------------------------------
var globalSessions *session.Manager

func initSession() {
	//globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)   //這是新版的用法 https://github.com/astaxie/beego/tree/master/session
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

//------------------------------------------------------------------------
// 20170118 tkli
// 主要的 預設 handler
// 基本上，如果沒有 設定的 路徑(url path)，全部都由這個 handler 處理
// 這個 handler 處理時，
//
// 若遇到 .html 則先檢查 該頁面是否已在 快取陣列中，
// 有則直接回傳，
// 無則，讀取檔案內容至快取陣列，然後回傳。
//------------------------------------------------------------------------
func defaultHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" || r.URL.Path == "" {
		indexHandler(w, r)
		return
	}

	//讀取 指定路徑的檔案 回傳給 client
	if cacheEnable && r.URL.Path[len(r.URL.Path)-5:] == ".html" {
		filename := strings.Trim(r.URL.Path[1:], " ")
		//fmt.Fprintf(w, filename)
		if htmlCache[filename] == "" {
			b, err := ioutil.ReadFile(filename) // just pass the file name
			if err != nil {
				fmt.Print(err)
				return
			} else {
				htmlCache[filename] = string(b)
			}
		}
		fmt.Fprint(w, htmlCache[filename]) //注意哦!!! 用 Fprintf 的話，會變成格式化輸出，這時如果 html 裡有 % 會被解釋成 格式化輸出的參數
	} else {
		http.ServeFile(w, r, r.URL.Path[1:])
	}
}

//------------------------------------------------------------------------
// 20170118 tkli
// 這個只是做個 範例，也就是，可以把 整個 html 寫成一個 字串，
//------------------------------------------------------------------------
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, index) //注意哦!!! 用 Fprintf 的話，會變成格式化輸出，這時如果 html 裡有 % 會被解釋成 格式化輸出的參數
}

//------------------------------------------------------------------------
// 20170118 tkli
// 這個單獨 對某個 網址路由的情況。
//------------------------------------------------------------------------
func mainHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("main.html") // just pass the file name
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Fprint(w, string(b))
	}
}

//------------------------------------------------------------------------
// 20170118 tkli
// 主程式
// 設定各項路由，並且啟動伺服器
//------------------------------------------------------------------------
func main() {

	http.HandleFunc("/", defaultHandler)

	http.HandleFunc("/index.html", indexHandler)
	//http.HandleFunc("/main.html", mainHandler)

	http.HandleFunc("/loginVerify", loginVerify)
	http.HandleFunc("/checkLogin", checkLogin)
	http.HandleFunc("/logout", logout)

	http.HandleFunc("/getPostList", getPostList)
	http.HandleFunc("/getOnePost", getOnePost)
	http.HandleFunc("/getReplyList", getReplyList)
	http.HandleFunc("/getMemberData", getMemberData)

	http.HandleFunc("/doPostArticle", doPostArticle)
	http.HandleFunc("/doReplyArticle", doReplyArticle)

	http.HandleFunc("/doChangePass", doChangePass)
	http.HandleFunc("/doChangeMemberData", doChangeMemberData)

	http.HandleFunc("/test", testHandler)

	initSession()
	http.ListenAndServe(":8080", nil)
}
