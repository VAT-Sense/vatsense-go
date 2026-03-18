// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"net/http"
	"slices"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
	"github.com/VAT-Sense/vatsense-go/packages/respjson"
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
	options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r UsageService) {
	r = UsageService{}
	r.options = opts
	return
}

// Check your used and remaining API requests.
func (r *UsageService) Get(ctx context.Context, opts ...option.RequestOption) (res *UsageGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UsageGetResponse struct {
	Code    int64                `json:"code"`
	Data    UsageGetResponseData `json:"data"`
	Success bool                 `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageGetResponseData struct {
	Requests UsageGetResponseDataRequests `json:"requests"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Requests    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageGetResponseDataRequests struct {
	// Requests remaining before the limit is reached.
	Remaining int64 `json:"remaining"`
	// Total requests allowed on your plan.
	Total int64 `json:"total"`
	// Requests used in the last 30 days.
	Used int64 `json:"used"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Remaining   respjson.Field
		Total       respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageGetResponseDataRequests) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponseDataRequests) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
