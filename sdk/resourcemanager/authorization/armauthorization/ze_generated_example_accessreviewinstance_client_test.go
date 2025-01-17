//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/StopAccessReviewInstance.json
func ExampleAccessReviewInstanceClient_Stop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewInstanceClient("fa73e90b-5bf1-45fd-a182-35ce5fc0674d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Stop(ctx,
		"fa73e90b-5bf1-45fd-a182-35ce5fc0674d",
		"d9b9e056-7004-470b-bf21-1635e98487da",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/AccessReviewInstanceResetDecisions.json
func ExampleAccessReviewInstanceClient_ResetDecisions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewInstanceClient("fa73e90b-5bf1-45fd-a182-35ce5fc0674d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ResetDecisions(ctx,
		"fa73e90b-5bf1-45fd-a182-35ce5fc0674d",
		"d9b9e056-7004-470b-bf21-1635e98487da",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/AccessReviewInstanceApplyDecisions.json
func ExampleAccessReviewInstanceClient_ApplyDecisions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewInstanceClient("fa73e90b-5bf1-45fd-a182-35ce5fc0674d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ApplyDecisions(ctx,
		"fa73e90b-5bf1-45fd-a182-35ce5fc0674d",
		"d9b9e056-7004-470b-bf21-1635e98487da",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/AccessReviewInstanceSendReminders.json
func ExampleAccessReviewInstanceClient_SendReminders() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewInstanceClient("fa73e90b-5bf1-45fd-a182-35ce5fc0674d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.SendReminders(ctx,
		"fa73e90b-5bf1-45fd-a182-35ce5fc0674d",
		"d9b9e056-7004-470b-bf21-1635e98487da",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/AccessReviewInstanceAcceptRecommendations.json
func ExampleAccessReviewInstanceClient_AcceptRecommendations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewInstanceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.AcceptRecommendations(ctx,
		"488a6d0e-0a63-4946-86e3-1f5bbc934661",
		"d9b9e056-7004-470b-bf21-1635e98487da",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
