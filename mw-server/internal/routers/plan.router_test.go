package routers

import (
	"context"
	openapiGeneral "mwserver/apiAutogenerated/general"
	"net/http"

	"mwserver/internal/auth"
	"mwserver/internal/config"
	"mwserver/internal/openapi"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlan(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91"

	t.Run("should create Plan and return it successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		request := openapiGeneral.SchemasCreatePlanPayload{
			DayReportUuid: "25ceb64e-7a57-4ce0-a4fd-45982d9fce38",
			Description:   "Description",
			IsDone:        false,
			OwnerUuid:     "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91",
			Time:          0,
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		plan, response, err := generalApi.PlanAPI.CreatePlan(ctx).Request(request).Execute()
		if err != nil {
			t.Fatalf("Failed to create plan: %v", err)
		}

		expectedData := &openapiGeneral.SchemasPlanPopulatedResponse{
			DayReportUuid: "25ceb64e-7a57-4ce0-a4fd-45982d9fce38",
			Description:   "Description",
			IsDone:        false,
			OwnerName:     "Alice Johnson",
			OwnerUuid:     "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91",
			Tags:          []openapiGeneral.SchemasJobTagResponse{},
			Time:          0,
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.DayReportUuid, plan.DayReportUuid)
		assert.Equal(t, expectedData.Description, plan.Description)
		assert.Equal(t, expectedData.IsDone, plan.IsDone)
		assert.Equal(t, expectedData.OwnerName, plan.OwnerName)
		assert.Equal(t, expectedData.OwnerUuid, plan.OwnerUuid)
		assert.Equal(t, expectedData.Tags, plan.Tags)
		assert.Equal(t, expectedData.Time, plan.Time)
	})
}

func TestUpdatePlan(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91"
	planID := "18cbbee6-5071-4608-b349-ffad514711cb"

	t.Run("should update Plan and return it successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		newDescription := "New Description"
		newIsDoneStatus := true
		newTime := int32(12)

		request := openapiGeneral.SchemasUpdatePlanPayload{
			Description: &newDescription,
			IsDone:      &newIsDoneStatus,
			Time:        &newTime,
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		plan, response, err := generalApi.PlanAPI.UpdatePlan(ctx, planID).Request(request).Execute()

		expectedData := &openapiGeneral.SchemasPlanPopulatedResponse{
			Uuid:        planID,
			Description: newDescription,
			IsDone:      newIsDoneStatus,
			Time:        newTime,
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Uuid, plan.Uuid)
		assert.Equal(t, expectedData.Description, plan.Description)
		assert.Equal(t, expectedData.IsDone, plan.IsDone)
		assert.Equal(t, expectedData.Time, plan.Time)
	})
}

func TestDeletePlanById(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91"
	planID := "18cbbee6-5071-4608-b349-ffad514711cb"

	t.Run("should delete Plan successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		response, err := generalApi.PlanAPI.DeletePlan(ctx, planID).Execute()

		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		wayID := "1d922e8a-5d58-4b82-9a3d-83e2e73b3f91"
		reports, response, err := generalApi.DayReportAPI.GetDayReports(ctx, wayID).Execute()

		isExists := false
		for _, report := range reports.DayReports {
			for i := range report.Plans {
				if report.Plans[i].Uuid == planID {
					isExists = true
					break
				}
			}
		}

		if isExists {
			t.Fatalf("Plan %s wasn't removed", planID)
		}
	})
}