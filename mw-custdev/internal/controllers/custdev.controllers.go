package controllers

import (
	"mwcustdev/internal/auth"
	"mwcustdev/internal/schemas"
	"mwcustdev/internal/services"
	"mwcustdev/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustdevController struct {
	custdevService *services.CustdevService
}

func NewCustDevController(custdevService *services.CustdevService) *CustdevController {
	return &CustdevController{custdevService: custdevService}
}

// @Summary Create proposal for user
// @Description
// @Tags proposal
// @ID create-proposal
// @Accept  json
// @Produce  json
// @Param request body schemas.CreateProposalPayload true "query params"
// @Success 200 {object} schemas.CreateProposalsRow
// @Router /proposa/add [post]
func (cs *CustdevController) CreateProposal(ctx *gin.Context) {
	var payload *schemas.CreateProposalPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creatorIDRaw, _ := ctx.Get(auth.ContextKeyUserID)
	creatorUUID := uuid.MustParse(creatorIDRaw.(string))

	params := &services.CreatreProposalsParams{
		UserUuid: creatorUUID,
		Proposal: payload.CreateAt,
		CreateAt: payload.CreateAt,
	}

	err := cs.custdevService.CreateProposal(ctx, params)

	utils.HandleErrorGin(ctx, err)

	ctx.IndentedJSON(http.StatusOK, "status: success")

}
