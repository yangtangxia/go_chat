package service

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/MySQL"
	"github.com/go-xorm/xorm"
	"../model"
)
var DbEngin *xorm.Engine

func init()  {
	drawname := "mysql"
	DsName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drawname, DsName) // 连接数据库
	if err != nil && err.Error() != "" {
		fmt.Print(err.Error())
		return
	}
	//是否显示sql语句
	DbEngin.ShowSQL(true)
	//数据库最大打开连接数
	DbEngin.SetMaxOpenConns(2)

	//自动创建user表
	DbEngin.Sync2(new(model.User), new(model.Contact))

	fmt.Println("init data base ok")
}