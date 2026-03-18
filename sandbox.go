// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
	"github.com/VAT-Sense/vatsense-go/packages/respjson"
)

// Temporary sandbox API keys for testing
//
// SandboxService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxService] method instead.
type SandboxService struct {
	options []option.RequestOption
}

// NewSandboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSandboxService(opts ...option.RequestOption) (r SandboxService) {
	r = SandboxService{}
	r.options = opts
	return
}

// Generate a temporary sandbox API key for testing. Sandbox keys have limited
// request allowances and restricted endpoint access (no invoice endpoints). Rate
// limited to 1 key per IP address per 6 hours.
func (r *SandboxService) GenerateKey(ctx context.Context, opts ...option.RequestOption) (res *SandboxGenerateKeyResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "sandbox/key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type SandboxGenerateKeyResponse struct {
	Code    int64                          `json:"code"`
	Data    SandboxGenerateKeyResponseData `json:"data"`
	Success bool                           `json:"success"`
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
func (r SandboxGenerateKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxGenerateKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxGenerateKeyResponseData struct {
	AllowedEndpoints []string  `json:"allowed_endpoints"`
	ExpiresAt        time.Time `json:"expires_at" format:"date-time"`
	// The temporary sandbox API key.
	Key               string `json:"key"`
	RequestsRemaining int64  `json:"requests_remaining"`
	SignupURL         string `json:"signup_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedEndpoints  respjson.Field
		ExpiresAt         respjson.Field
		Key               respjson.Field
		RequestsRemaining respjson.Field
		SignupURL         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxGenerateKeyResponseData) RawJSON() string { return r.JSON.raw }
func (r *SandboxGenerateKeyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
