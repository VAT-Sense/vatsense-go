// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"net/http"
	"slices"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
)

// API usage statistics
//
// UsageService contains methods and other services that help with interacting with
// the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r *UsageService) {
	r = &UsageService{}
	r.Options = opts
	return
}

// Check your used and remaining API requests.
func (r *UsageService) Get(ctx context.Context, opts ...option.RequestOption) (res *UsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UsageGetResponse struct {
	Code    int64                `json:"code"`
	Data    UsageGetResponseData `json:"data"`
	Success bool                 `json:"success"`
	JSON    usageGetResponseJSON `json:"-"`
}

// usageGetResponseJSON contains the JSON metadata for the struct
// [UsageGetResponse]
type usageGetResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponseData struct {
	Requests UsageGetResponseDataRequests `json:"requests"`
	JSON     usageGetResponseDataJSON     `json:"-"`
}

// usageGetResponseDataJSON contains the JSON metadata for the struct
// [UsageGetResponseData]
type usageGetResponseDataJSON struct {
	Requests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type UsageGetResponseDataRequests struct {
	// Requests remaining before the limit is reached.
	Remaining int64 `json:"remaining"`
	// Total requests allowed on your plan.
	Total int64 `json:"total"`
	// Requests used in the last 30 days.
	Used int64                            `json:"used"`
	JSON usageGetResponseDataRequestsJSON `json:"-"`
}

// usageGetResponseDataRequestsJSON contains the JSON metadata for the struct
// [UsageGetResponseDataRequests]
type usageGetResponseDataRequestsJSON struct {
	Remaining   apijson.Field
	Total       apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageGetResponseDataRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageGetResponseDataRequestsJSON) RawJSON() string {
	return r.raw
}
