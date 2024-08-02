package request

type UserCreateDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateDTO struct {
	Id       uint64 `json:"id"`       // ID
	Avatar   string `json:"avatar"`   // 头像
	Nickname string `json:"nickname"` // 昵称
}

type UserUpdatePasswordDTO struct {
	OldPassword string `json:"oldPassword"` // 原密码
	Password    string `json:"password"`    // 密码
}
