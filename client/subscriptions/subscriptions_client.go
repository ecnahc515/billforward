package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new subscriptions API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscriptions API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Create multiple subscriptions.

{"nickname":"Create multiple subscriptions","response":"createMultipleSubscriptionViaHelper.html","request":"createMultipleSubscriptionViaHelper.request.html"}
*/
func (a *Client) BatchCreateSubscriptions(params *BatchCreateSubscriptionsParams) (*BatchCreateSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchCreateSubscriptionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "batchCreateSubscriptions",
		Method:      "POST",
		PathPattern: "/subscriptions/batch",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &BatchCreateSubscriptionsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*BatchCreateSubscriptionsOK), nil
}

/*Retires the subscription specified by the subscription-ID parameter. Retiring a subscription causes it to cancel based on the specified retirement settings for the product.

{"nickname":"Cancel subscription","response":"deleteSubscription.html","request":"deleteSubscriptionRequest.html"}
*/
func (a *Client) CancelSubscription(params *CancelSubscriptionParams) (*CancelSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "cancelSubscription",
		Method:      "POST",
		PathPattern: "/subscriptions/{subscription-ID}/cancel",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &CancelSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelSubscriptionOK), nil
}

/*Create an aggregating subscription.

{"nickname":"Create aggregating subscription","response":"createAggregatingSubscription.html","request":"createAggregatingSubscription.request.html"}
*/
func (a *Client) CreateAggregatingSubscription(params *CreateAggregatingSubscriptionParams) (*CreateAggregatingSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAggregatingSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "createAggregatingSubscription",
		Method:      "POST",
		PathPattern: "/subscriptions/aggregating",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &CreateAggregatingSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAggregatingSubscriptionOK), nil
}

/*Create a new subscription.

{"nickname":"Create a new subscription","request":"createSubscriptionRequest.html","response":"createSubscriptionResponse.html"}
*/
func (a *Client) CreateSubscription(params *CreateSubscriptionParams) (*CreateSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "createSubscription",
		Method:      "POST",
		PathPattern: "/subscriptions",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &CreateSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSubscriptionOK), nil
}

/*Retrieves a collection of subscriptions, specified by the account-ID parameter. By default 10 values are returned. Records are returned in natural order.

{"nickname":"Retrieve by account","response":"getSubscriptionByAccount.html"}
*/
func (a *Client) GetSubscriptionByAccountID(params *GetSubscriptionByAccountIDParams) (*GetSubscriptionByAccountIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionByAccountIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "getSubscriptionByAccountID",
		Method:      "GET",
		PathPattern: "/subscriptions/account/{account-ID}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &GetSubscriptionByAccountIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionByAccountIDOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
