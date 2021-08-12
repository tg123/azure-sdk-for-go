// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// EligibleChildResourcesListResultPager provides iteration over EligibleChildResourcesListResult pages.
type EligibleChildResourcesListResultPager interface {
	azcore.Pager

	// PageResponse returns the current EligibleChildResourcesListResultResponse.
	PageResponse() EligibleChildResourcesListResultResponse
}

type eligibleChildResourcesListResultCreateRequest func(context.Context) (*azcore.Request, error)

type eligibleChildResourcesListResultHandleError func(*azcore.Response) error

type eligibleChildResourcesListResultHandleResponse func(*azcore.Response) (EligibleChildResourcesListResultResponse, error)

type eligibleChildResourcesListResultAdvancePage func(context.Context, EligibleChildResourcesListResultResponse) (*azcore.Request, error)

type eligibleChildResourcesListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester eligibleChildResourcesListResultCreateRequest
	// callback for handling response errors
	errorer eligibleChildResourcesListResultHandleError
	// callback for handling the HTTP response
	responder eligibleChildResourcesListResultHandleResponse
	// callback for advancing to the next page
	advancer eligibleChildResourcesListResultAdvancePage
	// contains the current response
	current EligibleChildResourcesListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *eligibleChildResourcesListResultPager) Err() error {
	return p.err
}

func (p *eligibleChildResourcesListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EligibleChildResourcesListResult.NextLink == nil || len(*p.current.EligibleChildResourcesListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *eligibleChildResourcesListResultPager) PageResponse() EligibleChildResourcesListResultResponse {
	return p.current
}

// RoleAssignmentListResultPager provides iteration over RoleAssignmentListResult pages.
type RoleAssignmentListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleAssignmentListResultResponse.
	PageResponse() RoleAssignmentListResultResponse
}

type roleAssignmentListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleAssignmentListResultHandleError func(*azcore.Response) error

type roleAssignmentListResultHandleResponse func(*azcore.Response) (RoleAssignmentListResultResponse, error)

type roleAssignmentListResultAdvancePage func(context.Context, RoleAssignmentListResultResponse) (*azcore.Request, error)

type roleAssignmentListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleAssignmentListResultCreateRequest
	// callback for handling response errors
	errorer roleAssignmentListResultHandleError
	// callback for handling the HTTP response
	responder roleAssignmentListResultHandleResponse
	// callback for advancing to the next page
	advancer roleAssignmentListResultAdvancePage
	// contains the current response
	current RoleAssignmentListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleAssignmentListResultPager) Err() error {
	return p.err
}

func (p *roleAssignmentListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentListResult.NextLink == nil || len(*p.current.RoleAssignmentListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleAssignmentListResultPager) PageResponse() RoleAssignmentListResultResponse {
	return p.current
}

// RoleAssignmentScheduleInstanceListResultPager provides iteration over RoleAssignmentScheduleInstanceListResult pages.
type RoleAssignmentScheduleInstanceListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleAssignmentScheduleInstanceListResultResponse.
	PageResponse() RoleAssignmentScheduleInstanceListResultResponse
}

type roleAssignmentScheduleInstanceListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleAssignmentScheduleInstanceListResultHandleError func(*azcore.Response) error

type roleAssignmentScheduleInstanceListResultHandleResponse func(*azcore.Response) (RoleAssignmentScheduleInstanceListResultResponse, error)

type roleAssignmentScheduleInstanceListResultAdvancePage func(context.Context, RoleAssignmentScheduleInstanceListResultResponse) (*azcore.Request, error)

type roleAssignmentScheduleInstanceListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleAssignmentScheduleInstanceListResultCreateRequest
	// callback for handling response errors
	errorer roleAssignmentScheduleInstanceListResultHandleError
	// callback for handling the HTTP response
	responder roleAssignmentScheduleInstanceListResultHandleResponse
	// callback for advancing to the next page
	advancer roleAssignmentScheduleInstanceListResultAdvancePage
	// contains the current response
	current RoleAssignmentScheduleInstanceListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleAssignmentScheduleInstanceListResultPager) Err() error {
	return p.err
}

func (p *roleAssignmentScheduleInstanceListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentScheduleInstanceListResult.NextLink == nil || len(*p.current.RoleAssignmentScheduleInstanceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleAssignmentScheduleInstanceListResultPager) PageResponse() RoleAssignmentScheduleInstanceListResultResponse {
	return p.current
}

// RoleAssignmentScheduleListResultPager provides iteration over RoleAssignmentScheduleListResult pages.
type RoleAssignmentScheduleListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleAssignmentScheduleListResultResponse.
	PageResponse() RoleAssignmentScheduleListResultResponse
}

type roleAssignmentScheduleListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleAssignmentScheduleListResultHandleError func(*azcore.Response) error

type roleAssignmentScheduleListResultHandleResponse func(*azcore.Response) (RoleAssignmentScheduleListResultResponse, error)

type roleAssignmentScheduleListResultAdvancePage func(context.Context, RoleAssignmentScheduleListResultResponse) (*azcore.Request, error)

type roleAssignmentScheduleListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleAssignmentScheduleListResultCreateRequest
	// callback for handling response errors
	errorer roleAssignmentScheduleListResultHandleError
	// callback for handling the HTTP response
	responder roleAssignmentScheduleListResultHandleResponse
	// callback for advancing to the next page
	advancer roleAssignmentScheduleListResultAdvancePage
	// contains the current response
	current RoleAssignmentScheduleListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleAssignmentScheduleListResultPager) Err() error {
	return p.err
}

func (p *roleAssignmentScheduleListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentScheduleListResult.NextLink == nil || len(*p.current.RoleAssignmentScheduleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleAssignmentScheduleListResultPager) PageResponse() RoleAssignmentScheduleListResultResponse {
	return p.current
}

// RoleAssignmentScheduleRequestListResultPager provides iteration over RoleAssignmentScheduleRequestListResult pages.
type RoleAssignmentScheduleRequestListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleAssignmentScheduleRequestListResultResponse.
	PageResponse() RoleAssignmentScheduleRequestListResultResponse
}

type roleAssignmentScheduleRequestListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleAssignmentScheduleRequestListResultHandleError func(*azcore.Response) error

type roleAssignmentScheduleRequestListResultHandleResponse func(*azcore.Response) (RoleAssignmentScheduleRequestListResultResponse, error)

type roleAssignmentScheduleRequestListResultAdvancePage func(context.Context, RoleAssignmentScheduleRequestListResultResponse) (*azcore.Request, error)

type roleAssignmentScheduleRequestListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleAssignmentScheduleRequestListResultCreateRequest
	// callback for handling response errors
	errorer roleAssignmentScheduleRequestListResultHandleError
	// callback for handling the HTTP response
	responder roleAssignmentScheduleRequestListResultHandleResponse
	// callback for advancing to the next page
	advancer roleAssignmentScheduleRequestListResultAdvancePage
	// contains the current response
	current RoleAssignmentScheduleRequestListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleAssignmentScheduleRequestListResultPager) Err() error {
	return p.err
}

func (p *roleAssignmentScheduleRequestListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentScheduleRequestListResult.NextLink == nil || len(*p.current.RoleAssignmentScheduleRequestListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleAssignmentScheduleRequestListResultPager) PageResponse() RoleAssignmentScheduleRequestListResultResponse {
	return p.current
}

// RoleEligibilityScheduleInstanceListResultPager provides iteration over RoleEligibilityScheduleInstanceListResult pages.
type RoleEligibilityScheduleInstanceListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleEligibilityScheduleInstanceListResultResponse.
	PageResponse() RoleEligibilityScheduleInstanceListResultResponse
}

type roleEligibilityScheduleInstanceListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleEligibilityScheduleInstanceListResultHandleError func(*azcore.Response) error

type roleEligibilityScheduleInstanceListResultHandleResponse func(*azcore.Response) (RoleEligibilityScheduleInstanceListResultResponse, error)

type roleEligibilityScheduleInstanceListResultAdvancePage func(context.Context, RoleEligibilityScheduleInstanceListResultResponse) (*azcore.Request, error)

type roleEligibilityScheduleInstanceListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleEligibilityScheduleInstanceListResultCreateRequest
	// callback for handling response errors
	errorer roleEligibilityScheduleInstanceListResultHandleError
	// callback for handling the HTTP response
	responder roleEligibilityScheduleInstanceListResultHandleResponse
	// callback for advancing to the next page
	advancer roleEligibilityScheduleInstanceListResultAdvancePage
	// contains the current response
	current RoleEligibilityScheduleInstanceListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleEligibilityScheduleInstanceListResultPager) Err() error {
	return p.err
}

func (p *roleEligibilityScheduleInstanceListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleEligibilityScheduleInstanceListResult.NextLink == nil || len(*p.current.RoleEligibilityScheduleInstanceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleEligibilityScheduleInstanceListResultPager) PageResponse() RoleEligibilityScheduleInstanceListResultResponse {
	return p.current
}

// RoleEligibilityScheduleListResultPager provides iteration over RoleEligibilityScheduleListResult pages.
type RoleEligibilityScheduleListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleEligibilityScheduleListResultResponse.
	PageResponse() RoleEligibilityScheduleListResultResponse
}

type roleEligibilityScheduleListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleEligibilityScheduleListResultHandleError func(*azcore.Response) error

type roleEligibilityScheduleListResultHandleResponse func(*azcore.Response) (RoleEligibilityScheduleListResultResponse, error)

type roleEligibilityScheduleListResultAdvancePage func(context.Context, RoleEligibilityScheduleListResultResponse) (*azcore.Request, error)

type roleEligibilityScheduleListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleEligibilityScheduleListResultCreateRequest
	// callback for handling response errors
	errorer roleEligibilityScheduleListResultHandleError
	// callback for handling the HTTP response
	responder roleEligibilityScheduleListResultHandleResponse
	// callback for advancing to the next page
	advancer roleEligibilityScheduleListResultAdvancePage
	// contains the current response
	current RoleEligibilityScheduleListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleEligibilityScheduleListResultPager) Err() error {
	return p.err
}

