//------------------------------------------------------------------------
// 20170120 tkli
// 底下的參考連結 是 go 對 mysql 存取的 函式庫及參考說明
// https://github.com/go-sql-driver/mysql
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md
//------------------------------------------------------------------------

package main

import (
	"./db"
	"fmt"
	"net/http"
)

//------------------------------------------------------------------------
// 20170121 tkli
// 對 mysql 資料表 query 的範例
//------------------------------------------------------------------------
func getPostList(w http.ResponseWriter, r *http.Request) {

	if !checkIfLogin(w, r) {
		fmt.Fprint(w, "{\"result\":\"fail\", \"msg\":\"尚未登入\"}")
		return
	}

	dbConn, _ := db.GetConn()

	//第一個參數的 prepare字串，其內有幾個 ?，後面就要多接幾個 參數，
	//執行時，會自動被代入到 prepare字串中，很方便
	//想想，十年前做這個事的時候，還要土法煉鋼的 用 自己設定特定字符，再用 字串取代函數 去把參數內容替代進去。還真是麻煩。
	sql := "select pno, userid, time, title from post where replyTo=? "
	rows, _ := dbConn.Query(sql, 0)

	jsonStr, _ := db.GetJsonFromRows(rows)
	fmt.Fprint(w, string(jsonStr))

	dbConn.Close()
}
