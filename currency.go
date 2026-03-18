// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/apiquery"
	"github.com/VAT-Sense/vatsense-go/internal/param"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
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
	Options []option.RequestOption
}

// NewCurrencyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrencyService(opts ...option.RequestOption) (r *CurrencyService) {
	r = &CurrencyService{}
	r.Options = opts
	return
}

// Returns a list of all currency conversion rates sourced from HMRC (GBP) and the
// European Central Bank (EUR).
//
// You can optionally filter by source and target currency.
func (r *CurrencyService) List(ctx context.Context, query CurrencyListParams, opts ...option.RequestOption) (res *CurrencyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "currency"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Calculate the inclusive and exclusive VAT price on a given amount and VAT rate.
// This is a standalone calculation that does not look up rates by country.
func (r *CurrencyService) CalculateVatPrice(ctx context.Context, query CurrencyCalculateVatPriceParams, opts ...option.RequestOption) (res *CurrencyCalculateVatPriceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	path := "currency/convert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type VatPrice struct {
	Object VatPriceObject `json:"object"`
	// The price provided.
	Price float64 `json:"price"`
	// The calculated price exclusive of VAT.
	PriceExclVat float64 `json:"price_excl_vat"`
	// The calculated price inclusive of VAT.
	PriceInclVat float64 `json:"price_incl_vat"`
	// Whether the price is inclusive or exclusive of VAT.
	TaxType VatPriceTaxType `json:"tax_type"`
	// The total VAT amount.
	Vat float64 `json:"vat"`
	// The VAT rate percentage.
	VatRate float64      `json:"vat_rate"`
	JSON    vatPriceJSON `json:"-"`
}

// vatPriceJSON contains the JSON metadata for the struct [VatPrice]
type vatPriceJSON struct {
	Object       apijson.Field
	Price        apijson.Field
	PriceExclVat apijson.Field
	PriceInclVat apijson.Field
	TaxType      apijson.Field
	Vat          apijson.Field
	VatRate      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VatPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vatPriceJSON) RawJSON() string {
	return r.raw
}

type VatPriceObject string

const (
	VatPriceObjectVatPrice VatPriceObject = "vat_price"
)

func (r VatPriceObject) IsKnown() bool {
	switch r {
	case VatPriceObjectVatPrice:
		return true
	}
	return false
}

// Whether the price is inclusive or exclusive of VAT.
type VatPriceTaxType string

const (
	VatPriceTaxTypeIncl VatPriceTaxType = "incl"
	VatPriceTaxTypeExcl VatPriceTaxType = "excl"
)

func (r VatPriceTaxType) IsKnown() bool {
	switch r {
	case VatPriceTaxTypeIncl, VatPriceTaxTypeExcl:
		return true
	}
	return false
}

type CurrencyListResponse struct {
	Code    int64                      `json:"code"`
	Data    []CurrencyListResponseData `json:"data"`
	Success bool                       `json:"success"`
	JSON    currencyListResponseJSON   `json:"-"`
}

// currencyListResponseJSON contains the JSON metadata for the struct
// [CurrencyListResponse]
type currencyListResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrencyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currencyListResponseJSON) RawJSON() string {
	return r.raw
}

type CurrencyListResponseData struct {
	// The 3-character source currency code.
	From   string                         `json:"from"`
	Object CurrencyListResponseDataObject `json:"object"`
	// The exchange rate.
	Rate float64 `json:"rate"`
	// The 3-character target currency code (GBP or EUR).
	To   string                       `json:"to"`
	JSON currencyListResponseDataJSON `json:"-"`
}

// currencyListResponseDataJSON contains the JSON metadata for the struct
// [CurrencyListResponseData]
type currencyListResponseDataJSON struct {
	From        apijson.Field
	Object      apijson.Field
	Rate        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrencyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currencyListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CurrencyListResponseDataObject string

const (
	CurrencyListResponseDataObjectConvertRate CurrencyListResponseDataObject = "convert_rate"
)

func (r CurrencyListResponseDataObject) IsKnown() bool {
	switch r {
	case CurrencyListResponseDataObjectConvertRate:
		return true
	}
	return false
}

type CurrencyCalculateVatPriceResponse struct {
	Code    int64                                 `json:"code"`
	Data    VatPrice                              `json:"data"`
	Success bool                                  `json:"success"`
	JSON    currencyCalculateVatPriceResponseJSON `json:"-"`
}

// currencyCalculateVatPriceResponseJSON contains the JSON metadata for the struct
// [CurrencyCalculateVatPriceResponse]
type currencyCalculateVatPriceResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrencyCalculateVatPriceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currencyCalculateVatPriceResponseJSON) RawJSON() string {
	return r.raw
}

