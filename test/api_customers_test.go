/*
Femsa API

Testing CustomersAPIService

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

func Test_digitalfemsa_CustomersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomersAPIService CreateCustomer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService CreateCustomerFiscalEntities", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CustomersAPI.CreateCustomerFiscalEntities(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService DeleteCustomerById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CustomersAPI.DeleteCustomerById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService GetCustomerById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CustomersAPI.GetCustomerById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService GetCustomers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomersAPI.GetCustomers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService UpdateCustomer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CustomersAPI.UpdateCustomer(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService UpdateCustomerFiscalEntities", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var fiscalEntitiesId string

		resp, httpRes, err := apiClient.CustomersAPI.UpdateCustomerFiscalEntities(context.Background(), id, fiscalEntitiesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
