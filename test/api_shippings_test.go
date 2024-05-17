/*
Femsa API

Testing ShippingsAPIService

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

func Test_digitalfemsa_ShippingsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShippingsAPIService OrdersCreateShipping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ShippingsAPI.OrdersCreateShipping(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingsAPIService OrdersDeleteShipping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var shippingId string

		resp, httpRes, err := apiClient.ShippingsAPI.OrdersDeleteShipping(context.Background(), id, shippingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingsAPIService OrdersUpdateShipping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var shippingId string

		resp, httpRes, err := apiClient.ShippingsAPI.OrdersUpdateShipping(context.Background(), id, shippingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
