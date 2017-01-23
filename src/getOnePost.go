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
//
// 其實 這個完全改自 getPostList，只修改了 sql ，
// 並自 request 中抓取 pno 參數
//------------------------------------------------------------------------
func getOnePost(w http.ResponseWriter, r *http.Request) {

	if !checkIfLogin(w, r) {
		fmt.Fprint(w, "{\"result\":\"fail\", \"msg\":\"尚未登入\"}")
		return
	}

	dbConn, _ := db.GetConn()

	var pno = r.URL.Query()["pno"][0] //上面的是 抓post 的資料，這個是 抓 Get 的資料
	sql := "select pno, users.userid, users.username, time, title,content from post, users where replyTo=0 and post.userid = users.userid and pno=? "
	rows, _ := dbConn.Query(sql, pno)

	jsonStr, _ := db.GetJsonFromRows(rows)
	fmt.Fprint(w, string(jsonStr))

	dbConn.Close()
}
