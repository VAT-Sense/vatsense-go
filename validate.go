// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"encoding/json"
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

// VAT and EORI number validation
//
// ValidateService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewValidateService] method instead.
type ValidateService struct {
	options []option.RequestOption
}

// NewValidateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewValidateService(opts ...option.RequestOption) (r ValidateService) {
	r = ValidateService{}
	r.options = opts
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
// `412` error and the request does not count against your usage quota. It also
// returns the `412`, uncounted, when a check takes longer than 55 seconds for a
// number checked with VIES or the EU EORI service, or 25 seconds for any other.
// With a webhook endpoint taking `validation.completed`, a number that got the
// `412` keeps being checked and the answer is sent by webhook
// (https://vatsense.com/documentation#webhooks).
//
// Provide either `vat_number` or `eori_number`, but not both.
func (r *ValidateService) Check(ctx context.Context, query ValidateCheckParams, opts ...option.RequestOption) (res *ValidateCheckResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ValidateCheckResponse struct {
	Code    int64                     `json:"code"`
	Data    ValidateCheckResponseData `json:"data"`
	Success bool                      `json:"success"`
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
func (r ValidateCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *ValidateCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateCheckResponseData struct {
	Company ValidateCheckResponseDataCompanyUnion `json:"company"`
	// Official consultation number (only returned when requester_vat_number is
	// provided).
	ConsultationNumber string `json:"consultation_number" api:"nullable"`
	// Only when a `match_` parameter was sent.
	Match ValidateCheckResponseDataMatch `json:"match"`
	// Whether the VAT/EORI number is valid.
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company            respjson.Field
		ConsultationNumber respjson.Field
		Match              respjson.Field
		Valid              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValidateCheckResponseData) RawJSON() string { return r.JSON.raw }
func (r *ValidateCheckResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only when a `match_` parameter was sent.
type ValidateCheckResponseDataMatch struct {
	// One entry for each detail sent (company_name, address, street, postcode,
	// city).
	Fields map[string]ValidateCheckResponseDataMatchField `json:"fields"`
	// likely (a score of 80 or more, and the name, if sent, is not a no_match),
	// possible (50 to 79), unlikely (under 50), or not_applicable when nothing could
	// be compared.
	//
	// Any of "likely", "possible", "unlikely", "not_applicable".
	Result string `json:"result"`
	// The average of the details sent, with the name counting twice. null when
	// not_applicable.
	Score int64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Result      respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValidateCheckResponseDataMatch) RawJSON() string { return r.JSON.raw }
func (r *ValidateCheckResponseDataMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How one detail you sent compared with the registered details.
type ValidateCheckResponseDataMatchField struct {
	// Any of "match", "partial", "no_match", "not_applicable".
	Result string `json:"result"`
	Score  int64  `json:"score" api:"nullable"`
	// VIES's own answer for that detail on a live EU check; empty when VIES said
	// nothing.
	//
	// Any of "valid", "invalid", "not_processed".
	Vies string `json:"vies" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Score       respjson.Field
		Vies        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValidateCheckResponseDataMatchField) RawJSON() string { return r.JSON.raw }
func (r *ValidateCheckResponseDataMatchField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ValidateCheckResponseDataCompanyUnion contains all possible properties and
// values from [ValidateCheckResponseDataCompanyValidationCompany],
// [ValidateCheckResponseDataCompanyEoriValidationCompany].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ValidateCheckResponseDataCompanyUnion struct {
	CompanyAddress string `json:"company_address"`
	CompanyName    string `json:"company_name"`
	CountryCode    string `json:"country_code"`
	// This field is from variant [ValidateCheckResponseDataCompanyValidationCompany].
	VatNumber string `json:"vat_number"`
	// This field is from variant
	// [ValidateCheckResponseDataCompanyEoriValidationCompany].
	EoriNumber string `json:"eori_number"`
	JSON       struct {
		CompanyAddress respjson.Field
		CompanyName    respjson.Field
		CountryCode    respjson.Field
		VatNumber      respjson.Field
		EoriNumber     respjson.Field
		raw            string
	} `json:"-"`
}

func (u ValidateCheckResponseDataCompanyUnion) AsValidateCheckResponseDataCompanyValidationCompany() (v ValidateCheckResponseDataCompanyValidationCompany) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ValidateCheckResponseDataCompanyUnion) AsValidateCheckResponseDataCompanyEoriValidationCompany() (v ValidateCheckResponseDataCompanyEoriValidationCompany) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ValidateCheckResponseDataCompanyUnion) RawJSON() string { return u.JSON.raw }

func (r *ValidateCheckResponseDataCompanyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateCheckResponseDataCompanyValidationCompany struct {
	CompanyAddress string `json:"company_address"`
	CompanyName    string `json:"company_name"`
	CountryCode    string `json:"country_code"`
	// The VAT number (without country code prefix).
	VatNumber string `json:"vat_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyAddress respjson.Field
		CompanyName    respjson.Field
		CountryCode    respjson.Field
		VatNumber      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValidateCheckResponseDataCompanyValidationCompany) RawJSON() string { return r.JSON.raw }
func (r *ValidateCheckResponseDataCompanyValidationCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateCheckResponseDataCompanyEoriValidationCompany struct {
	CompanyAddress string `json:"company_address"`
	CompanyName    string `json:"company_name"`
	CountryCode    string `json:"country_code"`
	// The EORI number (without country code prefix).
	EoriNumber string `json:"eori_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyAddress respjson.Field
		CompanyName    respjson.Field
		CountryCode    respjson.Field
		EoriNumber     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValidateCheckResponseDataCompanyEoriValidationCompany) RawJSON() string { return r.JSON.raw }
func (r *ValidateCheckResponseDataCompanyEoriValidationCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ValidateCheckParams struct {
	// The EORI number to validate. Must include the leading 2-character country code
	// (e.g. "GB123456789123"). UK and EU only.
	EoriNumber param.Opt[string] `query:"eori_number,omitzero" json:"-"`
	// The whole address you hold, as one line, matched against the registered
	// address.
	MatchAddress param.Opt[string] `query:"match_address,omitzero" json:"-"`
	// The town or city you hold, looked for in the registered address.
	MatchCity param.Opt[string] `query:"match_city,omitzero" json:"-"`
	// The company name you hold, matched against the registered name. Legal forms
	// such as Ltd or GmbH, accents, punctuation and word order are ignored. Sending
	// any `match_` parameter adds `match` to the response; matching costs nothing
	// extra.
	MatchCompanyName param.Opt[string] `query:"match_company_name,omitzero" json:"-"`
	// The postcode you hold, looked for in the registered address. Spaces are
	// ignored.
	MatchPostcode param.Opt[string] `query:"match_postcode,omitzero" json:"-"`
	// The street line you hold, looked for in the registered address.
	MatchStreet param.Opt[string] `query:"match_street,omitzero" json:"-"`
	// Your own VAT number. If supplied, the response will include a unique
	// consultation number issued by the relevant authority (VIES or HMRC). Must
	// include the leading 2-character country code.
	//
	// Note: GB requester numbers only work for GB validations; EU requester numbers
	// only work for EU validations. Cross-region is not supported.
	RequesterVatNumber param.Opt[string] `query:"requester_vat_number,omitzero" json:"-"`
	// The VAT number to validate. Must include the leading 2-character country code
	// (e.g. "GB288305674", "FR12345678901").
	VatNumber param.Opt[string] `query:"vat_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ValidateCheckParams]'s query parameters as `url.Values`.
func (r ValidateCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
