package user

// CreateUserRequest 创建用户请求结构体
type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

// CreateUserResponse 创建用户成功响应
type CreateUserResponse struct {
	Username string `json:"username"`
}

// LoginRequest 用户登录请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 用户登录成功响应结构体
type LoginResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	Token    string `json:"token"`
}

// ChangePasswordRequest 用户修改密码请求结构体
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
	Password2   string `json:"password2"`
}

// Response 响应
type Response struct {
	Message string `json:"msg"`
}
