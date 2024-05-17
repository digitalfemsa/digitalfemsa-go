/*
Femsa API

Testing CompaniesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package digitalfemsa

import (
	"context"
	openapiclient "github.com/digitalfemsa/digitalfemsa-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_digitalfemsa_CompaniesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CompaniesAPIService GetCompanies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CompaniesAPI.GetCompanies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CompaniesAPIService GetCompany", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CompaniesAPI.GetCompany(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
