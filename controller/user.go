package controller

import (
	"../service"
	"../util"
	"net/http"
)

var  userService service.UserService

func UserLogin (writer http.ResponseWriter, request *http.Request) {
	//数据库操作
	//逻辑处理
	//restapi json
	// 获取参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	user, err := userService.Login(mobile,passwd)
	if err != nil {
		util.ApiReturnFail(writer, nil, err.Error())
	} else {
		util.ApiReturnOk(writer, user, "登录成功")
	}
}


func UserRegister (writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plaionpwd := request.PostForm.Get("plaionpwd")
	password := request.PostForm.Get("passwd")

	if mobile == "" {
		util.ApiReturnFail(writer, nil, "手机号不能为空")
		return
	}


	if plaionpwd == "" || password == "" {
		util.ApiReturnFail(writer, nil, "请输入密码")
		return
	}

	if plaionpwd != password {
		util.ApiReturnFail(writer, nil, "两次密码不一致")
		return
	}

	user, err :=userService.Register(mobile, plaionpwd)
	if err != nil {
		util.ApiReturnFail(writer, nil, err.Error())
		return
	} else {
		util.ApiReturnOk(writer,user, "注册成功")
		return
	}
}
