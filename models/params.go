package models

// 定义请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Repassword string `json:"re_password"`
}
