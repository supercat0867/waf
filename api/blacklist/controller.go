package blacklist

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"waf/domain/blacklist"
	"waf/utils/api"
	"waf/utils/validator"
)

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
// @Success 200 {object} Response
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

	// 将秒数转换为 time.Duration
	expiryDuration := time.Duration(req.Expiry) * time.Second

	err := c.blacklistService.Add(req.IP, expiryDuration)
	if err != nil {
		api.HandleInternalServerError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Status: 200, Message: "success"})
}

// RemoveIPFromBlacklist godoc
// @Summary 将IP移除黑名单
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param RemoveIPTiBlacklistRequest body RemoveIPTiBlacklistRequest true "IP信息"
// @Success 200 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/blacklist [delete]
func (c *Controller) RemoveIPFromBlacklist(g *gin.Context) {
	var req RemoveIPTiBlacklistRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	if !validator.ValidateIPorCIDR(req.IP) {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	err := c.blacklistService.Remove(req.IP)
	if err != nil {
		api.HandleInternalServerError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Status: 200, Message: "success"})
}

// GetIps godoc
// @Summary 获取黑名单列表
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Success 200 {object} IpListResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /waf/blacklist [get]
func (c *Controller) GetIps(g *gin.Context) {
	ips, err := c.blacklistService.Get()
	if err != nil {
		api.HandleInternalServerError(g, err)
		return
	}
	resp := IpListResponse{
		Status:  200,
		Count:   len(ips),
		Message: "success",
		Data:    ips,
	}

	g.JSON(http.StatusOK, resp)
}
