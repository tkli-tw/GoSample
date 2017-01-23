//原則上，未來撰寫時，最好就是 檔名 和 裡頭的函數名 儘量相同，這樣會比較好找，好管理

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	//"os"
	"strings"
)

//------------------------------------------------------------------------
// 20170118 tkli
// 實作 一個簡單(而有效率的) POP3 認證，
//------------------------------------------------------------------------
/* 底下為 pop3 協定 認證時的對話
Message from server: +OK MPOP3D V6.0 at <server> starting.
Text to send: USER <uid>
Message from server: +OK Password required for user "<uid>".
Text to send: PASS <upass>
Message from server: +OK <uid> has 13269 message(s)(1310136320 octets).
Text to send: QUIT
Message from server: +OK Pop server at <server> signing off.bufio
*/
func pop3Auth(uid, upass, server, port string) bool {

	// connect to this socket
	conn, _ := net.Dial("tcp", server+":"+port) //mail server 位置
	//conn, _ := net.Dial("tcp", "mail.uch.edu.tw:110") //mail server 位置

	// read in input from conn
	reader := bufio.NewReader(conn)

	// listen for reply
	message, _ := reader.ReadString('\n')
	// 反正每個都是驗證一下 回應字串 是否是 "+OK" 起始
	if !strings.HasPrefix(message, "+OK") {
		return false
	}

	// send to socke
	// 送出 帳號名稱
	fmt.Fprintf(conn, "USER "+uid+"\n")

	message, _ = reader.ReadString('\n')
	if !strings.HasPrefix(message, "+OK") {
		return false
	}

	// 送出 密碼
	fmt.Fprintf(conn, "PASS "+upass+"\n")

	message, _ = reader.ReadString('\n')
	if !strings.HasPrefix(message, "+OK") {
		return false
	}

	// 送出 結束 連線
	fmt.Fprintf(conn, "QUIT\n")

	return true

}

//------------------------------------------------------------------------
// 20170118 tkli
// 實作 一個簡單 db 認證，
//------------------------------------------------------------------------
func dbAuth(uid, upass, server, port string) bool {

	return true

}

//------------------------------------------------------------------------
// 20170118 tkli
// 驗證 帳密
//
// 20170121
// 通過後，在 session 中加入 uid 的值，作為 辨識 使用者是否已登入的 證明
// 另一方面，有需要的人 可自行加入一些資料到 session 中
//
// 另外，若未來 官方有出 session 的標準函式庫，最好改用之。
//------------------------------------------------------------------------
func loginVerify(w http.ResponseWriter, r *http.Request) {
	// 	var uid = r.URL.Query()["uid"][0]       //這個是直接抓 url 上的字串
	// 	var upass = r.URL.Query()["upass"][0]

	r.ParseForm() //注意，使用 Form.Get 前需要利用這個做 解析的動作
	var uid = r.Form.Get("uid")
	var upass = r.Form.Get("upass")

	var authResult = pop3Auth(uid, upass, "mail.uch.edu.tw", "110")

	//map 資料型別
	//result 用來 輸出 驗證結果
	var result = make(map[string]interface{})

	if authResult == true {
		result["success"] = "ok"
		result["msg"] = "none"

		sess := globalSessions.SessionStart(w, r)
		//defer sess.SessionRelease(w)
		sess.Set("uid", uid)

	} else {
		result["success"] = "fail"
		//result["msg"] = "帳密錯誤 " + uid + "," + upass
		result["msg"] = "帳密錯誤 "
	}

	jsonStr, _ := json.Marshal(result)
	fmt.Fprint(w, string(jsonStr))
}

//{"success":"ok", "msg":"none"}
