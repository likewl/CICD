/**
 * @author  like
 * @date  2021/3/8 12:52
 * @description CI/CD测试程序
 */
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("时间123"))
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
