//原則上，未來撰寫時，最好就是 檔名 和 裡頭的函數名 儘量相同，這樣會比較好找，好管理

package main

import (
	"encoding/json"
	"net/http"
)

func checkIfLogin(w http.ResponseWriter, r *http.Request) bool {
	sess := globalSessions.SessionStart(w, r)
	uid := sess.Get("uid")

	if uid != nil {
		return true
	} else {
		return false
	}

}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	uid := sess.Get("uid")

	var result = make(map[string]interface{})
	var userInfo = make(map[string]interface{})
	enc := json.NewEncoder(w)

	if uid != nil {
		result["success"] = "ok"

		userInfo["uid"] = uid
		userInfo["uname"] = uid
		result["userInfo"] = userInfo
	} else {
		result["success"] = "fail"
		result["msg"] = "尚未登入"
	}

	enc.Encode(result)
}

//{"success":"ok", "msg":"none"}
