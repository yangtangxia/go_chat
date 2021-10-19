package service

import (
	"../model"
	"../util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {

}

//注册函数
func (s *UserService)Register(
	mobile,
	plaionpwd string) (user model.User, err error) {

	//参数检测
	tmp := model.User{}
	//查询数据库一条数据
	_, err = DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}

	tmp.Mobile = mobile
	tmp.Nickname = "用户"
	tmp.Sex = model.SEX_UNKNOW
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31())
	tmp.Passwd = util.MakePasswd(plaionpwd, tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Updateat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	_, err = DbEngin.InsertOne(&tmp)
	if err != nil {
		return tmp, errors.New("注册失败")
	}
	return tmp, err
}

//登录函数
func (s *UserService)Login(
	mobile,
	plaionpwd string) (user model.User, err error) {

	//查询数据
	_, err = DbEngin.Where("mobile =?", mobile).Get(&user)
	if err != nil {
		return user, errors.New("没有获取到手机号信息")
	}

	verify := util.ValidatePasswd(plaionpwd, user.Salt, user.Passwd)

	if !verify {
		return user, errors.New("密码不正确")
	}

	//刷新token
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	user.Token = token
	DbEngin.ID(user.Id).Cols("token").Update(&user)

	return user, nil
}