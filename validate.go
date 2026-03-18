// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/apiquery"
	"github.com/VAT-Sense/vatsense-go/internal/param"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
	"github.com/tidwall/gjson"
)

// VAT and EORI number validation
//
// ValidateService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewValidateService] method instead.
type ValidateService struct {
	Options []option.RequestOption
}

// NewValidateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewValidateService(opts ...option.RequestOption) (r *ValidateService) {
	r = &ValidateService{}
	r.Options = opts
	return
}

// Check whether a given VAT number or EORI number is valid against live government
// records.
//
// **VAT validation** checks against UK (HMRC), EU (VIES), Australia, Norway,
// Switzerland, South Africa, and Brazil records.
//
// **EORI validation** checks against UK and EU records only.
//
// If the external validation service is temporarily unavailable, the API returns a
// `412` error and the request does not count against your usage quota.
//
// Provide either `vat_number` or `eori_number`, but not both.
func (r *ValidateService) Check(ctx context.Context, query ValidateCheckParams, opts ...option.RequestOption) (res *ValidateCheckResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ValidateCheckResponse struct {
	Code    int64                     `json:"code"`
	Data    ValidateCheckResponseData `json:"data"`
	Success bool                      `json:"success"`
	JSON    validateCheckResponseJSON `json:"-"`
}

// validateCheckResponseJSON contains the JSON metadata for the struct
// [ValidateCheckResponse]
type validateCheckResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateCheckResponseJSON) RawJSON() string {
	return r.raw
}

type ValidateCheckResponseData struct {
	Company ValidateCheckResponseDataCompany `json:"company"`
	// Official consultation number (only returned when requester_vat_number is
	// provided).
	ConsultationNumber string `json:"consultation_number" api:"nullable"`
	// Whether the VAT/EORI number is valid.
	Valid bool                          `json:"valid"`
	JSON  validateCheckResponseDataJSON `json:"-"`
}

// validateCheckResponseDataJSON contains the JSON metadata for the struct
// [ValidateCheckResponseData]
type validateCheckResponseDataJSON struct {
	Company            apijson.Field
	ConsultationNumber apijson.Field
	Valid              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ValidateCheckResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateCheckResponseDataJSON) RawJSON() string {
	return r.raw
}

type ValidateCheckResponseDataCompany struct {
	CompanyAddress string `json:"company_address"`
	CompanyName    string `json:"company_name"`
	CountryCode    string `json:"country_code"`
	// The EORI number (without country code prefix).
	EoriNumber string `json:"eori_number"`
	// The VAT number (without country code prefix).
	VatNumber string                               `json:"vat_number"`
	JSON      validateCheckResponseDataCompanyJSON `json:"-"`
	union     ValidateCheckResponseDataCompanyUnion
}

// validateCheckResponseDataCompanyJSON contains the JSON metadata for the struct
// [ValidateCheckResponseDataCompany]
type validateCheckResponseDataCompanyJSON struct {
	CompanyAddress apijson.Field
	CompanyName    apijson.Field
	CountryCode    apijson.Field
	EoriNumber     apijson.Field
	VatNumber      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r validateCheckResponseDataCompanyJSON) RawJSON() string {
	return r.raw
}

func (r *ValidateCheckResponseDataCompany) UnmarshalJSON(data []byte) (err error) {
	*r = ValidateCheckResponseDataCompany{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ValidateCheckResponseDataCompanyUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ValidateCheckResponseDataCompanyValidationCompany],
// [ValidateCheckResponseDataCompanyEoriValidationCompany].
func (r ValidateCheckResponseDataCompany) AsUnion() ValidateCheckResponseDataCompanyUnion {
	return r.union
}

// Union satisfied by [ValidateCheckResponseDataCompanyValidationCompany] or
// [ValidateCheckResponseDataCompanyEoriValidationCompany].
type ValidateCheckResponseDataCompanyUnion interface {
	implementsValidateCheckResponseDataCompany()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ValidateCheckResponseDataCompanyUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ValidateCheckResponseDataCompanyValidationCompany{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ValidateCheckResponseDataCompanyEoriValidationCompany{}),
		},
	)
}

type ValidateCheckResponseDataCompanyValidationCompany struct {
	CompanyAddress string `json:"company_address"`
	CompanyName    string `json:"company_name"`
	CountryCode    string `json:"country_code"`
	// The VAT number (without country code prefix).
	VatNumber string                                                `json:"vat_number"`
	JSON      validateCheckResponseDataCompanyValidationCompanyJSON `json:"-"`
}

// validateCheckResponseDataCompanyValidationCompanyJSON contains the JSON metadata
// for the struct [ValidateCheckResponseDataCompanyValidationCompany]
type validateCheckResponseDataCompanyValidationCompanyJSON struct {
	CompanyAddress apijson.Field
	CompanyName    apijson.Field
	CountryCode    apijson.Field
	VatNumber      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ValidateCheckResponseDataCompanyValidationCompany) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateCheckResponseDataCompanyValidationCompanyJSON) RawJSON() string {
	return r.raw
}

func (r ValidateCheckResponseDataCompanyValidationCompany) implementsValidateCheckResponseDataCompany() {
}

type ValidateCheckResponseDataCompanyEoriValidationCompany struct {
	CompanyAddress string `json:"company_address"`
	CompanyName    string `json:"company_name"`
	CountryCode    string `json:"country_code"`
	// The EORI number (without country code prefix).
	EoriNumber string                                                    `json:"eori_number"`
	JSON       validateCheckResponseDataCompanyEoriValidationCompanyJSON `json:"-"`
}

// validateCheckResponseDataCompanyEoriValidationCompanyJSON contains the JSON
// metadata for the struct [ValidateCheckResponseDataCompanyEoriValidationCompany]
type validateCheckResponseDataCompanyEoriValidationCompanyJSON struct {
	CompanyAddress apijson.Field
	CompanyName    apijson.Field
	CountryCode    apijson.Field
	EoriNumber     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ValidateCheckResponseDataCompanyEoriValidationCompany) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateCheckResponseDataCompanyEoriValidationCompanyJSON) RawJSON() string {
	return r.raw
}

func (r ValidateCheckResponseDataCompanyEoriValidationCompany) implementsValidateCheckResponseDataCompany() {
}

type ValidateCheckParams struct {
	// The EORI number to validate. Must include the leading 2-character country code
	// (e.g. "GB123456789123"). UK and EU only.
	EoriNumber param.Field[string] `query:"eori_number"`
	// Your own VAT number. If supplied, the response will include a unique
	// consultation number issued by the relevant authority (VIES or HMRC). Must
	// include the leading 2-character country code.
	//
	// Note: GB requester numbers only work for GB validations; EU requester numbers
	// only work for EU validations. Cross-region is not supported.
	RequesterVatNumber param.Field[string] `query:"requester_vat_number"`
	// The VAT number to validate. Must include the leading 2-character country code
	// (e.g. "GB288305674", "FR12345678901").
	VatNumber param.Field[string] `query:"vat_number"`
}

// URLQuery serializes [ValidateCheckParams]'s query parameters as `url.Values`.
func (r ValidateCheckParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
