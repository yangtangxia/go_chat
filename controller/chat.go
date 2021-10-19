package controller

import (
	"../args"
	"../service"
	"../util"
	"net/http"
)

var contactService service.ContactService

func Addfriends(writer http.ResponseWriter, r *http.Request)  {
	var arg args.ContactArg
	util.Bind(r, &arg)

	err := contactService.AddFriends(arg.Owenid, arg.FriendId)
	if err != nil {
		util.ApiReturnFail(writer, nil, err.Error())
	} else {
		util.ApiReturnOk(writer, arg.Owenid, "添加好友成功")
	}
}

func Friends(writer http.ResponseWriter, r *http.Request)  {
	var arg args.ContactArg
	util.Bind(r, &arg)

	users := contactService.SearchFriends(arg.Owenid)

	util.ApiOkList(writer, users, len(users))
}