//原則上，未來撰寫時，最好就是 檔名 和 裡頭的函數名 儘量相同，這樣會比較好找，好管理

package main

import (
	"encoding/json"
	"net/http"
)

func logout(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	//這裡只是刪除 uid 的內容，
	//理論上，最好是全部刪除才對。懶得查 怎麼刪全部了。
	sess.Delete("uid")

	var result = make(map[string]interface{})
	enc := json.NewEncoder(w)

	result["success"] = "ok"
	result["msg"] = "none"

	enc.Encode(result)
}

//{"success":"ok", "msg":"none"}
