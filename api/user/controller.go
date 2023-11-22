package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strconv"
	"time"
	"waf/config"
	"waf/domain/user"
	"waf/utils/api"
	jwt2 "waf/utils/jwt"
)

// Controller 用户控制器
type Controller struct {
	userService *user.Service
	config      *config.Config
}

// NewUserController 实例化用户控制器
func NewUserController(service *user.Service, config *config.Config) *Controller {
	return &Controller{
		userService: service,
		config:      config,
	}
}

// CreateUser godoc
// @Summary 创建用户
// @Tags Auth
// @Accept json
// @Product json
// @Param Authorization header string true "Authorization header"
// @Param CreateUserRequest body CreateUserRequest true "用户名、密码、重复密码"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/user [post]
func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	newUser := user.NewUser(req.Username, req.Password, req.Password2)
	err := c.userService.Create(newUser)
	if err != nil {
		api.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, CreateUserResponse{Username: req.Username})
}

// Login godoc
// @Summary 用户登录
// @Tags Auth
// @Accept json
// @Product json
// @Param LoginRequest body LoginRequest true "用户名、密码"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/login [post]
func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	currentUser, err := c.userService.CheckUserAndPassword(req.Username, req.Password)
	if err != nil {
		api.HandleError(g, err)
		return
	}
	decodedClaims := jwt2.VerifyToken(currentUser.Token, c.config.JwtSetting.SecretKey)
	if decodedClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.UserName,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp":      time.Now().Add(48 * time.Hour).Unix(),
			})
		token, err := jwt2.GenerateToken(jwtClaims, c.config.JwtSetting.SecretKey)
		if err != nil {
			api.HandleError(g, api.ErrJwtGenerate)
			return
		}
		currentUser.Token = token
		err = c.userService.Update(&currentUser)
		if err != nil {
			api.HandleError(g, err)
			return
		}
	}
	g.JSON(http.StatusOK, LoginResponse{
		Username: currentUser.UserName,
		UserId:   currentUser.ID,
		Token:    currentUser.Token,
	})
}

// ChangePassword godoc
// @Summary 修改密码
// @Tags Auth
// @Accept json
// @Product json
// @Param Authorization header string true "Authorization header"
// @Param ChangePasswordRequest body ChangePasswordRequest true "旧密码、新密码、重复密码"
// @Success 200 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/user [patch]
func (c *Controller) ChangePassword(g *gin.Context) {
	var req ChangePasswordRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	// 获取用户id
	userId := api.GetUserId(g)
	// 获取用户信息
	userInfo, err := c.userService.GetUserByID(userId)
	if err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	userInfo.Password2 = req.Password2
	err = c.userService.ChangePassword(&userInfo, req.OldPassword, req.Password)
	if err != nil {
		api.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Message: "success"})
}
