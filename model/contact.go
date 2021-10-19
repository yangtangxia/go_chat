package model

import "time"

const (
	CATE_1 = 1
	CATE_2 = 2
)

type Contact struct {
	Id  int64 `xorm:"pk autoincr bigint(20) comment('主键id')" form:"id" json:"id"`
	Ownerid int64 `xorm:"bigint(20) comment('自己id')" form:"owner_id"`
	Friends_id int64 `xorm:"bigint(20) comment('好友id')" form:"friends_id"`
	Cate int8 `xorm:"int(2) comment('栏目： 1好友， 2群')" form:"cate"`
	Name string `xorm:"varchar(20) comment('名称')" form:"name" json:"name"`
	Icon string `xorm:"varchar(255) comment('图片')" form:"icon" josn:"icon"`
	Memo string `xorm:"varchar(255) comment('简介')" form:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime comment('创建时间')" form:"createat" json:"createat"`
	Updateat time.Time `xorm:"datetime comment('修改时间')" form:"updateat" json:"updateat"`
}
//ENGINE=InnoDB charset=utf8 comment="好友关系表"
