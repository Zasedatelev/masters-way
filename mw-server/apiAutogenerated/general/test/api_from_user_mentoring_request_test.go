/*


Testing FromUserMentoringRequestAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_FromUserMentoringRequestAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FromUserMentoringRequestAPIService CreateFromUserMentoringRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.FromUserMentoringRequestAPI.CreateFromUserMentoringRequest(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FromUserMentoringRequestAPIService DeleteFromUserMentoringRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userUuid string
		var wayUuid string

		httpRes, err := apiClient.FromUserMentoringRequestAPI.DeleteFromUserMentoringRequest(context.Background(), userUuid, wayUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
