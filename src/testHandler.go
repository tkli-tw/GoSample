//原則上，未來撰寫時，最好就是 檔名 和 裡頭的函數名 儘量相同，這樣會比較好找，好管理

package main

import (
	"fmt"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	var upass = r.URL.Query()["upass"]
	fmt.Fprintf(w, "Hi there, I love %s! Hi there, I love %s!", r.URL.Path, upass)
}
