// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/vat-sense-go/internal/apijson"
	"github.com/stainless-sdks/vat-sense-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/vat-sense-go/internal/encoding/json"
	"github.com/stainless-sdks/vat-sense-go/internal/requestconfig"
	"github.com/stainless-sdks/vat-sense-go/option"
	"github.com/stainless-sdks/vat-sense-go/packages/param"
	"github.com/stainless-sdks/vat-sense-go/packages/respjson"
)

// VAT-compliant invoice management
//
// InvoiceService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	options []option.RequestOption
	// VAT-compliant invoice management
	Item InvoiceItemService
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r InvoiceService) {
	r = InvoiceService{}
	r.options = opts
	r.Item = NewInvoiceItemService(opts...)
	return
}

// Create a new VAT-compliant invoice. VAT Sense will automatically calculate the
// totals based on the items provided.
//
// Not available with sandbox API keys.
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "invoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a specific invoice by its ID.
func (r *InvoiceService) Get(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s", url.PathEscape(invoiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing invoice. Only the fields provided will be updated.
func (r *InvoiceService) Update(ctx context.Context, invoiceID string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s", url.PathEscape(invoiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of all invoices.
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *InvoiceListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "invoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently delete an invoice.
func (r *InvoiceService) Delete(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *InvoiceDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s", url.PathEscape(invoiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The properties Business, CurrencyCode, Date, Items, TaxPoint are required.
type CreateInvoiceParam struct {
	Business InvoiceBusinessInputParam `json:"business,omitzero" api:"required"`
	// The 3-character currency code the invoice is billed in.
	CurrencyCode string `json:"currency_code" api:"required"`
	// The date the invoice was issued (YYYY-MM-DD or YYYY-MM-DD HH:MM:SS).
	Date  string                  `json:"date" api:"required"`
	Items []InvoiceItemInputParam `json:"items,omitzero" api:"required"`
	// The tax point or "time of supply" (YYYY-MM-DD or YYYY-MM-DD HH:MM:SS).
	TaxPoint string `json:"tax_point" api:"required"`
	// Whether the invoice is subject to VAT.
	HasVat param.Opt[bool] `json:"has_vat,omitzero"`
	// A unique invoice number. If not provided, defaults to an auto-incremented
	// number.
	InvoiceNumber param.Opt[string] `json:"invoice_number,omitzero"`
	// Whether the invoice is a copy of a primary invoice.
	IsCopy param.Opt[bool] `json:"is_copy,omitzero"`
	// Whether the invoice is zero-rated due to reverse charge.
	IsReverseCharge param.Opt[bool] `json:"is_reverse_charge,omitzero"`
	// Any additional notes for the invoice.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Pad the auto-generated invoice number with leading zeros to this length.
	PadInvoiceNumber param.Opt[int64] `json:"pad_invoice_number,omitzero"`
	// A serial prepended to the auto-generated invoice number. Each unique serial has
	// its own auto-increment range.
	Serial param.Opt[string] `json:"serial,omitzero"`
	// Whether the invoice has been zero-rated.
	ZeroRated  param.Opt[bool]             `json:"zero_rated,omitzero"`
	Conversion InvoiceConversionInputParam `json:"conversion,omitzero"`
	Customer   InvoiceCustomerInputParam   `json:"customer,omitzero"`
	// Whether item prices include or exclude VAT.
	//
	// Any of "incl", "excl".
	TaxType CreateInvoiceTaxType `json:"tax_type,omitzero"`
	// The type of invoice.
	//
	// Any of "sale", "refund".
	Type CreateInvoiceType `json:"type,omitzero"`
	paramObj
}

func (r CreateInvoiceParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateInvoiceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateInvoiceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether item prices include or exclude VAT.
type CreateInvoiceTaxType string

const (
	CreateInvoiceTaxTypeIncl CreateInvoiceTaxType = "incl"
	CreateInvoiceTaxTypeExcl CreateInvoiceTaxType = "excl"
)

// The type of invoice.
type CreateInvoiceType string

const (
	CreateInvoiceTypeSale   CreateInvoiceType = "sale"
	CreateInvoiceTypeRefund CreateInvoiceType = "refund"
)

type Invoice struct {
	ID            string                 `json:"id"`
	Business      InvoiceBusiness        `json:"business"`
	Conversion    InvoiceConversionInput `json:"conversion" api:"nullable"`
	Created       time.Time              `json:"created" format:"date-time"`
	CurrencyCode  string                 `json:"currency_code"`
	Customer      InvoiceCustomer        `json:"customer" api:"nullable"`
	Date          string                 `json:"date"`
	HasVat        bool                   `json:"has_vat"`
	InvoiceNumber string                 `json:"invoice_number"`
	// Unique URL to view the invoice. Append "/pdf" to download a PDF copy.
	InvoiceURL      string        `json:"invoice_url" format:"uri"`
	IsCopy          bool          `json:"is_copy"`
	IsReverseCharge bool          `json:"is_reverse_charge"`
	Items           []InvoiceItem `json:"items"`
	Notes           string        `json:"notes" api:"nullable"`
	NumItems        int64         `json:"num_items"`
	// Any of "invoice".
	Object   InvoiceObject `json:"object"`
	TaxPoint string        `json:"tax_point"`
	// Any of "incl", "excl".
	TaxType InvoiceTaxType `json:"tax_type"`
	Totals  InvoiceTotals  `json:"totals"`
	// Any of "sale", "refund".
	Type      InvoiceType `json:"type"`
	Updated   time.Time   `json:"updated" format:"date-time"`
	ZeroRated bool        `json:"zero_rated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Business        respjson.Field
		Conversion      respjson.Field
		Created         respjson.Field
		CurrencyCode    respjson.Field
		Customer        respjson.Field
		Date            respjson.Field
		HasVat          respjson.Field
		InvoiceNumber   respjson.Field
		InvoiceURL      respjson.Field
		IsCopy          respjson.Field
		IsReverseCharge respjson.Field
		Items           respjson.Field
		Notes           respjson.Field
		NumItems        respjson.Field
		Object          respjson.Field
		TaxPoint        respjson.Field
		TaxType         respjson.Field
		Totals          respjson.Field
		Type            respjson.Field
		Updated         respjson.Field
		ZeroRated       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invoice) RawJSON() string { return r.JSON.raw }
func (r *Invoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceBusiness struct {
	Address       string `json:"address"`
	CompanyNumber string `json:"company_number"`
	Logo          string `json:"logo" api:"nullable"`
	Name          string `json:"name"`
	VatNumber     string `json:"vat_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address       respjson.Field
		CompanyNumber respjson.Field
		Logo          respjson.Field
		Name          respjson.Field
		VatNumber     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceBusiness) RawJSON() string { return r.JSON.raw }
func (r *InvoiceBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCustomer struct {
	Address       string `json:"address"`
	CompanyNumber string `json:"company_number"`
	Logo          string `json:"logo" api:"nullable"`
	Name          string `json:"name"`
	VatNumber     string `json:"vat_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address       respjson.Field
		CompanyNumber respjson.Field
		Logo          respjson.Field
		Name          respjson.Field
		VatNumber     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceCustomer) RawJSON() string { return r.JSON.raw }
func (r *InvoiceCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceObject string

const (
	InvoiceObjectInvoice InvoiceObject = "invoice"
)

type InvoiceTaxType string

const (
	InvoiceTaxTypeIncl InvoiceTaxType = "incl"
	InvoiceTaxTypeExcl InvoiceTaxType = "excl"
)

type InvoiceTotals struct {
	// Total discount amount.
	Discount float64 `json:"discount"`
	// Total before VAT.
	Subtotal float64 `json:"subtotal"`
	// Grand total.
	Total float64 `json:"total"`
	// Total VAT amount.
	Vat float64 `json:"vat"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Discount    respjson.Field
		Subtotal    respjson.Field
		Total       respjson.Field
		Vat         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceTotals) RawJSON() string { return r.JSON.raw }
func (r *InvoiceTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceType string

const (
	InvoiceTypeSale   InvoiceType = "sale"
	InvoiceTypeRefund InvoiceType = "refund"
)

// The properties Address, Name, VatNumber are required.
type InvoiceBusinessInputParam struct {
	// Your business trading address.
	Address string `json:"address" api:"required"`
	// Your business trading name.
	Name string `json:"name" api:"required"`
	// Your business VAT number.
	VatNumber   string            `json:"vat_number" api:"required"`
	BankAccount param.Opt[string] `json:"bank_account,omitzero"`
	// Your business company number.
	CompanyNumber param.Opt[string] `json:"company_number,omitzero"`
	Email         param.Opt[string] `json:"email,omitzero" format:"email"`
	// URL to your company logo (HTTPS only, .svg/.jpg/.png). Recommended 240px by
	// 60px.
	Logo    param.Opt[string] `json:"logo,omitzero" format:"uri"`
	Phone   param.Opt[string] `json:"phone,omitzero"`
	Website param.Opt[string] `json:"website,omitzero" format:"uri"`
	paramObj
}

func (r InvoiceBusinessInputParam) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceBusinessInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceBusinessInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceConversionInput struct {
	// The 3-character currency code for the conversion.
	CurrencyCode string `json:"currency_code" api:"required"`
	// The rate of conversion.
	Rate float64 `json:"rate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode respjson.Field
		Rate         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceConversionInput) RawJSON() string { return r.JSON.raw }
func (r *InvoiceConversionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InvoiceConversionInput to a InvoiceConversionInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InvoiceConversionInputParam.Overrides()
func (r InvoiceConversionInput) ToParam() InvoiceConversionInputParam {
	return param.Override[InvoiceConversionInputParam](json.RawMessage(r.RawJSON()))
}

// The properties CurrencyCode, Rate are required.
type InvoiceConversionInputParam struct {
	// The 3-character currency code for the conversion.
	CurrencyCode string `json:"currency_code" api:"required"`
	// The rate of conversion.
	Rate float64 `json:"rate" api:"required"`
	paramObj
}

func (r InvoiceConversionInputParam) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceConversionInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceConversionInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type InvoiceCustomerInputParam struct {
	// The customer's trading name.
	Name          string            `json:"name" api:"required"`
	Address       param.Opt[string] `json:"address,omitzero"`
	CompanyNumber param.Opt[string] `json:"company_number,omitzero"`
	CountryCode   param.Opt[string] `json:"country_code,omitzero"`
	Email         param.Opt[string] `json:"email,omitzero" format:"email"`
	// URL to the customer logo (HTTPS only, .jpg/.png).
	Logo      param.Opt[string] `json:"logo,omitzero" format:"uri"`
	VatNumber param.Opt[string] `json:"vat_number,omitzero"`
	paramObj
}

func (r InvoiceCustomerInputParam) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceCustomerInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceCustomerInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceResponse struct {
	Code    int64   `json:"code"`
	Data    Invoice `json:"data"`
	Success bool    `json:"success"`
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
func (r InvoiceResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListResponse struct {
	Code    int64     `json:"code"`
	Data    []Invoice `json:"data"`
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
func (r InvoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceDeleteResponse struct {
	Code    int64 `json:"code" api:"required"`
	Success bool  `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceNewParams struct {
	CreateInvoice CreateInvoiceParam
	paramObj
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateInvoice)
}
func (r *InvoiceNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateInvoice)
}

type InvoiceUpdateParams struct {
	CreateInvoice CreateInvoiceParam
	paramObj
}

func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateInvoice)
}
func (r *InvoiceUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateInvoice)
}

type InvoiceListParams struct {
	// Number of invoices to return (default 10, max 100).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of invoices to skip (default 0).
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Search query to filter invoices.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
