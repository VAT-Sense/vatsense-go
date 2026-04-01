// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/VAT-Sense/vatsense-go/internal/apijson"
	shimjson "github.com/VAT-Sense/vatsense-go/internal/encoding/json"
	"github.com/VAT-Sense/vatsense-go/internal/requestconfig"
	"github.com/VAT-Sense/vatsense-go/option"
	"github.com/VAT-Sense/vatsense-go/packages/param"
	"github.com/VAT-Sense/vatsense-go/packages/respjson"
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
	options []option.RequestOption
}

// NewInvoiceItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvoiceItemService(opts ...option.RequestOption) (r InvoiceItemService) {
	r = InvoiceItemService{}
	r.options = opts
	return
}

// Retrieve a specific line item from an invoice.
func (r *InvoiceItemService) Get(ctx context.Context, itemID string, query InvoiceItemGetParams, opts ...option.RequestOption) (res *InvoiceItemGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.InvoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item/%s", url.PathEscape(query.InvoiceID), url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a specific line item on an invoice.
func (r *InvoiceItemService) Update(ctx context.Context, itemID string, params InvoiceItemUpdateParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.InvoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item/%s", url.PathEscape(params.InvoiceID), url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Remove a specific line item from an invoice.
func (r *InvoiceItemService) Delete(ctx context.Context, itemID string, body InvoiceItemDeleteParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.InvoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item/%s", url.PathEscape(body.InvoiceID), url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Add one or more line items to an existing invoice.
func (r *InvoiceItemService) Add(ctx context.Context, invoiceID string, body InvoiceItemAddParams, opts ...option.RequestOption) (res *InvoiceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("invoice/%s/item", url.PathEscape(invoiceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type InvoiceItem struct {
	ID           string  `json:"id"`
	DiscountRate float64 `json:"discount_rate" api:"nullable"`
	Item         string  `json:"item"`
	// Any of "item".
	Object     InvoiceItemObject `json:"object"`
	PriceEach  float64           `json:"price_each"`
	PriceTotal float64           `json:"price_total"`
	Quantity   float64           `json:"quantity"`
	VatRate    float64           `json:"vat_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DiscountRate respjson.Field
		Item         respjson.Field
		Object       respjson.Field
		PriceEach    respjson.Field
		PriceTotal   respjson.Field
		Quantity     respjson.Field
		VatRate      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceItem) RawJSON() string { return r.JSON.raw }
func (r *InvoiceItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceItemObject string

const (
	InvoiceItemObjectItem InvoiceItemObject = "item"
)

// The properties Item, PriceEach, Quantity, VatRate are required.
type InvoiceItemInputParam struct {
	// The description of the line item.
	Item string `json:"item" api:"required"`
	// The price per item. Must be a decimal with 2 decimal places.
	PriceEach float64 `json:"price_each" api:"required"`
	// The quantity of the item.
	Quantity float64 `json:"quantity" api:"required"`
	// A percentage VAT rate for this item.
	VatRate float64 `json:"vat_rate" api:"required"`
	// A percentage discount to apply to the price.
	DiscountRate param.Opt[float64] `json:"discount_rate,omitzero"`
	paramObj
}

func (r InvoiceItemInputParam) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceItemInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceItemInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceItemGetResponse struct {
	Code    int64       `json:"code"`
	Data    InvoiceItem `json:"data"`
	Success bool        `json:"success"`
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
func (r InvoiceItemGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceItemGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceItemGetParams struct {
	InvoiceID string `path:"invoice_id" api:"required" json:"-"`
	paramObj
}

type InvoiceItemUpdateParams struct {
	InvoiceID        string `path:"invoice_id" api:"required" json:"-"`
	InvoiceItemInput InvoiceItemInputParam
	paramObj
}

func (r InvoiceItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.InvoiceItemInput)
}
func (r *InvoiceItemUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceItemDeleteParams struct {
	InvoiceID string `path:"invoice_id" api:"required" json:"-"`
	paramObj
}

type InvoiceItemAddParams struct {
	Items []InvoiceItemInputParam `json:"items,omitzero" api:"required"`
	paramObj
}

func (r InvoiceItemAddParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceItemAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceItemAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
