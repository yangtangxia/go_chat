package main

import (
	"./controller"
	_ "github.com/go-sql-driver/MySQL"
	"html/template"
	"log"
	"net/http"
)



//var DbEngin *xorm.Engine
//var  userService service.UserService
//func init()  {
//	drawname := "mysql"
//	DsName := "root:root@(127.0.0.1:3306)/chat?chrset=utf8"
//	DbEngin, err := xorm.NewEngine(drawname, DsName)
//	if err != nil {
//		fmt.Print(err.Error())
//		return
//	}
//	//是否显示sql语句
//	DbEngin.ShowSQL(true)
//	//数据库最大打开连接数
//	DbEngin.SetMaxOpenConns(2)
//
//	//自动创建user表
//	//DbEngin.Sync2(new(user))
//
//	fmt.Println("init data base ok")
//}



//万能模板渲染函数
func registerView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}

	 for _, v := range tpl.Templates() {
	 	tplname := v.Name()
	 	http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	 }
}

func main() {
	//绑定请求和处理函数
	//http.HandleFunc("/user/login", controller.UserLogin)

	// 提供静态资源目录支持
	http.Handle("/", http.FileServer(http.Dir(".")))
	//渲染登录模板


	//http.HandleFunc("/user/login.html", func(w http.ResponseWriter,
	//	r *http.Request) {
	//	tpl, err := template.ParseFiles("view/user/login.html")
	//	if nil != err {
	//		//打印并直接退出
	//		log.Fatal(err.Error())
	//	}
	//
	//	tpl.ExecuteTemplate(w, "/user/login.html", nil)
	//})
	/*
	http.HandleFunc("/user/register.shtml", func(writer http.ResponseWriter,
		request *http.Request) {
		tpl,err := template.ParseFiles("view/user/register.html")
		if err != nil {
			println(1111)
		}

		tpl.ExecuteTemplate(writer, "/user/register.shtml", nil)
	})
*/
	registerView()
	//注册
	http.HandleFunc("/user/register", controller.UserRegister)
	//登录
	http.HandleFunc("/user/login", controller.UserLogin)

	// 添加好友
	http.HandleFunc("/chat/addfriends", controller.Addfriends)
	//好友列表
	http.HandleFunc("/chat/friends", controller.Friends)
	
	//mysql 依赖包
	//go get github.com/go-xorm/xorm
	// MySQL驱动
	//go get github.com/go-sql-driver/mysql
	//提供指定文件目录支持
	//http.Handle("/asset/", http.FileServer(http.Dir(".")))
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}