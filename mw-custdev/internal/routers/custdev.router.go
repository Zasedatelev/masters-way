package routers

import (
	"mwcustdev/internal/auth"
	"mwcustdev/internal/controllers"

	"github.com/gin-gonic/gin"
)

type CustdevRouter struct {
	custdevController *controllers.CustdevController
}

func newCustdevController(proposalController *controllers.CustdevController) *CustdevRouter {
	return &CustdevRouter{proposalController}
}

func (cr *CustdevRouter) SetCustdevRouter(rg *gin.RouterGroup) {
	proposals := rg.Group("/proposals", auth.AuthMiddleware())
	proposals.POST("/proposal/add", cr.custdevController.CreateProposal)
}
