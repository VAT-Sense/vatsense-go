// Batches: bulk validation through `/batches` (November 2026).

package vatsense

import (
	"context"
	"errors"
	"fmt"
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

// Bulk validation: check many VAT numbers in one batch.
//
// Batch calls are free and work even when your requests are used up; each number
// a batch checks counts as one request. With a test key a batch is answered from
// the test numbers, is free and holds at most 1,000 numbers.
//
// BatchService contains methods and other services that help with interacting
// with the vat-sense API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.options = opts
	return
}

// Create a batch of VAT numbers and, by default, start checking it. Set exactly
// one of Numbers or Rows.
func (r *BatchService) New(ctx context.Context, body BatchNewParams, opts ...option.RequestOption) (res *BatchResponse, err error) {
	if body.IdempotencyKey.Valid() {
		opts = slices.Concat([]option.RequestOption{option.WithHeader("Idempotency-Key", body.IdempotencyKey.Value)}, opts)
	}
	opts = slices.Concat(r.options, opts)
	path := "batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The account's batches, newest first.
func (r *BatchService) List(ctx context.Context, query BatchListParams, opts ...option.RequestOption) (res *BatchListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A batch and its progress. Poll until Status is `complete`, or take the
// `batch.completed` webhook.
func (r *BatchService) Get(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchResponse, err error) {
	return r.action(ctx, http.MethodGet, batchID, "", nil, opts)
}

// Delete a draft. A batch that has started cannot be deleted (409).
func (r *BatchService) Delete(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batches/%s", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Start a draft, with the same options as [BatchService.New].
func (r *BatchService) Start(ctx context.Context, batchID string, body BatchStartParams, opts ...option.RequestOption) (res *BatchResponse, err error) {
	return r.action(ctx, http.MethodPost, batchID, "/start", body, opts)
}

// Pause a running batch.
func (r *BatchService) Pause(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchResponse, err error) {
	return r.action(ctx, http.MethodPost, batchID, "/pause", nil, opts)
}

// Continue a paused batch. Set Shortfall to `overage` to run into overage.
func (r *BatchService) Continue(ctx context.Context, batchID string, body BatchContinueParams, opts ...option.RequestOption) (res *BatchResponse, err error) {
	return r.action(ctx, http.MethodPost, batchID, "/continue", body, opts)
}

// Mark a paused batch complete, leaving the rest unchecked.
func (r *BatchService) Complete(ctx context.Context, batchID string, opts ...option.RequestOption) (res *BatchResponse, err error) {
	return r.action(ctx, http.MethodPost, batchID, "/complete", nil, opts)
}

// The batch's checked rows as JSON, in row order. Follow NextAfter while HasMore
// is true.
func (r *BatchService) Results(ctx context.Context, batchID string, query BatchResultsParams, opts ...option.RequestOption) (res *BatchResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batches/%s/results", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// The batch's CSV report, the same file the dashboard gives.
func (r *BatchService) ResultsCsv(ctx context.Context, batchID string, opts ...option.RequestOption) (res []byte, err error) {
	opts = slices.Concat(r.options, []option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batches/%s/results.csv", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *BatchService) action(ctx context.Context, method string, batchID string, suffix string, body any, opts []option.RequestOption) (res *BatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("batches/%s%s", url.PathEscape(batchID), suffix)
	err = requestconfig.ExecuteNewRequest(ctx, method, path, body, &res, opts...)
	return res, err
}

// A batch of VAT numbers checked together.
type Batch struct {
	// The batch's 26-character id.
	ID string `json:"id"`
	// Always `batch`.
	Object       string `json:"object"`
	AutoContinue bool   `json:"auto_continue"`
	Consultation bool   `json:"consultation"`
	// When the batch was made. Times are UTC, ISO 8601.
	CreatedAt string `json:"created_at"`
	// The batch's page in the dashboard; empty for a test batch.
	DashboardURL      string `json:"dashboard_url" api:"nullable"`
	EstimatedFinishAt string `json:"estimated_finish_at" api:"nullable"`
	ExpiresAt         string `json:"expires_at" api:"nullable"`
	FinishedAt        string `json:"finished_at" api:"nullable"`
	// False for a test batch, made with a test key.
	Livemode    bool   `json:"livemode"`
	Number      int64  `json:"number"`
	PauseReason string `json:"pause_reason" api:"nullable"`
	// For drafts and paused batches.
	Quote        BatchQuote   `json:"quote" api:"nullable"`
	Reference    string       `json:"reference" api:"nullable"`
	RequestsUsed int64        `json:"requests_used"`
	Results      BatchResults `json:"results"`
	Rows         BatchRows    `json:"rows"`
	Source       string       `json:"source"`
	StartedAt    string       `json:"started_at" api:"nullable"`
	// `draft`, `queued`, `processing`, `paused`, `complete` or `expired`.
	Status      string `json:"status"`
	Trial       bool   `json:"trial"`
	UsesOverage bool   `json:"uses_overage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Object            respjson.Field
		AutoContinue      respjson.Field
		Consultation      respjson.Field
		CreatedAt         respjson.Field
		DashboardURL      respjson.Field
		EstimatedFinishAt respjson.Field
		ExpiresAt         respjson.Field
		FinishedAt        respjson.Field
		Livemode          respjson.Field
		Number            respjson.Field
		PauseReason       respjson.Field
		Quote             respjson.Field
		Reference         respjson.Field
		RequestsUsed      respjson.Field
		Results           respjson.Field
		Rows              respjson.Field
		Source            respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		Trial             respjson.Field
		UsesOverage       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Batch) RawJSON() string { return r.JSON.raw }
func (r *Batch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchQuote struct {
	// Whether the requests left cover the batch.
	Covered        bool   `json:"covered"`
	OverageOn      bool   `json:"overage_on"`
	RenewsAt       string `json:"renews_at" api:"nullable"`
	RequestsLeft   int64  `json:"requests_left"`
	RequestsNeeded int64  `json:"requests_needed"`
	Shortfall      int64  `json:"shortfall"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Covered        respjson.Field
		OverageOn      respjson.Field
		RenewsAt       respjson.Field
		RequestsLeft   respjson.Field
		RequestsNeeded respjson.Field
		Shortfall      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchQuote) RawJSON() string { return r.JSON.raw }
func (r *BatchQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResults struct {
	Duplicate   int64 `json:"duplicate"`
	Invalid     int64 `json:"invalid"`
	NotChecked  int64 `json:"not_checked"`
	Repaired    int64 `json:"repaired"`
	Retrying    int64 `json:"retrying"`
	Unsupported int64 `json:"unsupported"`
	Valid       int64 `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duplicate   respjson.Field
		Invalid     respjson.Field
		NotChecked  respjson.Field
		Repaired    respjson.Field
		Retrying    respjson.Field
		Unsupported respjson.Field
		Valid       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResults) RawJSON() string { return r.JSON.raw }
func (r *BatchResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchRows struct {
	// Rows checked so far.
	Done int64 `json:"done"`
	// Rows that will be sent for checking.
	ToCheck int64 `json:"to_check"`
	// Rows in the batch.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Done        respjson.Field
		ToCheck     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchRows) RawJSON() string { return r.JSON.raw }
func (r *BatchRows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One checked row of a batch, as the CSV report lists it.
type BatchResult struct {
	Row                int64            `json:"row"`
	Authority          string           `json:"authority" api:"nullable"`
	CheckedAt          string           `json:"checked_at" api:"nullable"`
	CompanyAddress     string           `json:"company_address" api:"nullable"`
	CompanyName        string           `json:"company_name" api:"nullable"`
	ConfirmationNumber string           `json:"confirmation_number" api:"nullable"`
	ConsultationNumber string           `json:"consultation_number" api:"nullable"`
	CountryCode        string           `json:"country_code" api:"nullable"`
	CountryUploaded    string           `json:"country_uploaded" api:"nullable"`
	DuplicateOfRow     int64            `json:"duplicate_of_row" api:"nullable"`
	Match              BatchResultMatch `json:"match" api:"nullable"`
	Reason             string           `json:"reason" api:"nullable"`
	Repaired           bool             `json:"repaired"`
	// Such as `valid` or `invalid`; `pending` while the row is being checked.
	Result string `json:"result"`
	// The number as it was checked.
	VatNumber         string `json:"vat_number" api:"nullable"`
	VatNumberUploaded string `json:"vat_number_uploaded" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Row                respjson.Field
		Authority          respjson.Field
		CheckedAt          respjson.Field
		CompanyAddress     respjson.Field
		CompanyName        respjson.Field
		ConfirmationNumber respjson.Field
		ConsultationNumber respjson.Field
		CountryCode        respjson.Field
		CountryUploaded    respjson.Field
		DuplicateOfRow     respjson.Field
		Match              respjson.Field
		Reason             respjson.Field
		Repaired           respjson.Field
		Result             respjson.Field
		VatNumber          respjson.Field
		VatNumberUploaded  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResult) RawJSON() string { return r.JSON.raw }
func (r *BatchResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResultMatch struct {
	Fields map[string]BatchResultMatchField `json:"fields"`
	// `likely`, `possible`, `unlikely` or `not_applicable`.
	Result string `json:"result"`
	Score  int64  `json:"score" api:"nullable"`
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
func (r BatchResultMatch) RawJSON() string { return r.JSON.raw }
func (r *BatchResultMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResultMatchField struct {
	Result string `json:"result" api:"nullable"`
	Score  int64  `json:"score" api:"nullable"`
	// The detail as it was sent.
	Sent string `json:"sent" api:"nullable"`
	Vies string `json:"vies" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Score       respjson.Field
		Sent        respjson.Field
		Vies        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResultMatchField) RawJSON() string { return r.JSON.raw }
func (r *BatchResultMatchField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponse struct {
	Code    int64 `json:"code"`
	Data    Batch `json:"data"`
	Success bool  `json:"success"`
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
func (r BatchResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponse struct {
	Code    int64                 `json:"code"`
	Data    BatchListResponseData `json:"data"`
	Success bool                  `json:"success"`
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
func (r BatchListResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListResponseData struct {
	Data    []Batch `json:"data"`
	HasMore bool    `json:"has_more"`
	Limit   int64   `json:"limit"`
	Page    int64   `json:"page"`
	Total   int64   `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResultsResponse struct {
	Code    int64                    `json:"code"`
	Data    BatchResultsResponseData `json:"data"`
	Success bool                     `json:"success"`
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
func (r BatchResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResultsResponseData struct {
	Data    []BatchResult `json:"data"`
	HasMore bool          `json:"has_more"`
	// Pass as `After` for the next page while `HasMore` is true.
	NextAfter int64 `json:"next_after" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextAfter   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResultsResponseData) RawJSON() string { return r.JSON.raw }
func (r *BatchResultsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchDeleteResponse struct {
	Code    int64                   `json:"code"`
	Data    BatchDeleteResponseData `json:"data"`
	Success bool                    `json:"success"`
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
func (r BatchDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchDeleteResponseData struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Object  string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *BatchDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What to do when your requests do not cover the batch.
type BatchShortfall string

const (
	// Start nothing and return the quote (the default).
	BatchShortfallWait BatchShortfall = "wait"
	// Check what is covered and pause the rest.
	BatchShortfallPause BatchShortfall = "pause"
	// Run into overage, which must already be on in the dashboard.
	BatchShortfallOverage BatchShortfall = "overage"
)

// A row to check, with optional company details for matching.
//
// The property VatNumber is required.
type BatchRowParam struct {
	VatNumber    string            `json:"vat_number" api:"required"`
	Address      param.Opt[string] `json:"address,omitzero"`
	AddressLine2 param.Opt[string] `json:"address_line_2,omitzero"`
	City         param.Opt[string] `json:"city,omitzero"`
	CompanyName  param.Opt[string] `json:"company_name,omitzero"`
	CountryCode  param.Opt[string] `json:"country_code,omitzero"`
	Postcode     param.Opt[string] `json:"postcode,omitzero"`
	Street       param.Opt[string] `json:"street,omitzero"`
	paramObj
}

func (r BatchRowParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchRowParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchRowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewParams struct {
	// Sent as the `Idempotency-Key` header, at most 64 characters. The same key
	// again returns the batch already made.
	IdempotencyKey param.Opt[string] `json:"-"`
	// The VAT numbers to check, each with its 2-character country code.
	Numbers []string `json:"numbers,omitzero"`
	// Rows with VatNumber and, for company matching, any of the other details.
	Rows []BatchRowParam `json:"rows,omitzero"`
	// Your own label for the batch, at most 255 characters.
	Reference param.Opt[string] `json:"reference,omitzero"`
	// Defaults to true. False keeps a draft for 24 hours, to start with
	// [BatchService.Start].
	Start param.Opt[bool] `json:"start,omitzero"`
	// Any of "wait", "pause", "overage".
	Shortfall BatchShortfall `json:"shortfall,omitzero"`
	// Defaults to true. With `pause`, continue when your requests renew.
	AutoContinue param.Opt[bool] `json:"auto_continue,omitzero"`
	// Defaults to false. Send the requester VAT number from your account settings
	// with each check.
	Consultation param.Opt[bool] `json:"consultation,omitzero"`
	paramObj
}

func (r BatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchStartParams struct {
	// Any of "wait", "pause", "overage".
	Shortfall    BatchShortfall  `json:"shortfall,omitzero"`
	AutoContinue param.Opt[bool] `json:"auto_continue,omitzero"`
	Consultation param.Opt[bool] `json:"consultation,omitzero"`
	paramObj
}

func (r BatchStartParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchContinueParams struct {
	// Any of "wait", "pause", "overage".
	Shortfall BatchShortfall `json:"shortfall,omitzero"`
	paramObj
}

func (r BatchContinueParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchContinueParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchContinueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchListParams struct {
	// Batches a page, 1 to 100 (default 25).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page, from 1.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// `draft` lists the drafts not yet started instead.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchListParams]'s query parameters as `url.Values`.
func (r BatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchResultsParams struct {
	// Return rows after this row number (default 0).
	After param.Opt[int64] `query:"after,omitzero" json:"-"`
	// Rows a page, 1 to 1,000 (default 100).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchResultsParams]'s query parameters as `url.Values`.
func (r BatchResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
