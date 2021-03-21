/**
 * @author  like
 * @date  2021/3/8 12:52
 * @description CI/CD测试程序
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	open, err2 := os.OpenFile("log/test.log", os.O_RDWR, 0666)
	if err2 != nil {
		fmt.Println(err2)
	}
	log.Default()
	logger := log.New(open, "", log.LstdFlags)
	logger.Println("log")
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		w.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
