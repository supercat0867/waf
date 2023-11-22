package whitelist

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"waf/domain/whitelist"
	"waf/utils/api"
)

var ctx = context.Background()

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
// @Param AddIPToWhitelistRequest body AddIPToWhitelistRequest true "IP信息"
// @Success 201 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /whitelist [post]
func (c *Controller) AddIPToWhitelist(g *gin.Context) {
	var req AddIPToWhitelistRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	err := c.whitelistService.Add(req.IP, ctx)
	if err != nil {
		api.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, Response{Message: "success"})
}

// RemoveIPFromWhitelist godoc
// @Summary 将IP移除白名单
// @Tags IP黑白名单
// @Accept json
// @Produce json
// @Param AddIPToWhitelistRequest body AddIPToWhitelistRequest true "IP信息"
// @Success 200 {object} Response
// @Failure 400 {object} api.ErrorResponse
// @Router /whitelist [delete]
func (c *Controller) RemoveIPFromWhitelist(g *gin.Context) {
	var req AddIPToWhitelistRequest
	if err := g.ShouldBind(&req); err != nil {
		api.HandleError(g, api.ErrInvalidBody)
		return
	}
	err := c.whitelistService.Remove(req.IP, ctx)
	if err != nil {
		api.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Message: "success"})
}