type CurrencyConvertResponse struct {
	Code    int64                       `json:"code"`
	Data    CurrencyConvertResponseData `json:"data"`
	Success bool                        `json:"success"`
	JSON    currencyConvertResponseJSON `json:"-"`
}

// currencyConvertResponseJSON contains the JSON metadata for the struct
// [CurrencyConvertResponse]
type currencyConvertResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrencyConvertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currencyConvertResponseJSON) RawJSON() string {
	return r.raw
}

type CurrencyConvertResponseData struct {
	// The original amount.
	Amount float64 `json:"amount"`
	// The converted amount.
	Converted float64                           `json:"converted"`
	From      string                            `json:"from"`
	Object    CurrencyConvertResponseDataObject `json:"object"`
	// The exchange rate used.
	Rate float64                         `json:"rate"`
	To   string                          `json:"to"`
	JSON currencyConvertResponseDataJSON `json:"-"`
}

// currencyConvertResponseDataJSON contains the JSON metadata for the struct
// [CurrencyConvertResponseData]
type currencyConvertResponseDataJSON struct {
	Amount      apijson.Field
	Converted   apijson.Field
	From        apijson.Field
	Object      apijson.Field
	Rate        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrencyConvertResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currencyConvertResponseDataJSON) RawJSON() string {
	return r.raw
}

type CurrencyConvertResponseDataObject string

const (
	CurrencyConvertResponseDataObjectConversion CurrencyConvertResponseDataObject = "conversion"
)

func (r CurrencyConvertResponseDataObject) IsKnown() bool {
	switch r {
	case CurrencyConvertResponseDataObjectConversion:
		return true
	}
	return false
}

type CurrencyListParams struct {
	// The 3-character currency code(s) to convert from (e.g. "USD", "CAD"). Can be a
	// comma-separated list without spaces (e.g. "USD,CAD,AUD").
	From param.Field[string] `query:"from"`
	// The 3-character target currency code. Must be either "GBP" or "EUR".
	To param.Field[CurrencyListParamsTo] `query:"to"`
}

// URLQuery serializes [CurrencyListParams]'s query parameters as `url.Values`.
func (r CurrencyListParams) URLQuery() (v url.Values) {
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

func (r CurrencyListParamsTo) IsKnown() bool {
	switch r {
	case CurrencyListParamsToGbp, CurrencyListParamsToEur:
		return true
	}
	return false
}

type CurrencyCalculateVatPriceParams struct {
	// The price to calculate on. Must be a string with exactly 2 decimal places (e.g.
	// "30.00", "59.95").
	Price param.Field[string] `query:"price" api:"required"`
	// Whether the provided price is inclusive or exclusive of VAT.
	TaxType param.Field[CurrencyCalculateVatPriceParamsTaxType] `query:"tax_type" api:"required"`
	// A percentage VAT rate to use for the calculation.
	VatRate param.Field[float64] `query:"vat_rate" api:"required"`
}

// URLQuery serializes [CurrencyCalculateVatPriceParams]'s query parameters as
// `url.Values`.
func (r CurrencyCalculateVatPriceParams) URLQuery() (v url.Values) {
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

func (r CurrencyCalculateVatPriceParamsTaxType) IsKnown() bool {
	switch r {
	case CurrencyCalculateVatPriceParamsTaxTypeIncl, CurrencyCalculateVatPriceParamsTaxTypeExcl:
		return true
	}
	return false
}

type CurrencyConvertParams struct {
	// The amount to convert. Must be a string with exactly 2 decimal places (e.g.
	// "39.99").
	Amount param.Field[string] `query:"amount" api:"required"`
	// The 3-character source currency code (e.g. "USD", "CAD").
	From param.Field[string] `query:"from" api:"required"`
	// The 3-character target currency code. Must be either "GBP" or "EUR".
	To param.Field[CurrencyConvertParamsTo] `query:"to" api:"required"`
}

// URLQuery serializes [CurrencyConvertParams]'s query parameters as `url.Values`.
func (r CurrencyConvertParams) URLQuery() (v url.Values) {
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

func (r CurrencyConvertParamsTo) IsKnown() bool {
	switch r {
	case CurrencyConvertParamsToGbp, CurrencyConvertParamsToEur:
		return true
	}
	return false
}
