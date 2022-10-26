package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 初始化一个包含两个文件的切片。
	// 注意：包含 base template 的必须是切片中的第一个元素
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// 使用 template.ParseFiles() 函数读取文件并保存模板。
	// 注意：我们能通过文件路径的切片作为一个可变参数？
	// 如果有错误，我们可以通过 log 输出详细的错误信息
	// 并使用 http.Error() 函数发送
	// 通用的内部服务出错的 500 错误信息给用户。
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 使用 ExecuteTemplate() 方法来写 base 模板中的内容作为响应体。
	// 其中最后一个参数代表任意一个我们想传入的动态数据，
	// 这里我们什么也不传，使用 nil 填充。
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
