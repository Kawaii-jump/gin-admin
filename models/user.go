package models

//User user message
type User struct {
	UserName  string `json:"username"` //用户名
	Passworld string `json:"password"` //密码
}

//LoginResponse login response
type LoginResponse struct {
	Token   string `json:"token"`   //登录token
	Message string `json:"message"` //反馈信息
}
