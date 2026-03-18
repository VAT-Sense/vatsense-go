// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/apiquery"
	"github.com/VAT-Sense/vatsense-go/internal/param"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
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
	Options []option.RequestOption
	// VAT-compliant invoice management
	Item *InvoiceItemService
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	r.Item = NewInvoiceItemService(opts...)
	return
}

// Create a new VAT-compliant invoice. VAT Sense will automatically calculate the
// totals based on the items provided.
//
// Not available with sandbox API keys.
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "invoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a specific invoice by its ID.
func (r *InvoiceService) Get(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing invoice. Only the fields provided will be updated.
func (r *InvoiceService) Update(ctx context.Context, invoiceID string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of all invoices.
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *InvoiceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "invoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently delete an invoice.
func (r *InvoiceService) Delete(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *InvoiceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type CreateInvoiceParam struct {
	Business param.Field[InvoiceBusinessInputParam] `json:"business" api:"required"`
	// The 3-character currency code the invoice is billed in.
	CurrencyCode param.Field[string] `json:"currency_code" api:"required"`
	// The date the invoice was issued (YYYY-MM-DD or YYYY-MM-DD HH:MM:SS).
	Date  param.Field[string]                  `json:"date" api:"required"`
	Items param.Field[[]InvoiceItemInputParam] `json:"items" api:"required"`
	// The tax point or "time of supply" (YYYY-MM-DD or YYYY-MM-DD HH:MM:SS).
	TaxPoint   param.Field[string]                      `json:"tax_point" api:"required"`
	Conversion param.Field[InvoiceConversionInputParam] `json:"conversion"`
	Customer   param.Field[InvoiceCustomerInputParam]   `json:"customer"`
	// Whether the invoice is subject to VAT.
	HasVat param.Field[bool] `json:"has_vat"`
	// A unique invoice number. If not provided, defaults to an auto-incremented
	// number.
	InvoiceNumber param.Field[string] `json:"invoice_number"`
	// Whether the invoice is a copy of a primary invoice.
	IsCopy param.Field[bool] `json:"is_copy"`
	// Whether the invoice is zero-rated due to reverse charge.
	IsReverseCharge param.Field[bool] `json:"is_reverse_charge"`
	// Any additional notes for the invoice.
	Notes param.Field[string] `json:"notes"`
	// Pad the auto-generated invoice number with leading zeros to this length.
	PadInvoiceNumber param.Field[int64] `json:"pad_invoice_number"`
	// A serial prepended to the auto-generated invoice number. Each unique serial has
	// its own auto-increment range.
	Serial param.Field[string] `json:"serial"`
	// Whether item prices include or exclude VAT.
	TaxType param.Field[CreateInvoiceTaxType] `json:"tax_type"`
	// The type of invoice.
	Type param.Field[CreateInvoiceType] `json:"type"`
	// Whether the invoice has been zero-rated.
	ZeroRated param.Field[bool] `json:"zero_rated"`
}

func (r CreateInvoiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether item prices include or exclude VAT.
type CreateInvoiceTaxType string

const (
	CreateInvoiceTaxTypeIncl CreateInvoiceTaxType = "incl"
	CreateInvoiceTaxTypeExcl CreateInvoiceTaxType = "excl"
)

func (r CreateInvoiceTaxType) IsKnown() bool {
	switch r {
	case CreateInvoiceTaxTypeIncl, CreateInvoiceTaxTypeExcl:
		return true
	}
	return false
}

// The type of invoice.
type CreateInvoiceType string

const (
	CreateInvoiceTypeSale   CreateInvoiceType = "sale"
	CreateInvoiceTypeRefund CreateInvoiceType = "refund"
)

func (r CreateInvoiceType) IsKnown() bool {
	switch r {
	case CreateInvoiceTypeSale, CreateInvoiceTypeRefund:
		return true
	}
	return false
}

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
	InvoiceURL      string         `json:"invoice_url" format:"uri"`
	IsCopy          bool           `json:"is_copy"`
	IsReverseCharge bool           `json:"is_reverse_charge"`
	Items           []InvoiceItem  `json:"items"`
	Notes           string         `json:"notes" api:"nullable"`
	NumItems        int64          `json:"num_items"`
	Object          InvoiceObject  `json:"object"`
	TaxPoint        string         `json:"tax_point"`
	TaxType         InvoiceTaxType `json:"tax_type"`
	Totals          InvoiceTotals  `json:"totals"`
	Type            InvoiceType    `json:"type"`
	Updated         time.Time      `json:"updated" format:"date-time"`
	ZeroRated       bool           `json:"zero_rated"`
	JSON            invoiceJSON    `json:"-"`
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	ID              apijson.Field
	Business        apijson.Field
	Conversion      apijson.Field
	Created         apijson.Field
	CurrencyCode    apijson.Field
	Customer        apijson.Field
	Date            apijson.Field
	HasVat          apijson.Field
	InvoiceNumber   apijson.Field
	InvoiceURL      apijson.Field
	IsCopy          apijson.Field
	IsReverseCharge apijson.Field
	Items           apijson.Field
	Notes           apijson.Field
	NumItems        apijson.Field
	Object          apijson.Field
	TaxPoint        apijson.Field
	TaxType         apijson.Field
	Totals          apijson.Field
	Type            apijson.Field
	Updated         apijson.Field
	ZeroRated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceBusiness struct {
	Address       string              `json:"address"`
	CompanyNumber string              `json:"company_number"`
	Logo          string              `json:"logo" api:"nullable"`
	Name          string              `json:"name"`
	VatNumber     string              `json:"vat_number"`
	JSON          invoiceBusinessJSON `json:"-"`
}

// invoiceBusinessJSON contains the JSON metadata for the struct [InvoiceBusiness]
type invoiceBusinessJSON struct {
	Address       apijson.Field
	CompanyNumber apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	VatNumber     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceBusiness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceBusinessJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomer struct {
	Address       string              `json:"address"`
	CompanyNumber string              `json:"company_number"`
	Logo          string              `json:"logo" api:"nullable"`
	Name          string              `json:"name"`
	VatNumber     string              `json:"vat_number"`
	JSON          invoiceCustomerJSON `json:"-"`
}

// invoiceCustomerJSON contains the JSON metadata for the struct [InvoiceCustomer]
type invoiceCustomerJSON struct {
	Address       apijson.Field
	CompanyNumber apijson.Field
	Logo          apijson.Field
	Name          apijson.Field
	VatNumber     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerJSON) RawJSON() string {
	return r.raw
}

type InvoiceObject string

const (
	InvoiceObjectInvoice InvoiceObject = "invoice"
)

func (r InvoiceObject) IsKnown() bool {
	switch r {
	case InvoiceObjectInvoice:
		return true
	}
	return false
}

type InvoiceTaxType string

const (
	InvoiceTaxTypeIncl InvoiceTaxType = "incl"
	InvoiceTaxTypeExcl InvoiceTaxType = "excl"
)

func (r InvoiceTaxType) IsKnown() bool {
	switch r {
	case InvoiceTaxTypeIncl, InvoiceTaxTypeExcl:
		return true
	}
	return false
}

type InvoiceTotals struct {
	// Total discount amount.
	Discount float64 `json:"discount"`
	// Total before VAT.
	Subtotal float64 `json:"subtotal"`
	// Grand total.
	Total float64 `json:"total"`
	// Total VAT amount.
	Vat  float64           `json:"vat"`
	JSON invoiceTotalsJSON `json:"-"`
}

// invoiceTotalsJSON contains the JSON metadata for the struct [InvoiceTotals]
type invoiceTotalsJSON struct {
	Discount    apijson.Field
	Subtotal    apijson.Field
	Total       apijson.Field
	Vat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceTotalsJSON) RawJSON() string {
	return r.raw
}

type InvoiceType string

const (
	InvoiceTypeSale   InvoiceType = "sale"
	InvoiceTypeRefund InvoiceType = "refund"
)

func (r InvoiceType) IsKnown() bool {
	switch r {
	case InvoiceTypeSale, InvoiceTypeRefund:
		return true
	}
	return false
}

type InvoiceBusinessInputParam struct {
	// Your business trading address.
	Address param.Field[string] `json:"address" api:"required"`
	// Your business trading name.
	Name param.Field[string] `json:"name" api:"required"`
	// Your business VAT number.
	VatNumber   param.Field[string] `json:"vat_number" api:"required"`
	BankAccount param.Field[string] `json:"bank_account"`
	// Your business company number.
	CompanyNumber param.Field[string] `json:"company_number"`
	Email         param.Field[string] `json:"email" format:"email"`
	// URL to your company logo (HTTPS only, .svg/.jpg/.png). Recommended 240px by
	// 60px.
	Logo    param.Field[string] `json:"logo" format:"uri"`
	Phone   param.Field[string] `json:"phone"`
	Website param.Field[string] `json:"website" format:"uri"`
}

func (r InvoiceBusinessInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceConversionInput struct {
	// The 3-character currency code for the conversion.
	CurrencyCode string `json:"currency_code" api:"required"`
	// The rate of conversion.
	Rate float64                    `json:"rate" api:"required"`
	JSON invoiceConversionInputJSON `json:"-"`
}

// invoiceConversionInputJSON contains the JSON metadata for the struct
// [InvoiceConversionInput]
type invoiceConversionInputJSON struct {
	CurrencyCode apijson.Field
	Rate         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceConversionInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceConversionInputJSON) RawJSON() string {
	return r.raw
}

type InvoiceConversionInputParam struct {
	// The 3-character currency code for the conversion.
	CurrencyCode param.Field[string] `json:"currency_code" api:"required"`
	// The rate of conversion.
	Rate param.Field[float64] `json:"rate" api:"required"`
}

func (r InvoiceConversionInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceCustomerInputParam struct {
	// The customer's trading name.
	Name          param.Field[string] `json:"name" api:"required"`
	Address       param.Field[string] `json:"address"`
	CompanyNumber param.Field[string] `json:"company_number"`
	CountryCode   param.Field[string] `json:"country_code"`
	Email         param.Field[string] `json:"email" format:"email"`
	// URL to the customer logo (HTTPS only, .jpg/.png).
	Logo      param.Field[string] `json:"logo" format:"uri"`
	VatNumber param.Field[string] `json:"vat_number"`
}

func (r InvoiceCustomerInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceResponse struct {
	Code    int64               `json:"code"`
	Data    Invoice             `json:"data"`
	Success bool                `json:"success"`
	JSON    invoiceResponseJSON `json:"-"`
}

// invoiceResponseJSON contains the JSON metadata for the struct [InvoiceResponse]
type invoiceResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceListResponse struct {
	Code    int64                   `json:"code"`
	Data    []Invoice               `json:"data"`
	Success bool                    `json:"success"`
	JSON    invoiceListResponseJSON `json:"-"`
}

// invoiceListResponseJSON contains the JSON metadata for the struct
// [InvoiceListResponse]
type invoiceListResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceListResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceDeleteResponse struct {
	Code    int64                     `json:"code" api:"required"`
	Success bool                      `json:"success" api:"required"`
	JSON    invoiceDeleteResponseJSON `json:"-"`
}

// invoiceDeleteResponseJSON contains the JSON metadata for the struct
// [InvoiceDeleteResponse]
type invoiceDeleteResponseJSON struct {
	Code        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceNewParams struct {
	CreateInvoice CreateInvoiceParam `json:"create_invoice" api:"required"`
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateInvoice)
}

type InvoiceUpdateParams struct {
	CreateInvoice CreateInvoiceParam `json:"create_invoice" api:"required"`
}

func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateInvoice)
}

type InvoiceListParams struct {
	// Number of invoices to return (default 10, max 100).
	Limit param.Field[int64] `query:"limit"`
	// Number of invoices to skip (default 0).
	Offset param.Field[int64] `query:"offset"`
	// Search query to filter invoices.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
