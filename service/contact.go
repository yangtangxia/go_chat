package service

import (
	"errors"
	"../model"
	"time"
)

type ContactService struct {

}

func (service *ContactService) AddFriends(owenid, friends_id int64) error {
	if owenid == friends_id {
		return errors.New("不能添加自己为好友")
	}

	tem := model.Contact{}
	DbEngin.Where("ownerid = ?", owenid).And("friends_id = ?", friends_id).And("cate = ?", model.CATE_1).Get(&tem)
	if tem.Id > 0 {
		return errors.New("该用户已经被添加为好友")
	}

	// 开启事务
	session := DbEngin.NewSession()
	session.Begin()

	_, e2 := session.InsertOne(model.Contact{
		Ownerid: owenid,
		Friends_id: friends_id,
		Cate: model.CATE_1,
		Createat: time.Now(),
	})
	_, e3 := session.InsertOne(model.Contact{
		Ownerid: friends_id,
		Friends_id: owenid,
		Cate: model.CATE_1,
		Createat: time.Now(),
	})

	if e2 == nil && e3 == nil {
		session.Commit()
		return nil
	} else {
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

func (service *ContactService) SearchFriends(owenid int64) ([]model.User) {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)

	DbEngin.Where("ownerid = ? and cate = ?", owenid, model.CATE_1).Find(&conconts)
	for _, v := range conconts{
		objIds = append(objIds, v.Friends_id)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngin.In("id", objIds).Find(&coms)
	return coms
}