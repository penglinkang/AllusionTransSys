/*author kpl 2020/1/14 创建工程*/
// main 用于server的主入口
package main

import (
	"log"
	"net/http"
	"service"
	"strconv"
)

type admin struct {
	addr string
	port int
}

// NewAdminServer application单例
func NewAdminServer() *admin {
	return &admin{}
}

// 设置端口 默认为80
func (a *admin)SetdefaultConfig(num interface{}) {
	if port, ok := num.(int); ok {
		a.port = port
	} else {
		a.port = 80
	}
}

func main() {
	app := NewAdminServer()
	app.SetdefaultConfig(nil )
	// 添加必要的路由
	service.AddRouterService()
	// 构造通道进行阻塞
	var ch = make(chan int, 1)
	// 闭包进行监听
	go func() {
		log.Println(http.ListenAndServe(app.addr + ":" + strconv.Itoa(app.port), nil))
		ch <- 2
	}()
	<- ch
}
