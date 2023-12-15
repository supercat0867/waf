package whitelist

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"waf/domain/whitelist"
	"waf/utils/api"
	"waf/utils/validator"
)

// Controller 白名单控制器
type Controller struct {
	whitelistService *whitelist.Service
}

// NewWhitelistController 实例化白名单控制器
func NewWhitelistController(service *whitelist.Service) *Controller {
	return &Controller{
		whitelistService: service,
	}
}

// AddIPToWhitelist godoc
// @Summary 添加IP至白名单
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param AddIPToWhitelistRequest body AddIPToWhitelistRequest true "IP信息"
// @Success 201 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/whitelist [post]
func (c *Controller) AddIPToWhitelist(g *gin.Context) {
	var req AddIPToWhitelistRequest
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

	err := c.whitelistService.Add(req.IP, expiryDuration)
	if err != nil {
		api.HandleInternalServerError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Status: 200, Message: "success"})
}

// RemoveIPFromWhitelist godoc
// @Summary 将IP移除白名单
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param RemoveIPTiWhitelistRequest body RemoveIPTiWhitelistRequest true "IP信息"
// @Success 200 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /waf/whitelist [delete]
func (c *Controller) RemoveIPFromWhitelist(g *gin.Context) {
	var req RemoveIPTiWhitelistRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	if !validator.ValidateIPorCIDR(req.IP) {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	err := c.whitelistService.Remove(req.IP)
	if err != nil {
		api.HandleInternalServerError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Status: 200, Message: "success"})
}

// GetIps godoc
// @Summary 获取白名单列表
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Success 200 {object} IpListResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /waf/whitelist [get]
func (c *Controller) GetIps(g *gin.Context) {
	ips, err := c.whitelistService.Get()
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
