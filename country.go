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

// Country and province information
//
// CountryService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCountryService] method instead.
type CountryService struct {
	Options []option.RequestOption
}

// NewCountryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCountryService(opts ...option.RequestOption) (r *CountryService) {
	r = &CountryService{}
	r.Options = opts
	return
}

// Returns a list of all countries, including whether they are subject to VAT/GST
// and whether they are subject to EU VAT. Each country is returned as a country
// object.
//
// You can optionally filter by country code or IP address.
func (r *CountryService) List(ctx context.Context, query CountryListParams, opts ...option.RequestOption) (res *CountryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a list of all provinces within a given country.
func (r *CountryService) ListProvinces(ctx context.Context, query CountryListProvincesParams, opts ...option.RequestOption) (res *CountryListProvincesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "countries/provinces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Country struct {
	// 2-character ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	// Whether the country is subject to EU VAT.
	Eu        bool          `json:"eu"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Object    CountryObject `json:"object"`
	// Whether the country is subject to VAT/GST.
	Vat  bool        `json:"vat"`
	JSON countryJSON `json:"-"`
}

// countryJSON contains the JSON metadata for the struct [Country]
type countryJSON struct {
	CountryCode apijson.Field
	CountryName apijson.Field
	Eu          apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Object      apijson.Field
	Vat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Country) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryJSON) RawJSON() string {
	return r.raw
}

type CountryObject string

const (
	CountryObjectCountry CountryObject = "country"
)

func (r CountryObject) IsKnown() bool {
	switch r {
	case CountryObjectCountry:
		return true
	}
	return false
}

type CountryListResponse struct {
	Code    int64                   `json:"code"`
	Data    []Country               `json:"data"`
	Success bool                    `json:"success"`
	JSON    countryListResponseJSON `json:"-"`
}

// countryListResponseJSON contains the JSON metadata for the struct
// [CountryListResponse]
type countryListResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CountryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryListResponseJSON) RawJSON() string {
	return r.raw
}

type CountryListProvincesResponse struct {
	Code    int64                              `json:"code"`
	Data    []CountryListProvincesResponseData `json:"data"`
	Success bool                               `json:"success"`
	JSON    countryListProvincesResponseJSON   `json:"-"`
}

// countryListProvincesResponseJSON contains the JSON metadata for the struct
// [CountryListProvincesResponse]
type countryListProvincesResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CountryListProvincesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryListProvincesResponseJSON) RawJSON() string {
	return r.raw
}

type CountryListProvincesResponseData struct {
	CountryCode  string                                 `json:"country_code"`
	Object       CountryListProvincesResponseDataObject `json:"object"`
	ProvinceCode string                                 `json:"province_code"`
	ProvinceName string                                 `json:"province_name"`
	JSON         countryListProvincesResponseDataJSON   `json:"-"`
}

// countryListProvincesResponseDataJSON contains the JSON metadata for the struct
// [CountryListProvincesResponseData]
type countryListProvincesResponseDataJSON struct {
	CountryCode  apijson.Field
	Object       apijson.Field
	ProvinceCode apijson.Field
	ProvinceName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CountryListProvincesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryListProvincesResponseDataJSON) RawJSON() string {
	return r.raw
}

type CountryListProvincesResponseDataObject string

const (
	CountryListProvincesResponseDataObjectProvince CountryListProvincesResponseDataObject = "province"
)

func (r CountryListProvincesResponseDataObject) IsKnown() bool {
	switch r {
	case CountryListProvincesResponseDataObjectProvince:
		return true
	}
	return false
}

type CountryListParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "GB", "FR").
	CountryCode param.Field[string] `query:"country_code"`
	// An IPv4 or IPv6 address. If provided, the country will be determined from the IP
	// address.
	IPAddress param.Field[string] `query:"ip_address"`
}

// URLQuery serializes [CountryListParams]'s query parameters as `url.Values`.
func (r CountryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CountryListProvincesParams struct {
	// A 2-character ISO 3166-1 alpha-2 country code (e.g. "CA").
	CountryCode param.Field[string] `query:"country_code" api:"required"`
}

// URLQuery serializes [CountryListProvincesParams]'s query parameters as
// `url.Values`.
func (r CountryListProvincesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
