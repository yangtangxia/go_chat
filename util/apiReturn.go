package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//定义结构体
type H struct {
	Code int `json:"code"` // 转换为小写
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"` // 把大写D转小写， 返回数据data为空时不显示data字段
	Rows interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

func ApiReturnOk(writer http.ResponseWriter, data interface{}, msg string)  {
	returnApi(writer, 0, data, msg)

}

func ApiReturnFail(writer http.ResponseWriter, data interface{}, msg string)  {
	returnApi(writer, -1, data, msg)

}

func ApiOkList(writer http.ResponseWriter, lists interface{}, total interface{})  {
	apiList(writer, 0, lists, total)
}


func returnApi(writer http.ResponseWriter, code int, data interface{}, msg string)  {
	writer.Header().Set("Content-Type", "application/json")
	//设置200状态
	writer.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Msg: msg,
		Data: data,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	writer.Write(ret)

	return
}

func apiList(writer http.ResponseWriter, code int, data interface{}, total interface{} )  {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Rows:data,
		Total:total,
	}

	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	writer.Write(ret)
	return
}