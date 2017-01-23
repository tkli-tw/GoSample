//------------------------------------------------------------------------
// 20170120 tkli
// 底下的參考連結
// https://github.com/go-sql-driver/mysql
//------------------------------------------------------------------------
//
// 要先安裝套件
// Simple install the package to your $GOPATH with the go tool from shell:
// $ go get github.com/go-sql-driver/mysql
// Make sure Git is installed on your machine and in your system's PATH.
// $ go install github.com/go-sql-driver/mysql

package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 20170120 tkli
//連線前 請特別注意，mysql 是否是執行中，不然試了半天，也不會有反應
func GetConn() (*sql.DB, error) {
	//[username[:password]@][protocol[(address:port)]]/dbname[?param1=value1&...&paramN=valueN]
	username := "tkli"
	password := ""
	address := "0.0.0.0"
	dbname := "test2"

	conStr := username + ":" + password + "@tcp(" + address + ":3306)/" + dbname + "?charset=utf8"
	dbConn, err := sql.Open("mysql", conStr)

	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println("已連線至資料庫")
	}
	// fmt.Println(conStr)
	// fmt.Println(dbConn)
	return dbConn, err
}

// 20170122 tkli
// 參考：http://stackoverflow.com/questions/21986780/is-it-possible-to-retrieve-a-column-value-by-name-using-golang-database-sql
// go 是一個 嚴格型別的程式語言，要處理 json 的格式化輸出 還真是困難呀!!
// 因為，要解決 如何 自動化的 讀入各種型別的資料 有點複雜
//
// 將這個 功能 抽出來是因為 基本上我們每個 後端的輸出都是 json格式，所以，有使用到資料庫時，利用這個就能快速的產生結果了
func GetJsonFromRows(rows *sql.Rows) ([]byte, error) {
	var result = make(map[string]interface{})
	var dataList []interface{}

	columnNames, _ := rows.Columns() // []string{"id", "name"}

	//20170123 tkli
	//底下這個部份 只是為了 取得 scanArgs[i] 的型別，參考：https://codedump.io/share/QcvmbfcBk8J5/1/dumping-mysql-tables-to-json-with-golang
	//不過，我還搞不清楚 為什麼會 work... 哈，真是不負責任呀!!! 沒辦法囉，就是速成的嘛...有輿趣的自己弄清楚吧!!!
	scanArgs := make([]interface{}, len(columnNames))
	values := make([]interface{}, len(columnNames))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	//var postIndex = 0
	for rows.Next() {
		rows.Scan(scanArgs...)

		record := make(map[string]interface{})

		for i, col := range values {
			if col != nil {
				switch col.(type) {
				default:
				case bool:
					record[columnNames[i]] = col.(bool)
				case int:
					record[columnNames[i]] = col.(int)
				case int64:
					record[columnNames[i]] = col.(int64)
				case float64:
					record[columnNames[i]] = col.(float64)
				case string:
					record[columnNames[i]] = col.(string)
				case []byte: // -- all cases go HERE!
					record[columnNames[i]] = string(col.([]byte))
					//case time.Time:
					// record[columns[i]] = col.(string)
				}
			}
		}

		dataList = append(dataList, record)
	}

	result["result"] = "ok"
	result["data"] = dataList

	//回傳 json 文件，但 型別是 []byte, error，如果要用 string 輸出時，記得轉 string
	return json.Marshal(result)

}
