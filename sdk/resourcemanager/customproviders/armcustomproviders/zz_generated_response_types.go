//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomproviders

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AssociationsClientCreateOrUpdatePollerResponse contains the response from method AssociationsClient.CreateOrUpdate.
type AssociationsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AssociationsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AssociationsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AssociationsClientCreateOrUpdateResponse, error) {
	respType := AssociationsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Association)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AssociationsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *AssociationsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *AssociationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AssociationsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AssociationsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AssociationsClientCreateOrUpdateResponse contains the response from method AssociationsClient.CreateOrUpdate.
type AssociationsClientCreateOrUpdateResponse struct {
	AssociationsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssociationsClientCreateOrUpdateResult contains the result from method AssociationsClient.CreateOrUpdate.
type AssociationsClientCreateOrUpdateResult struct {
	Association
}

// AssociationsClientDeletePollerResponse contains the response from method AssociationsClient.Delete.
type AssociationsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AssociationsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AssociationsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AssociationsClientDeleteResponse, error) {
	respType := AssociationsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AssociationsClientDeletePollerResponse from the provided client and resume token.
func (l *AssociationsClientDeletePollerResponse) Resume(ctx context.Context, client *AssociationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AssociationsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AssociationsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AssociationsClientDeleteResponse contains the response from method AssociationsClient.Delete.
type AssociationsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssociationsClientGetResponse contains the response from method AssociationsClient.Get.
type AssociationsClientGetResponse struct {
	AssociationsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssociationsClientGetResult contains the result from method AssociationsClient.Get.
type AssociationsClientGetResult struct {
	Association
}

// AssociationsClientListAllResponse contains the response from method AssociationsClient.ListAll.
type AssociationsClientListAllResponse struct {
	AssociationsClientListAllResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssociationsClientListAllResult contains the result from method AssociationsClient.ListAll.
type AssociationsClientListAllResult struct {
	AssociationsList
}

// CustomResourceProviderClientCreateOrUpdatePollerResponse contains the response from method CustomResourceProviderClient.CreateOrUpdate.
type CustomResourceProviderClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CustomResourceProviderClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CustomResourceProviderClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CustomResourceProviderClientCreateOrUpdateResponse, error) {
	respType := CustomResourceProviderClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.CustomRPManifest)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CustomResourceProviderClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *CustomResourceProviderClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *CustomResourceProviderClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CustomResourceProviderClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CustomResourceProviderClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CustomResourceProviderClientCreateOrUpdateResponse contains the response from method CustomResourceProviderClient.CreateOrUpdate.
type CustomResourceProviderClientCreateOrUpdateResponse struct {
	CustomResourceProviderClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomResourceProviderClientCreateOrUpdateResult contains the result from method CustomResourceProviderClient.CreateOrUpdate.
type CustomResourceProviderClientCreateOrUpdateResult struct {
	CustomRPManifest
}

// CustomResourceProviderClientDeletePollerResponse contains the response from method CustomResourceProviderClient.Delete.
type CustomResourceProviderClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CustomResourceProviderClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CustomResourceProviderClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CustomResourceProviderClientDeleteResponse, error) {
	respType := CustomResourceProviderClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CustomResourceProviderClientDeletePollerResponse from the provided client and resume token.
func (l *CustomResourceProviderClientDeletePollerResponse) Resume(ctx context.Context, client *CustomResourceProviderClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CustomResourceProviderClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &CustomResourceProviderClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CustomResourceProviderClientDeleteResponse contains the response from method CustomResourceProviderClient.Delete.
type CustomResourceProviderClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomResourceProviderClientGetResponse contains the response from method CustomResourceProviderClient.Get.
type CustomResourceProviderClientGetResponse struct {
	CustomResourceProviderClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomResourceProviderClientGetResult contains the result from method CustomResourceProviderClient.Get.
type CustomResourceProviderClientGetResult struct {
	CustomRPManifest
}

// CustomResourceProviderClientListByResourceGroupResponse contains the response from method CustomResourceProviderClient.ListByResourceGroup.
type CustomResourceProviderClientListByResourceGroupResponse struct {
	CustomResourceProviderClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomResourceProviderClientListByResourceGroupResult contains the result from method CustomResourceProviderClient.ListByResourceGroup.
type CustomResourceProviderClientListByResourceGroupResult struct {
	ListByCustomRPManifest
}

// CustomResourceProviderClientListBySubscriptionResponse contains the response from method CustomResourceProviderClient.ListBySubscription.
type CustomResourceProviderClientListBySubscriptionResponse struct {
	CustomResourceProviderClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomResourceProviderClientListBySubscriptionResult contains the result from method CustomResourceProviderClient.ListBySubscription.
type CustomResourceProviderClientListBySubscriptionResult struct {
	ListByCustomRPManifest
}

// CustomResourceProviderClientUpdateResponse contains the response from method CustomResourceProviderClient.Update.
type CustomResourceProviderClientUpdateResponse struct {
	CustomResourceProviderClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomResourceProviderClientUpdateResult contains the result from method CustomResourceProviderClient.Update.
type CustomResourceProviderClientUpdateResult struct {
	CustomRPManifest
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	ResourceProviderOperationList
}