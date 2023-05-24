/*
Seagate Management Controller (MC) OpenAPI

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_client_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService CopyVolumeDestinationPoolNameSourceGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var destinationPoolOption string
		var nameOption string
		var sourceOption string

		resp, httpRes, err := apiClient.DefaultApi.CopyVolumeDestinationPoolNameSourceGet(context.Background(), destinationPoolOption, nameOption, sourceOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateSnapshotsVolumesNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var volumesOption string
		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.CreateSnapshotsVolumesNamesGet(context.Background(), volumesOption, namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateVolumePoolSizeTierAffinityNameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolOption string
		var sizeOption string
		var tierAffinityOption string
		var nameOption string

		resp, httpRes, err := apiClient.DefaultApi.CreateVolumePoolSizeTierAffinityNameGet(context.Background(), poolOption, sizeOption, tierAffinityOption, nameOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteHostsNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.DeleteHostsNamesGet(context.Background(), namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteInitiatorNicknameNameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOption string

		resp, httpRes, err := apiClient.DefaultApi.DeleteInitiatorNicknameNameGet(context.Background(), nameOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteSnapshotNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.DeleteSnapshotNamesGet(context.Background(), namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteVolumesNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.DeleteVolumesNamesGet(context.Background(), namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ExpandVolumeSizeNameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sizeOption string
		var nameOption string

		resp, httpRes, err := apiClient.DefaultApi.ExpandVolumeSizeNameGet(context.Background(), sizeOption, nameOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService LoginGetByHash", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var loginHash string

		resp, httpRes, err := apiClient.DefaultApi.LoginGetByHash(context.Background(), loginHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService MapVolumeAccessLunInitiatorNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accessOption string
		var lunOption string
		var initiatorOption string
		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.MapVolumeAccessLunInitiatorNamesGet(context.Background(), accessOption, lunOption, initiatorOption, namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService SchemaGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var schemaId string

		resp, httpRes, err := apiClient.DefaultApi.SchemaGet(context.Background(), schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService SetInitiatorIdNicknameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOption string
		var nicknameOption string

		resp, httpRes, err := apiClient.DefaultApi.SetInitiatorIdNicknameGet(context.Background(), idOption, nicknameOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowAdvancedSettingsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowAdvancedSettingsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowCacheParametersGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowCacheParametersGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowCertificateGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowCertificateGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowControllersGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowControllersGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowDisksGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowDisksGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowEnclosuresGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowEnclosuresGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowFansGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowFansGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowHostGroupsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowHostGroupsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowInitiatorNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.ShowInitiatorNamesGet(context.Background(), namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowInitiatorsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowInitiatorsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowMapsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowMapsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowMapsInitiatorGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowMapsInitiatorGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowMapsNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.ShowMapsNamesGet(context.Background(), namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowPoolsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowPoolsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowPowerSuppliesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowPowerSuppliesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowSnapshotsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowSnapshotsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowSnapshotsPatternGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var patternOption string

		resp, httpRes, err := apiClient.DefaultApi.ShowSnapshotsPatternGet(context.Background(), patternOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowSnapshotsVolumeGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var volumeOption string

		resp, httpRes, err := apiClient.DefaultApi.ShowSnapshotsVolumeGet(context.Background(), volumeOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowSystemGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowSystemGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowVersionsDetailGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowVersionsDetailGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowVersionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.ShowVersionsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ShowVolumesNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.ShowVolumesNamesGet(context.Background(), namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService UnmapVolumeInitiatorNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var initiatorOption string
		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.UnmapVolumeInitiatorNamesGet(context.Background(), initiatorOption, namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService UnmapVolumeNamesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namesOption string

		resp, httpRes, err := apiClient.DefaultApi.UnmapVolumeNamesGet(context.Background(), namesOption).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
