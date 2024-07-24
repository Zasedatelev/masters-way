/*


Testing RoomAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_RoomAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoomAPIService AddUserToRoom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomId string
		var userId string

		resp, httpRes, err := apiClient.RoomAPI.AddUserToRoom(context.Background(), roomId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService CreateRoom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomType string

		resp, httpRes, err := apiClient.RoomAPI.CreateRoom(context.Background(), roomType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService DeleteUserToGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomId string
		var userId string

		resp, httpRes, err := apiClient.RoomAPI.DeleteUserToGroup(context.Background(), roomId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService GetChatPreview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RoomAPI.GetChatPreview(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService GetRoomById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomId string

		resp, httpRes, err := apiClient.RoomAPI.GetRoomById(context.Background(), roomId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService GetRooms", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomType string

		resp, httpRes, err := apiClient.RoomAPI.GetRooms(context.Background(), roomType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService MakeMessageInRoom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomId string

		resp, httpRes, err := apiClient.RoomAPI.MakeMessageInRoom(context.Background(), roomId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService UpdateRoom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomId string

		resp, httpRes, err := apiClient.RoomAPI.UpdateRoom(context.Background(), roomId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
