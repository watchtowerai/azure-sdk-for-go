//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_ListByProfile.json
func ExampleAFDOriginGroupsClient_ListByProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	pager := client.ListByProfile("<resource-group-name>",
		"<profile-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Get.json
func ExampleAFDOriginGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<origin-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AFDOriginGroupsClientGetResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Create.json
func ExampleAFDOriginGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<origin-group-name>",
		armcdn.AFDOriginGroup{
			Properties: &armcdn.AFDOriginGroupProperties{
				HealthProbeSettings: &armcdn.HealthProbeParameters{
					ProbeIntervalInSeconds: to.Int32Ptr(10),
					ProbePath:              to.StringPtr("<probe-path>"),
					ProbeProtocol:          armcdn.ProbeProtocolNotSet.ToPtr(),
					ProbeRequestType:       armcdn.HealthProbeRequestTypeNotSet.ToPtr(),
				},
				LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
					AdditionalLatencyInMilliseconds: to.Int32Ptr(1000),
					SampleSize:                      to.Int32Ptr(3),
					SuccessfulSamplesRequired:       to.Int32Ptr(3),
				},
				TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Int32Ptr(5),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AFDOriginGroupsClientCreateResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Update.json
func ExampleAFDOriginGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<origin-group-name>",
		armcdn.AFDOriginGroupUpdateParameters{
			Properties: &armcdn.AFDOriginGroupUpdatePropertiesParameters{
				HealthProbeSettings: &armcdn.HealthProbeParameters{
					ProbeIntervalInSeconds: to.Int32Ptr(10),
					ProbePath:              to.StringPtr("<probe-path>"),
					ProbeProtocol:          armcdn.ProbeProtocolNotSet.ToPtr(),
					ProbeRequestType:       armcdn.HealthProbeRequestTypeNotSet.ToPtr(),
				},
				LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
					AdditionalLatencyInMilliseconds: to.Int32Ptr(1000),
					SampleSize:                      to.Int32Ptr(3),
					SuccessfulSamplesRequired:       to.Int32Ptr(3),
				},
				TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Int32Ptr(5),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AFDOriginGroupsClientUpdateResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Delete.json
func ExampleAFDOriginGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<origin-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_ListResourceUsage.json
func ExampleAFDOriginGroupsClient_ListResourceUsage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	pager := client.ListResourceUsage("<resource-group-name>",
		"<profile-name>",
		"<origin-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}