package routers

import (
	"bytes"
	"encoding/json"
	"mwcustdev/internal/controllers"
	"mwcustdev/internal/schemas"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProposal(t *testing.T) {
	router := gin.Default()
	c := controllers.CustdevController{}
	router.POST("", c.CreateProposal)
	userUuid := "f4cbdb32-0fc0-4704-a6d9-9cf357e5cd3e"
	proposal := "New proposals"

	newProposal := schemas.CreateProposalPayload{
		UserUuid: &userUuid,
		Proposal: &proposal,
		CreateAt: "2018-11-12 01:02:03.123456789",
	}

	jsonValue, _ := json.Marshal(newProposal)
	req, _ := http.NewRequest("POST", "/proposal/add", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}
