// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	"github.com/VAT-Sense/vatsense-go/internal/param"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
)

// VAT-compliant invoice management
//
// InvoiceItemService contains methods and other services that help with
// interacting with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceItemService] method instead.
type InvoiceItemService struct {
	Options []option.RequestOption
}

// NewInvoiceItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvoiceItemService(opts ...option.RequestOption) (r *InvoiceItemService) {
	r = &InvoiceItemService{}
	r.Options = opts
	return
}

// Retrieve a specific line item from an invoice.
func (r *InvoiceItemService) Get(ctx context.Context, invoiceID string, itemID string, opts ...option.RequestOption) (res *InvoiceItemGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item/%s", invoiceID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a specific line item on an invoice.
func (r *InvoiceItemService) Update(ctx context.Context, invoiceID string, itemID string, body InvoiceItemUpdateParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item/%s", invoiceID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Remove a specific line item from an invoice.
func (r *InvoiceItemService) Delete(ctx context.Context, invoiceID string, itemID string, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item/%s", invoiceID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Add one or more line items to an existing invoice.
func (r *InvoiceItemService) Add(ctx context.Context, invoiceID string, body InvoiceItemAddParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type InvoiceItem struct {
	ID           string            `json:"id"`
	DiscountRate float64           `json:"discount_rate" api:"nullable"`
	Item         string            `json:"item"`
	Object       InvoiceItemObject `json:"object"`
	PriceEach    float64           `json:"price_each"`
	PriceTotal   float64           `json:"price_total"`
	Quantity     float64           `json:"quantity"`
	VatRate      float64           `json:"vat_rate"`
	JSON         invoiceItemJSON   `json:"-"`
}

// invoiceItemJSON contains the JSON metadata for the struct [InvoiceItem]
type invoiceItemJSON struct {
	ID           apijson.Field
	DiscountRate apijson.Field
	Item         apijson.Field
	Object       apijson.Field
	PriceEach    apijson.Field
	PriceTotal   apijson.Field
	Quantity     apijson.Field
	VatRate      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceItemJSON) RawJSON() string {
	return r.raw
}

type InvoiceItemObject string

const (
	InvoiceItemObjectItem InvoiceItemObject = "item"
)

func (r InvoiceItemObject) IsKnown() bool {
	switch r {
	case InvoiceItemObjectItem:
		return true
	}
	return false
}

type InvoiceItemInputParam struct {
	// The description of the line item.
	Item param.Field[string] `json:"item" api:"required"`
	// The price per item. Must be a decimal with 2 decimal places.
	PriceEach param.Field[float64] `json:"price_each" api:"required"`
	// The quantity of the item.
	Quantity param.Field[float64] `json:"quantity" api:"required"`
	// A percentage VAT rate for this item.
	VatRate param.Field[float64] `json:"vat_rate" api:"required"`
	// A percentage discount to apply to the price.
	DiscountRate param.Field[float64] `json:"discount_rate"`
}

func (r InvoiceItemInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceItemGetResponse struct {
	Code    int64                      `json:"code"`
	Data    InvoiceItem                `json:"data"`
	Success bool                       `json:"success"`
	JSON    invoiceItemGetResponseJSON `json:"-"`
}

// invoiceItemGetResponseJSON contains the JSON metadata for the struct
// [InvoiceItemGetResponse]
type invoiceItemGetResponseJSON struct {
	Code        apijson.Field
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceItemUpdateParams struct {
	InvoiceItemInput InvoiceItemInputParam `json:"invoice_item_input" api:"required"`
}

func (r InvoiceItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InvoiceItemInput)
}

type InvoiceItemAddParams struct {
	Items param.Field[[]InvoiceItemInputParam] `json:"items" api:"required"`
}

func (r InvoiceItemAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
