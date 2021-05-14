package models

//User user message
type User struct {
	UserName  string `json:"userName"` //用户名
	Passworld string `json:"password"` //密码
}

//LoginResponse login response
type LoginResponse struct {
	Token   string `json:"token"`   //登录token
	Message string `json:"message"` //反馈信息
}

//UserInfo user info
type UserInfo struct {
	Name   string   `json:"name"`
	UserID string   `json:"user_id"`
	Access []string `json:"access"`
	Token  string   `json:"token"`
	Avator string   `json:"avator"`
}
