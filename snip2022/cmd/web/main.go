package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// 创建一个文件服务器指向 "./ui/static/" 文件夹。
	// 注意 http.Dir 函数给的路径是相对于项目根目录的。
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// 使用 mux.Handle() 函数将文件服务器注册为所有以“/static/”开头的
	// URL 路径的处理程序。
	// 对于匹配路径，我们在请求到达文件服务器之前去除“/static”前缀。
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// 注册其他应用路由
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
