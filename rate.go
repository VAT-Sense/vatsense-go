// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/apiquery"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
	"github.com/VAT-Sense/vatsense-go/packages/param"
	"github.com/VAT-Sense/vatsense-go/packages/respjson"
)

// VAT/GST rate lookups for countries worldwide
//
// RateService contains methods and other services that help with interacting with
// the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRateService] method instead.
type RateService struct {
	options []option.RequestOption
}

// NewRateService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRateService(opts ...option.RequestOption) (r RateService) {
	r = RateService{}
	r.options = opts
	return
}

// Returns a list of VAT/GST rates for all countries, sorted alphabetically by
// country code. Each rate is returned as a rate object containing the standard
// rate and any other applicable rates.
//
// You can optionally filter by country code, IP address, or EU membership.
func (r *RateService) List(ctx context.Context, query RateListParams, opts ...option.RequestOption) (res *RateListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Combines the functionality of the "Find a tax rate" and "VAT price calculation"
// endpoints to return the particular VAT price for an applicable VAT rate.
// Requires both a location (country_code or ip_address) and a price to calculate.
func (r *RateService) CalculatePrice(ctx context.Context, query RateCalculatePriceParams, opts ...option.RequestOption) (res *RateCalculatePriceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rates/price"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get detailed tax rate information for a location, including all applicable rate
// classes (standard, reduced, zero, etc.).
func (r *RateService) Details(ctx context.Context, query RateDetailsParams, opts ...option.RequestOption) (res *FindRate, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rates/tax_rate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A handy endpoint for finding a rate that applies to a particular country and
// optional product type, based on country code or IP address.
//
// If no type is provided, or no specific rate is applied to the given type, then
// the standard rate will be returned if the country is subject to tax.
//
// If the country is not subject to VAT/GST then an error response will be
// returned, indicating no tax applies.
func (r *RateService) Find(ctx context.Context, query RateFindParams, opts ...option.RequestOption) (res *FindRate, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rates/rate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a list of all available product types that can be used to filter tax
// rates.
func (r *RateService) ListTypes(ctx context.Context, opts ...option.RequestOption) (res *RateListTypesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rates/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type FindRate struct {
	Code    int64           `json:"code"`
	Data    RateWithTaxRate `json:"data"`
	Success bool            `json:"success"`
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
func (r FindRate) RawJSON() string { return r.JSON.raw }
func (r *FindRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Rate struct {
	// 2-character ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	// Whether the country is an EU member.
	Eu bool `json:"eu"`
	// Any of "rate".
	Object RateObject `json:"object"`
	// A list of other tax rates. Null if no additional rates exist.
	Other    []RateOther `json:"other" api:"nullable"`
	Standard TaxRate     `json:"standard"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode respjson.Field
		CountryName respjson.Field
		Eu          respjson.Field
		Object      respjson.Field
		Other       respjson.Field
		Standard    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rate) RawJSON() string { return r.JSON.raw }
func (r *Rate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateObject string

const (
	RateObjectRate RateObject = "rate"
)

type RateOther struct {
	// The province this rate applies to, if applicable.
	Province string `json:"province" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Province    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	TaxRate
}

// Returns the unmodified JSON received from the API
func (r RateOther) RawJSON() string { return r.JSON.raw }
func (r *RateOther) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateWithTaxRate struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Eu          bool   `json:"eu"`
	// Any of "rate".
	Object  RateWithTaxRateObject `json:"object"`
	TaxRate TaxRate               `json:"tax_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode respjson.Field
		CountryName respjson.Field
		Eu          respjson.Field
		Object      respjson.Field
		TaxRate     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RateWithTaxRate) RawJSON() string { return r.JSON.raw }
func (r *RateWithTaxRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateWithTaxRateObject string

const (
	RateWithTaxRateObjectRate RateWithTaxRateObject = "rate"
)

type TaxRate struct {
	// The rate class (e.g. "standard", "reduced", "zero").
	Class string `json:"class"`
	// A description of what goods/services this rate applies to.
	Description string `json:"description"`
	// Any of "tax_rate".
	Object TaxRateObject `json:"object"`
	// The tax rate percentage.
	Rate float64 `json:"rate"`
	// Comma-separated list of product types this rate applies to, or false if it
	// applies generally.
	Types TaxRateTypesUnion `json:"types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Class       respjson.Field
		Description respjson.Field
		Object      respjson.Field
		Rate        respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaxRate) RawJSON() string { return r.JSON.raw }
func (r *TaxRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaxRateObject string

const (
	TaxRateObjectTaxRate TaxRateObject = "tax_rate"
)

// TaxRateTypesUnion contains all possible properties and values from [string],
// [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfBool]
type TaxRateTypesUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u TaxRateTypesUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaxRateTypesUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaxRateTypesUnion) RawJSON() string { return u.JSON.raw }

func (r *TaxRateTypesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateListResponse struct {
	Code    int64  `json:"code"`
	Data    []Rate `json:"data"`
	Success bool   `json:"success"`
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
func (r RateListResponse) RawJSON() string { return r.JSON.raw }
func (r *RateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateCalculatePriceResponse struct {
	Code    int64                          `json:"code"`
	Data    RateCalculatePriceResponseData `json:"data"`
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
func (r RateCalculatePriceResponse) RawJSON() string { return r.JSON.raw }
func (r *RateCalculatePriceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateCalculatePriceResponseData struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Eu          bool   `json:"eu"`
	// Any of "rate".
	Object   string   `json:"object"`
	TaxRate  TaxRate  `json:"tax_rate"`
	VatPrice VatPrice `json:"vat_price"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode respjson.Field
		CountryName respjson.Field
		Eu          respjson.Field
		Object      respjson.Field
		TaxRate     respjson.Field
		VatPrice    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RateCalculatePriceResponseData) RawJSON() string { return r.JSON.raw }
func (r *RateCalculatePriceResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateListTypesResponse struct {
	Code    int64    `json:"code"`
	Data    []string `json:"data"`
	Success bool     `json:"success"`
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
func (r RateListTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *RateListTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RateListParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Opt[bool] `query:"eu,omitzero" json:"-"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// A historical date to retrieve rates for (format "YYYY-MM-DD HH:MM:SS"). Must be
	// a past date.
	Period param.Opt[time.Time] `query:"period,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [RateListParams]'s query parameters as `url.Values`.
func (r RateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RateCalculatePriceParams struct {
	// The price to calculate on. Must be a string with exactly 2 decimal places (e.g.
	// "30.00", "59.95").
	Price string `query:"price" api:"required" json:"-"`
	// Whether the provided price is inclusive or exclusive of VAT.
	//
	// Any of "incl", "excl".
	TaxType RateCalculatePriceParamsTaxType `query:"tax_type,omitzero" api:"required" json:"-"`
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Opt[bool] `query:"eu,omitzero" json:"-"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// A 2-character province code (e.g. "NU", "NT"). If providing a province code, you
	// must also provide the relevant country_code.
	ProvinceCode param.Opt[string] `query:"province_code,omitzero" json:"-"`
	// The product type to find the applicable rate for. See the /rates/types endpoint
	// for a full list of valid values.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RateCalculatePriceParams]'s query parameters as
// `url.Values`.
func (r RateCalculatePriceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether the provided price is inclusive or exclusive of VAT.
type RateCalculatePriceParamsTaxType string

const (
	RateCalculatePriceParamsTaxTypeIncl RateCalculatePriceParamsTaxType = "incl"
	RateCalculatePriceParamsTaxTypeExcl RateCalculatePriceParamsTaxType = "excl"
)

type RateDetailsParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Opt[bool] `query:"eu,omitzero" json:"-"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// A historical date to retrieve rates for (format "YYYY-MM-DD HH:MM:SS"). Must be
	// a past date.
	Period param.Opt[time.Time] `query:"period,omitzero" format:"date-time" json:"-"`
	// A 2-character province code (e.g. "NU", "NT"). If providing a province code, you
	// must also provide the relevant country_code.
	ProvinceCode param.Opt[string] `query:"province_code,omitzero" json:"-"`
	// The product type to find the applicable rate for. See the /rates/types endpoint
	// for a full list of valid values.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RateDetailsParams]'s query parameters as `url.Values`.
func (r RateDetailsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RateFindParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Opt[bool] `query:"eu,omitzero" json:"-"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// A historical date to retrieve rates for (format "YYYY-MM-DD HH:MM:SS"). Must be
	// a past date.
	Period param.Opt[time.Time] `query:"period,omitzero" format:"date-time" json:"-"`
	// A 2-character province code (e.g. "NU", "NT"). If providing a province code, you
	// must also provide the relevant country_code.
	ProvinceCode param.Opt[string] `query:"province_code,omitzero" json:"-"`
	// The product type to find the applicable rate for. See the /rates/types endpoint
	// for a full list of valid values.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RateFindParams]'s query parameters as `url.Values`.
func (r RateFindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
