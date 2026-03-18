// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/apiquery"
	"github.com/VAT-Sense/vatsense-go/internal/param"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
	"github.com/VAT-Sense/vatsense-go/shared"
	"github.com/tidwall/gjson"
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
	Options []option.RequestOption
}

// NewRateService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRateService(opts ...option.RequestOption) (r *RateService) {
	r = &RateService{}
	r.Options = opts
	return
}

// Returns a list of VAT/GST rates for all countries, sorted alphabetically by
// country code. Each rate is returned as a rate object containing the standard
// rate and any other applicable rates.
//
// You can optionally filter by country code, IP address, or EU membership.
func (r *RateService) List(ctx context.Context, query RateListParams, opts ...option.RequestOption) (res *RateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Combines the functionality of the "Find a tax rate" and "VAT price calculation"
// endpoints to return the particular VAT price for an applicable VAT rate.
// Requires both a location (country_code or ip_address) and a price to calculate.
func (r *RateService) CalculatePrice(ctx context.Context, query RateCalculatePriceParams, opts ...option.RequestOption) (res *RateCalculatePriceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rates/price"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get detailed tax rate information for a location, including all applicable rate
// classes (standard, reduced, zero, etc.).
func (r *RateService) Details(ctx context.Context, query RateDetailsParams, opts ...option.RequestOption) (res *FindRate, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	path := "rates/rate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a list of all available product types that can be used to filter tax
// rates.
func (r *RateService) ListTypes(ctx context.Context, opts ...option.RequestOption) (res *RateListTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rates/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type FindRate struct {
	Code    int64           `json:"code"`
	Data    RateWithTaxRate `json:"data"`
	Success bool            `json:"success"`
	JSON    findRateJSON    `json:"-"`
}

// findRateJSON contains the JSON metadata for the struct [FindRate]
type findRateJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FindRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r findRateJSON) RawJSON() string {
	return r.raw
}

type Rate struct {
	// 2-character ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	// Whether the country is an EU member.
	Eu     bool       `json:"eu"`
	Object RateObject `json:"object"`
	// A list of other tax rates. Null if no additional rates exist.
	Other    []RateOther `json:"other" api:"nullable"`
	Standard TaxRate     `json:"standard"`
	JSON     rateJSON    `json:"-"`
}

// rateJSON contains the JSON metadata for the struct [Rate]
type rateJSON struct {
	CountryCode apijson.Field
	CountryName apijson.Field
	Eu          apijson.Field
	Object      apijson.Field
	Other       apijson.Field
	Standard    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Rate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateJSON) RawJSON() string {
	return r.raw
}

type RateObject string

const (
	RateObjectRate RateObject = "rate"
)

func (r RateObject) IsKnown() bool {
	switch r {
	case RateObjectRate:
		return true
	}
	return false
}

type RateOther struct {
	// The province this rate applies to, if applicable.
	Province string        `json:"province" api:"nullable"`
	JSON     rateOtherJSON `json:"-"`
	TaxRate
}

// rateOtherJSON contains the JSON metadata for the struct [RateOther]
type rateOtherJSON struct {
	Province    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateOther) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateOtherJSON) RawJSON() string {
	return r.raw
}

type RateWithTaxRate struct {
	CountryCode string                `json:"country_code"`
	CountryName string                `json:"country_name"`
	Eu          bool                  `json:"eu"`
	Object      RateWithTaxRateObject `json:"object"`
	TaxRate     TaxRate               `json:"tax_rate"`
	JSON        rateWithTaxRateJSON   `json:"-"`
}

// rateWithTaxRateJSON contains the JSON metadata for the struct [RateWithTaxRate]
type rateWithTaxRateJSON struct {
	CountryCode apijson.Field
	CountryName apijson.Field
	Eu          apijson.Field
	Object      apijson.Field
	TaxRate     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateWithTaxRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateWithTaxRateJSON) RawJSON() string {
	return r.raw
}

type RateWithTaxRateObject string

const (
	RateWithTaxRateObjectRate RateWithTaxRateObject = "rate"
)

func (r RateWithTaxRateObject) IsKnown() bool {
	switch r {
	case RateWithTaxRateObjectRate:
		return true
	}
	return false
}

type TaxRate struct {
	// The rate class (e.g. "standard", "reduced", "zero").
	Class string `json:"class"`
	// A description of what goods/services this rate applies to.
	Description string        `json:"description"`
	Object      TaxRateObject `json:"object"`
	// The tax rate percentage.
	Rate float64 `json:"rate"`
	// Comma-separated list of product types this rate applies to, or false if it
	// applies generally.
	Types TaxRateTypesUnion `json:"types"`
	JSON  taxRateJSON       `json:"-"`
}

// taxRateJSON contains the JSON metadata for the struct [TaxRate]
type taxRateJSON struct {
	Class       apijson.Field
	Description apijson.Field
	Object      apijson.Field
	Rate        apijson.Field
	Types       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaxRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taxRateJSON) RawJSON() string {
	return r.raw
}

type TaxRateObject string

const (
	TaxRateObjectTaxRate TaxRateObject = "tax_rate"
)

func (r TaxRateObject) IsKnown() bool {
	switch r {
	case TaxRateObjectTaxRate:
		return true
	}
	return false
}

// Comma-separated list of product types this rate applies to, or false if it
// applies generally.
//
// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type TaxRateTypesUnion interface {
	ImplementsTaxRateTypesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TaxRateTypesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type RateListResponse struct {
	Code    int64                `json:"code"`
	Data    []Rate               `json:"data"`
	Success bool                 `json:"success"`
	JSON    rateListResponseJSON `json:"-"`
}

// rateListResponseJSON contains the JSON metadata for the struct
// [RateListResponse]
type rateListResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateListResponseJSON) RawJSON() string {
	return r.raw
}

type RateCalculatePriceResponse struct {
	Code    int64                          `json:"code"`
	Data    RateCalculatePriceResponseData `json:"data"`
	Success bool                           `json:"success"`
	JSON    rateCalculatePriceResponseJSON `json:"-"`
}

// rateCalculatePriceResponseJSON contains the JSON metadata for the struct
// [RateCalculatePriceResponse]
type rateCalculatePriceResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateCalculatePriceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateCalculatePriceResponseJSON) RawJSON() string {
	return r.raw
}

