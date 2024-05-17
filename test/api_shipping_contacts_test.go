/*
Femsa API

Testing ShippingContactsAPIService

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

func Test_digitalfemsa_ShippingContactsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShippingContactsAPIService CreateCustomerShippingContacts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ShippingContactsAPI.CreateCustomerShippingContacts(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingContactsAPIService DeleteCustomerShippingContacts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var shippingContactsId string

		resp, httpRes, err := apiClient.ShippingContactsAPI.DeleteCustomerShippingContacts(context.Background(), id, shippingContactsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShippingContactsAPIService UpdateCustomerShippingContacts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var shippingContactsId string

		resp, httpRes, err := apiClient.ShippingContactsAPI.UpdateCustomerShippingContacts(context.Background(), id, shippingContactsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
