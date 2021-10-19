package args

type ContactArg struct {
	PageArg
	Owenid int64 `json:"owenid" form:owenid`
	FriendId int64 `json:"friend_id" form:"friends_id"`
}