type RateCalculatePriceResponseData struct {
	CountryCode string                               `json:"country_code"`
	CountryName string                               `json:"country_name"`
	Eu          bool                                 `json:"eu"`
	Object      RateCalculatePriceResponseDataObject `json:"object"`
	TaxRate     TaxRate                              `json:"tax_rate"`
	VatPrice    VatPrice                             `json:"vat_price"`
	JSON        rateCalculatePriceResponseDataJSON   `json:"-"`
}

// rateCalculatePriceResponseDataJSON contains the JSON metadata for the struct
// [RateCalculatePriceResponseData]
type rateCalculatePriceResponseDataJSON struct {
	CountryCode apijson.Field
	CountryName apijson.Field
	Eu          apijson.Field
	Object      apijson.Field
	TaxRate     apijson.Field
	VatPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateCalculatePriceResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateCalculatePriceResponseDataJSON) RawJSON() string {
	return r.raw
}

type RateCalculatePriceResponseDataObject string

const (
	RateCalculatePriceResponseDataObjectRate RateCalculatePriceResponseDataObject = "rate"
)

func (r RateCalculatePriceResponseDataObject) IsKnown() bool {
	switch r {
	case RateCalculatePriceResponseDataObjectRate:
		return true
	}
	return false
}

type RateListTypesResponse struct {
	Code    int64                     `json:"code"`
	Data    []string                  `json:"data"`
	Success bool                      `json:"success"`
	JSON    rateListTypesResponseJSON `json:"-"`
}

// rateListTypesResponseJSON contains the JSON metadata for the struct
// [RateListTypesResponse]
type rateListTypesResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateListTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateListTypesResponseJSON) RawJSON() string {
	return r.raw
}

type RateListParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Field[string] `query:"country_code"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Field[bool] `query:"eu"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Field[string] `query:"ip_address"`
	// A historical date to retrieve rates for (format "YYYY-MM-DD HH:MM:SS"). Must be
	// a past date.
	Period param.Field[time.Time] `query:"period" format:"date-time"`
}

// URLQuery serializes [RateListParams]'s query parameters as `url.Values`.
func (r RateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RateCalculatePriceParams struct {
	// The price to calculate on. Must be a string with exactly 2 decimal places (e.g.
	// "30.00", "59.95").
	Price param.Field[string] `query:"price" api:"required"`
	// Whether the provided price is inclusive or exclusive of VAT.
	TaxType param.Field[RateCalculatePriceParamsTaxType] `query:"tax_type" api:"required"`
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Field[string] `query:"country_code"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Field[bool] `query:"eu"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Field[string] `query:"ip_address"`
	// A 2-character province code (e.g. "NU", "NT"). If providing a province code, you
	// must also provide the relevant country_code.
	ProvinceCode param.Field[string] `query:"province_code"`
	// The product type to find the applicable rate for. See the /rates/types endpoint
	// for a full list of valid values.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [RateCalculatePriceParams]'s query parameters as
// `url.Values`.
func (r RateCalculatePriceParams) URLQuery() (v url.Values) {
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

func (r RateCalculatePriceParamsTaxType) IsKnown() bool {
	switch r {
	case RateCalculatePriceParamsTaxTypeIncl, RateCalculatePriceParamsTaxTypeExcl:
		return true
	}
	return false
}

type RateDetailsParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Field[string] `query:"country_code"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Field[bool] `query:"eu"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Field[string] `query:"ip_address"`
	// A historical date to retrieve rates for (format "YYYY-MM-DD HH:MM:SS"). Must be
	// a past date.
	Period param.Field[time.Time] `query:"period" format:"date-time"`
	// A 2-character province code (e.g. "NU", "NT"). If providing a province code, you
	// must also provide the relevant country_code.
	ProvinceCode param.Field[string] `query:"province_code"`
	// The product type to find the applicable rate for. See the /rates/types endpoint
	// for a full list of valid values.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [RateDetailsParams]'s query parameters as `url.Values`.
func (r RateDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RateFindParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Field[string] `query:"country_code"`
	// Filter results by EU membership. Use 1 for EU countries only, 0 for non-EU only.
	Eu param.Field[bool] `query:"eu"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Field[string] `query:"ip_address"`
	// A historical date to retrieve rates for (format "YYYY-MM-DD HH:MM:SS"). Must be
	// a past date.
	Period param.Field[time.Time] `query:"period" format:"date-time"`
	// A 2-character province code (e.g. "NU", "NT"). If providing a province code, you
	// must also provide the relevant country_code.
	ProvinceCode param.Field[string] `query:"province_code"`
	// The product type to find the applicable rate for. See the /rates/types endpoint
	// for a full list of valid values.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [RateFindParams]'s query parameters as `url.Values`.
func (r RateFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
