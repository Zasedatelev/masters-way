package controllers

import (
	"mwcustdev/internal/services"
)

type Controller struct {
	CustdevController *CustdevController
}

func NewController(services *services.Service) *Controller {
	return &Controller{
		CustdevController: NewCustDevController(services.CustdevService),
	}
}
