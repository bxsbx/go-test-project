package authenticate

type UserInfo struct {
	UserId   string
	UserName string
	UserType int
}

func GetUserInfoFromToken(token string) UserInfo {
	return UserInfo{
		UserId:   "2323",
		UserName: "vsav",
		UserType: 1,
	}
}
