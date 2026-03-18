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

// Currency exchange rates and conversion
//
// CurrencyService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyService] method instead.
type CurrencyService struct {
	options []option.RequestOption
}

// NewCurrencyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrencyService(opts ...option.RequestOption) (r CurrencyService) {
	r = CurrencyService{}
	r.options = opts
	return
}

// Returns a list of all currency conversion rates sourced from HMRC (GBP) and the
// European Central Bank (EUR).
//
// You can optionally filter by source and target currency.
func (r *CurrencyService) List(ctx context.Context, query CurrencyListParams, opts ...option.RequestOption) (res *CurrencyListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "currency"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Calculate the inclusive and exclusive VAT price on a given amount and VAT rate.
// This is a standalone calculation that does not look up rates by country.
func (r *CurrencyService) CalculateVatPrice(ctx context.Context, query CurrencyCalculateVatPriceParams, opts ...option.RequestOption) (res *CurrencyCalculateVatPriceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "currency/price"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Convert a foreign currency amount to either GBP or EUR using official exchange
// rates.
//
// GBP rates are from HMRC (updated on the 1st of every month). EUR rates are from
// the European Central Bank (updated around 16:00 CET on working days).
func (r *CurrencyService) Convert(ctx context.Context, query CurrencyConvertParams, opts ...option.RequestOption) (res *CurrencyConvertResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "currency/convert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type VatPrice struct {
	// Any of "vat_price".
	Object VatPriceObject `json:"object"`
	// The price provided.
	Price float64 `json:"price"`
	// The calculated price exclusive of VAT.
	PriceExclVat float64 `json:"price_excl_vat"`
	// The calculated price inclusive of VAT.
	PriceInclVat float64 `json:"price_incl_vat"`
	// Whether the price is inclusive or exclusive of VAT.
	//
	// Any of "incl", "excl".
	TaxType VatPriceTaxType `json:"tax_type"`
	// The total VAT amount.
	Vat float64 `json:"vat"`
	// The VAT rate percentage.
	VatRate float64 `json:"vat_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object       respjson.Field
		Price        respjson.Field
		PriceExclVat respjson.Field
		PriceInclVat respjson.Field
		TaxType      respjson.Field
		Vat          respjson.Field
		VatRate      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VatPrice) RawJSON() string { return r.JSON.raw }
func (r *VatPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VatPriceObject string

const (
	VatPriceObjectVatPrice VatPriceObject = "vat_price"
)

// Whether the price is inclusive or exclusive of VAT.
type VatPriceTaxType string

const (
	VatPriceTaxTypeIncl VatPriceTaxType = "incl"
	VatPriceTaxTypeExcl VatPriceTaxType = "excl"
)

type CurrencyListResponse struct {
	Code    int64                      `json:"code"`
	Data    []CurrencyListResponseData `json:"data"`
	Success bool                       `json:"success"`
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
func (r CurrencyListResponse) RawJSON() string { return r.JSON.raw }
func (r *CurrencyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyListResponseData struct {
	// The 3-character source currency code.
	From string `json:"from"`
	// Any of "convert_rate".
	Object string `json:"object"`
	// The exchange rate.
	Rate float64 `json:"rate"`
	// The 3-character target currency code (GBP or EUR).
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Object      respjson.Field
		Rate        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CurrencyListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyCalculateVatPriceResponse struct {
	Code    int64    `json:"code"`
	Data    VatPrice `json:"data"`
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
func (r CurrencyCalculateVatPriceResponse) RawJSON() string { return r.JSON.raw }
func (r *CurrencyCalculateVatPriceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyConvertResponse struct {
	Code    int64                       `json:"code"`
	Data    CurrencyConvertResponseData `json:"data"`
	Success bool                        `json:"success"`
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
func (r CurrencyConvertResponse) RawJSON() string { return r.JSON.raw }
func (r *CurrencyConvertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyConvertResponseData struct {
	// The original amount.
	Amount float64 `json:"amount"`
	// The converted amount.
	Converted float64 `json:"converted"`
	From      string  `json:"from"`
	// Any of "conversion".
	Object string `json:"object"`
	// The exchange rate used.
	Rate float64 `json:"rate"`
	To   string  `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Converted   respjson.Field
		From        respjson.Field
		Object      respjson.Field
		Rate        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyConvertResponseData) RawJSON() string { return r.JSON.raw }
func (r *CurrencyConvertResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyListParams struct {
	// The 3-character currency code(s) to convert from (e.g. "USD", "CAD"). Can be a
	// comma-separated list without spaces (e.g. "USD,CAD,AUD").
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// The 3-character target currency code. Must be either "GBP" or "EUR".
	//
	// Any of "GBP", "EUR".
	To CurrencyListParamsTo `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListParams]'s query parameters as `url.Values`.
func (r CurrencyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The 3-character target currency code. Must be either "GBP" or "EUR".
type CurrencyListParamsTo string

const (
	CurrencyListParamsToGbp CurrencyListParamsTo = "GBP"
	CurrencyListParamsToEur CurrencyListParamsTo = "EUR"
)

type CurrencyCalculateVatPriceParams struct {
	// The price to calculate on. Must be a string with exactly 2 decimal places (e.g.
	// "30.00", "59.95").
	Price string `query:"price" api:"required" json:"-"`
	// Whether the provided price is inclusive or exclusive of VAT.
	//
	// Any of "incl", "excl".
	TaxType CurrencyCalculateVatPriceParamsTaxType `query:"tax_type,omitzero" api:"required" json:"-"`
	// A percentage VAT rate to use for the calculation.
	VatRate float64 `query:"vat_rate" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyCalculateVatPriceParams]'s query parameters as
// `url.Values`.
func (r CurrencyCalculateVatPriceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether the provided price is inclusive or exclusive of VAT.
type CurrencyCalculateVatPriceParamsTaxType string

const (
	CurrencyCalculateVatPriceParamsTaxTypeIncl CurrencyCalculateVatPriceParamsTaxType = "incl"
	CurrencyCalculateVatPriceParamsTaxTypeExcl CurrencyCalculateVatPriceParamsTaxType = "excl"
)

type CurrencyConvertParams struct {
	// The amount to convert. Must be a string with exactly 2 decimal places (e.g.
	// "39.99").
	Amount string `query:"amount" api:"required" json:"-"`
	// The 3-character source currency code (e.g. "USD", "CAD").
	From string `query:"from" api:"required" json:"-"`
	// The 3-character target currency code. Must be either "GBP" or "EUR".
	//
	// Any of "GBP", "EUR".
	To CurrencyConvertParamsTo `query:"to,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyConvertParams]'s query parameters as `url.Values`.
func (r CurrencyConvertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The 3-character target currency code. Must be either "GBP" or "EUR".
type CurrencyConvertParamsTo string

const (
	CurrencyConvertParamsToGbp CurrencyConvertParamsTo = "GBP"
	CurrencyConvertParamsToEur CurrencyConvertParamsTo = "EUR"
)
