/*author kpl 2020/1/14*/

package service

import "net/http"

// AddRouterService 添加必要的路由
func AddRouterService() {
	http.Handle("/hello", http.HandlerFunc(Test))
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}