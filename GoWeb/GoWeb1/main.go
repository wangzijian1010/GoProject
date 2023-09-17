package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

// 数据库连接信息，请替换为您自己的数据库信息
const (
	DBUsername = "root"
	DBPassword = "123456"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "test"
)

var db *sql.DB

// 登录页面处理函数
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 显示登录表单页面
	tmpl, err := template.ParseFiles("D:\\gopro\\GoWeb\\GoWeb1\\login.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// 处理登录表单提交
// 处理登录表单提交
func loginSubmitHandler(w http.ResponseWriter, r *http.Request) {
	// 从表单获取用户名和密码
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 连接到数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 查询数据库以验证用户名和密码
	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// 用户不存在
			http.Error(w, "此用户不存在，请注册", http.StatusUnauthorized)
		} else {
			// 处理查询错误
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// 检查密码是否匹配
	if storedPassword == password {
		// 登录成功
		// http的重定向函数定位到登录成功界面
		http.Redirect(w, r, "/login-success?username="+username, http.StatusSeeOther)
	} else {
		// 密码错误
		http.Error(w, "密码错误，请重新输入", http.StatusUnauthorized)
	}
}

// 创建login成功的处理函数
// 处理登录成功后的页面显示
func loginSuccessHandler(w http.ResponseWriter, r *http.Request) {
	// 获取用户名
	username := r.URL.Query().Get("username")

	// 渲染登录成功页面
	tmpl, err := template.ParseFiles("login_success.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
	}{
		Username: username,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	// 连接数据库
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	var err error
	db, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 初始化路由
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register", registerHandler)
	// 在 main 函数中添加以下路由
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login-submit", loginSubmitHandler)
	http.HandleFunc("/login-success", loginSuccessHandler)

	// 启动服务器
	port := ":8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 首页处理函数
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 仅显示一个链接到注册页面的链接
	fmt.Fprint(w, `<a href="/register">Go to Registration</a>`)
}

// 注册页面处理函数
// 注册页面处理函数
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// 显示注册表单页面
		tmpl, err := template.ParseFiles("D:\\gopro\\GoWeb\\GoWeb1\\register.html")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// 处理注册逻辑
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 添加输入验证逻辑
		if username == "" || password == "" {
			// 显示错误页面
			tmpl, err := template.ParseFiles("D:\\gopro\\GoWeb\\GoWeb1\\error.html")
			if err != nil {
				log.Fatal(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, "用户名和密码不能为空，请重新输入")
			return
		}

		// 连接到数据库
		db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// 查询数据库以检查是否已经存在相同的用户名
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username=?", username).Scan(&count)
		if err != nil {
			// 处理查询错误
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if count > 0 {
			// 用户名已经存在，显示错误消息
			tmpl, err := template.ParseFiles("D:\\gopro\\GoWeb\\GoWeb1\\error.html")
			if err != nil {
				log.Fatal(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, "用户名已经存在，请选择另一个用户名")
			return
		}

		// 如果用户名不存在，将数据插入数据库
		_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// 显示注册成功页面
		tmpl, err := template.ParseFiles("D:\\gopro\\GoWeb\\GoWeb1\\registersuccess.html")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data := struct {
			Username string
			Password string
		}{
			Username: username,
			Password: password,
		}
		tmpl.Execute(w, data)
	}
}
