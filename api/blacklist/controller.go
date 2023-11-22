package blacklist

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"waf/domain/blacklist"
	"waf/utils/api"
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
// @Param AddIPToBlacklistRequest body AddIPToBlacklistRequest true "IP信息"
// @Success 201 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /blacklist [post]
func (c *Controller) AddIPToBlacklist(g *gin.Context) {
	var req AddIPToBlacklistRequest
	if err := g.ShouldBind(&req); err != nil {
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
// @Param AddIPToBlacklistRequest body AddIPToBlacklistRequest true "IP信息"
// @Success 200 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /blacklist [delete]
func (c *Controller) RemoveIPFromBlacklist(g *gin.Context) {
	var req AddIPToBlacklistRequest
	if err := g.ShouldBind(&req); err != nil {
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
