package blacklist

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"waf/domain/blacklist"
	"waf/utils/api"
	"waf/utils/validator"
)

var ctx = context.Background()

// Controller 黑名单控制器
type Controller struct {
	blacklistService *blacklist.Service
}

// NewBlacklistController 实例化黑名单控制器
func NewBlacklistController(service *blacklist.Service) *Controller {
	return &Controller{
		blacklistService: service,
	}
}

// AddIPToBlacklist godoc
// @Summary 添加IP至黑名单
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param AddIPToBlacklistRequest body AddIPToBlacklistRequest true "IP信息"
// @Success 201 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/blacklist [post]
func (c *Controller) AddIPToBlacklist(g *gin.Context) {
	var req AddIPToBlacklistRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	if !validator.ValidateIPorCIDR(req.IP) {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	err := c.blacklistService.Add(req.IP, ctx)
	if err != nil {
		api.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, Response{Message: "success"})
}

// RemoveIPFromBlacklist godoc
// @Summary 将IP移除黑名单
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param AddIPToBlacklistRequest body AddIPToBlacklistRequest true "IP信息"
// @Success 200 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/blacklist [delete]
func (c *Controller) RemoveIPFromBlacklist(g *gin.Context) {
	var req AddIPToBlacklistRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	if !validator.ValidateIPorCIDR(req.IP) {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	err := c.blacklistService.Remove(req.IP, ctx)
	if err != nil {
		api.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Message: "success"})
}

// GetIps godoc
// @Summary 获取黑名单列表
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Success 200 {object} IpResponse
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/blacklist [get]
func (c *Controller) GetIps(g *gin.Context) {
	ips, err := c.blacklistService.Get(ctx)
	if err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	g.JSON(http.StatusOK, IpResponse{Data: ips})
}
