// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/apiquery"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
	"github.com/VAT-Sense/vatsense-go/packages/param"
	"github.com/VAT-Sense/vatsense-go/packages/respjson"
)

// Country and province information
//
// CountryService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCountryService] method instead.
type CountryService struct {
	options []option.RequestOption
}

// NewCountryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCountryService(opts ...option.RequestOption) (r CountryService) {
	r = CountryService{}
	r.options = opts
	return
}

// Returns a list of all countries, including whether they are subject to VAT/GST
// and whether they are subject to EU VAT. Each country is returned as a country
// object.
//
// You can optionally filter by country code or IP address.
func (r *CountryService) List(ctx context.Context, query CountryListParams, opts ...option.RequestOption) (res *CountryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a list of all provinces within a given country.
func (r *CountryService) ListProvinces(ctx context.Context, query CountryListProvincesParams, opts ...option.RequestOption) (res *CountryListProvincesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "countries/provinces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Country struct {
	// 2-character ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	// Whether the country is subject to EU VAT.
	Eu        bool    `json:"eu"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	// Any of "country".
	Object CountryObject `json:"object"`
	// Whether the country is subject to VAT/GST.
	Vat bool `json:"vat"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode respjson.Field
		CountryName respjson.Field
		Eu          respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		Object      respjson.Field
		Vat         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Country) RawJSON() string { return r.JSON.raw }
func (r *Country) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryObject string

const (
	CountryObjectCountry CountryObject = "country"
)

type CountryListResponse struct {
	Code    int64     `json:"code"`
	Data    []Country `json:"data"`
	Success bool      `json:"success"`
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
func (r CountryListResponse) RawJSON() string { return r.JSON.raw }
func (r *CountryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryListProvincesResponse struct {
	Code    int64                              `json:"code"`
	Data    []CountryListProvincesResponseData `json:"data"`
	Success bool                               `json:"success"`
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
func (r CountryListProvincesResponse) RawJSON() string { return r.JSON.raw }
func (r *CountryListProvincesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryListProvincesResponseData struct {
	CountryCode string `json:"country_code"`
	// Any of "province".
	Object       string `json:"object"`
	ProvinceCode string `json:"province_code"`
	ProvinceName string `json:"province_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode  respjson.Field
		Object       respjson.Field
		ProvinceCode respjson.Field
		ProvinceName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CountryListProvincesResponseData) RawJSON() string { return r.JSON.raw }
func (r *CountryListProvincesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CountryListParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR"). Overseas
	// territories that carry their own ISO code but are modelled as provinces of a
	// parent country (e.g. "NC" New Caledonia, "MF" Saint Martin, "GP", "MQ", "RE",
	// "PF", "GF", "YT", "BL", "PM", "WF" under "FR") may be queried directly; the
	// response identifies the territory and the rate is the one the
	// parent-plus-province query returns.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CountryListParams]'s query parameters as `url.Values`.
func (r CountryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CountryListProvincesParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "CA").
	CountryCode string `query:"country_code" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CountryListProvincesParams]'s query parameters as
// `url.Values`.
func (r CountryListProvincesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
