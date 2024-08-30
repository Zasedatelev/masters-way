package controllers

import (
	"mwchat/internal/services"
	"mwchat/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DevController struct {
	devService *services.DevService
}

func NewDevController(devService *services.DevService) *DevController {
	return &DevController{devService}
}

// Reset db
// @Summary resets db
// @Description resets db
// @Tags dev
// @Success 200
// @Router /dev/reset-db [get]
func (dc *DevController) ResetDB(ctx *gin.Context) {
	err := dc.devService.ResetDB(ctx)
	utils.HandleErrorGin(ctx, err)

	ctx.Status(http.StatusOK)
}
