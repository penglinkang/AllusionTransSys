/*author kpl 2020/1/14*/

package service

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// StaticDir 设置静态路由
type StaticDir struct {
	dir string
}

var config = NewStaticDir()

// NewStaticDir 静态资源服务单例
func NewStaticDir() *StaticDir {
	return &StaticDir{dir:"..\\..\\webresource"}
}

// AddRouterService 添加必要的路由
func AddRouterService() {
	http.Handle("/hello", http.HandlerFunc(Test))
	http.Handle("/index.html", http.HandlerFunc(Index))
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	buff := bytes.Buffer{}
	absPath, err := filepath.Abs(config.dir)

	if err != nil {
		fmt.Println("abspath error", err)
	}
	fmt.Println(filepath.Abs("./"))
	path :=  absPath + r.URL.String()

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	f.Read(buff.Bytes())

	w.Write(buff.Bytes())
}