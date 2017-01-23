//------------------------------------------------------------------------
// 20170120 tkli
// 底下的參考連結 是 go 對 mysql 存取的 函式庫及參考說明
// https://github.com/go-sql-driver/mysql
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md
//------------------------------------------------------------------------

package main

import (
	"./db"
	"encoding/json"
	"fmt"
	"net/http"
)

//------------------------------------------------------------------------
// 20170121 tkli
// 對 mysql 資料表 query 的範例
//------------------------------------------------------------------------
func doChangeMemberData(w http.ResponseWriter, r *http.Request) {

	if !checkIfLogin(w, r) {
		fmt.Fprint(w, "{\"result\":\"fail\", \"msg\":\"尚未登入\"}")
		return
	}

	sess := globalSessions.SessionStart(w, r)
	uid := sess.Get("uid")

	r.ParseForm() //注意，使用 Form.Get 前需要利用這個做 解析的動作
	var uname = r.Form.Get("uname")
	var age = r.Form.Get("age")
	var phone = r.Form.Get("phone")

	dbConn, _ := db.GetConn()

	sql := "update users set username=?, age=?, phone=? where userid = ? "
	stmt, err := dbConn.Prepare(sql)
	rec, err := stmt.Exec(uname, age, phone, uid)

	var result = make(map[string]interface{})

	if err == nil {
		result["result"] = "ok"
		affect, _ := rec.RowsAffected()
		result["msg"] = affect
	} else {
		result["result"] = "fail"
		result["msg"] = err
	}

	//底下這個是為了 產生 json 文件的
	jsonStr, _ := json.Marshal(result)
	fmt.Fprint(w, string(jsonStr))

	dbConn.Close()
}
