package main

import (
	"log"
	"net/http"
)

// 定义一个 home handler 函数，"Hello from Snippetbox" 作为响应体
func home(w http.ResponseWriter, r *http.Request) {
	// 检查当前请求的 URL 路径是不是匹配 "/"。如果不匹配，使用
	// http.NotFound() 函数发送 404 响应给客户端
	// 重要地，我们从 handler 返回。如果我们不从 handler 返回，它将继续执行
	// 并返回 "Hello from Snippetbox" 消息。
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// 添加一个 snippetView handler 函数
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// 添加一个 snippetCreate handler 函数
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// 使用 r.Method 查检请求方法使用的是否为 POST
	if r.Method != "POST" {
		// 使用 Header().Set() 方法添加一个 'Allow: POST' 头到响应头中。
		// 第一个参数是 header name，第二个参数是 header value.
		w.Header().Set("Allow", "POST")
		// 使用 http.Error() 函数发送 405 状态码和 "Method Not Allowed"
		// 字符串作为响应体。
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// 使用 http.NewServeMux() 函数初始化一个新的 servemux，
	// 然后将主函数注册为“/”URL 模式的处理程序。
	// 在 servemux 中添加两个新的 handler 函数和相应的 URL patterns
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