func (p *roleEligibilityScheduleListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleEligibilityScheduleListResult.NextLink == nil || len(*p.current.RoleEligibilityScheduleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleEligibilityScheduleListResultPager) PageResponse() RoleEligibilityScheduleListResultResponse {
	return p.current
}

// RoleEligibilityScheduleRequestListResultPager provides iteration over RoleEligibilityScheduleRequestListResult pages.
type RoleEligibilityScheduleRequestListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleEligibilityScheduleRequestListResultResponse.
	PageResponse() RoleEligibilityScheduleRequestListResultResponse
}

type roleEligibilityScheduleRequestListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleEligibilityScheduleRequestListResultHandleError func(*azcore.Response) error

type roleEligibilityScheduleRequestListResultHandleResponse func(*azcore.Response) (RoleEligibilityScheduleRequestListResultResponse, error)

type roleEligibilityScheduleRequestListResultAdvancePage func(context.Context, RoleEligibilityScheduleRequestListResultResponse) (*azcore.Request, error)

type roleEligibilityScheduleRequestListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleEligibilityScheduleRequestListResultCreateRequest
	// callback for handling response errors
	errorer roleEligibilityScheduleRequestListResultHandleError
	// callback for handling the HTTP response
	responder roleEligibilityScheduleRequestListResultHandleResponse
	// callback for advancing to the next page
	advancer roleEligibilityScheduleRequestListResultAdvancePage
	// contains the current response
	current RoleEligibilityScheduleRequestListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleEligibilityScheduleRequestListResultPager) Err() error {
	return p.err
}

func (p *roleEligibilityScheduleRequestListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleEligibilityScheduleRequestListResult.NextLink == nil || len(*p.current.RoleEligibilityScheduleRequestListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleEligibilityScheduleRequestListResultPager) PageResponse() RoleEligibilityScheduleRequestListResultResponse {
	return p.current
}

// RoleManagementPolicyAssignmentListResultPager provides iteration over RoleManagementPolicyAssignmentListResult pages.
type RoleManagementPolicyAssignmentListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleManagementPolicyAssignmentListResultResponse.
	PageResponse() RoleManagementPolicyAssignmentListResultResponse
}

type roleManagementPolicyAssignmentListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleManagementPolicyAssignmentListResultHandleError func(*azcore.Response) error

type roleManagementPolicyAssignmentListResultHandleResponse func(*azcore.Response) (RoleManagementPolicyAssignmentListResultResponse, error)

type roleManagementPolicyAssignmentListResultAdvancePage func(context.Context, RoleManagementPolicyAssignmentListResultResponse) (*azcore.Request, error)

type roleManagementPolicyAssignmentListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleManagementPolicyAssignmentListResultCreateRequest
	// callback for handling response errors
	errorer roleManagementPolicyAssignmentListResultHandleError
	// callback for handling the HTTP response
	responder roleManagementPolicyAssignmentListResultHandleResponse
	// callback for advancing to the next page
	advancer roleManagementPolicyAssignmentListResultAdvancePage
	// contains the current response
	current RoleManagementPolicyAssignmentListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleManagementPolicyAssignmentListResultPager) Err() error {
	return p.err
}

func (p *roleManagementPolicyAssignmentListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleManagementPolicyAssignmentListResult.NextLink == nil || len(*p.current.RoleManagementPolicyAssignmentListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleManagementPolicyAssignmentListResultPager) PageResponse() RoleManagementPolicyAssignmentListResultResponse {
	return p.current
}

// RoleManagementPolicyListResultPager provides iteration over RoleManagementPolicyListResult pages.
type RoleManagementPolicyListResultPager interface {
	azcore.Pager

	// PageResponse returns the current RoleManagementPolicyListResultResponse.
	PageResponse() RoleManagementPolicyListResultResponse
}

type roleManagementPolicyListResultCreateRequest func(context.Context) (*azcore.Request, error)

type roleManagementPolicyListResultHandleError func(*azcore.Response) error

type roleManagementPolicyListResultHandleResponse func(*azcore.Response) (RoleManagementPolicyListResultResponse, error)

type roleManagementPolicyListResultAdvancePage func(context.Context, RoleManagementPolicyListResultResponse) (*azcore.Request, error)

type roleManagementPolicyListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester roleManagementPolicyListResultCreateRequest
	// callback for handling response errors
	errorer roleManagementPolicyListResultHandleError
	// callback for handling the HTTP response
	responder roleManagementPolicyListResultHandleResponse
	// callback for advancing to the next page
	advancer roleManagementPolicyListResultAdvancePage
	// contains the current response
	current RoleManagementPolicyListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *roleManagementPolicyListResultPager) Err() error {
	return p.err
}

func (p *roleManagementPolicyListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleManagementPolicyListResult.NextLink == nil || len(*p.current.RoleManagementPolicyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleManagementPolicyListResultPager) PageResponse() RoleManagementPolicyListResultResponse {
	return p.current
}