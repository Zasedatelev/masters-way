/*


Testing FavoriteUserWayAPIService

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

func Test_openapi_FavoriteUserWayAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FavoriteUserWayAPIService CreateFavoriteUserWay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.FavoriteUserWayAPI.CreateFavoriteUserWay(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FavoriteUserWayAPIService DeleteFavoriteUserWay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userUuid string
		var wayUuid string

		httpRes, err := apiClient.FavoriteUserWayAPI.DeleteFavoriteUserWay(context.Background(), userUuid, wayUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
