package vatsense_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/VAT-Sense/vatsense-go"
	"github.com/VAT-Sense/vatsense-go/option"
)

// captureTransport records each request and answers with a fixed body.
type captureTransport struct {
	requests []*http.Request
	bodies   []string
	answer   string
}

func (c *captureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	body := ""
	if req.Body != nil {
		b, _ := io.ReadAll(req.Body)
		body = string(b)
	}
	c.requests = append(c.requests, req)
	c.bodies = append(c.bodies, body)
	return &http.Response{
		StatusCode: 200,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(stringsReader(c.answer)),
		Request:    req,
	}, nil
}

func stringsReader(s string) io.Reader { return &stringReader{s: s} }

type stringReader struct{ s string }

func (r *stringReader) Read(p []byte) (int, error) {
	if r.s == "" {
		return 0, io.EOF
	}
	n := copy(p, r.s)
	r.s = r.s[n:]
	return n, nil
}

func newCaptureClient(answer string) (vatsense.Client, *captureTransport) {
	t := &captureTransport{answer: answer}
	client := vatsense.NewClient(
		option.WithBaseURL("http://localhost:4010"),
		option.WithUsername("user"),
		option.WithPassword("key"),
		option.WithHTTPClient(&http.Client{Transport: t}),
	)
	return client, t
}

func TestInvoiceNewSendsIdempotencyKeyHeader(t *testing.T) {
	client, capture := newCaptureClient(`{"success":true,"code":200,"data":{}}`)
	_, err := client.Invoice.New(context.TODO(), vatsense.InvoiceNewParams{
		IdempotencyKey: vatsense.String("order-203"),
		CreateInvoice: vatsense.CreateInvoiceParam{
			Business:          vatsense.InvoiceBusinessInputParam{Name: "VAT Sense", Address: "London", VatNumber: "GB12345678"},
			CurrencyCode:      "GBP",
			Date:              "2026-11-05",
			TaxPoint:          "2026-11-05",
			Items:             []vatsense.InvoiceItemInputParam{{Item: "Plan", PriceEach: 10, Quantity: 1, VatRate: 20}},
			OriginalInvoiceID: vatsense.String("in5aeae457cda2a"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := capture.requests[0].Header.Get("Idempotency-Key"); got != "order-203" {
		t.Fatalf("Idempotency-Key header: got %q", got)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(capture.bodies[0]), &body); err != nil {
		t.Fatal(err)
	}
	if body["original_invoice_id"] != "in5aeae457cda2a" {
		t.Fatalf("original_invoice_id: got %v", body["original_invoice_id"])
	}
	if _, ok := body["IdempotencyKey"]; ok {
		t.Fatal("the key leaked into the body")
	}
}

func TestInvoiceNewWithoutKeySendsNoHeader(t *testing.T) {
	client, capture := newCaptureClient(`{"success":true,"code":200,"data":{}}`)
	_, err := client.Invoice.New(context.TODO(), vatsense.InvoiceNewParams{
		CreateInvoice: vatsense.CreateInvoiceParam{
			Business:     vatsense.InvoiceBusinessInputParam{Name: "VAT Sense", Address: "London", VatNumber: "GB12345678"},
			CurrencyCode: "GBP",
			Date:         "2026-11-05",
			TaxPoint:     "2026-11-05",
			Items:        []vatsense.InvoiceItemInputParam{{Item: "Plan", PriceEach: 10, Quantity: 1, VatRate: 20}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := capture.requests[0].Header.Get("Idempotency-Key"); got != "" {
		t.Fatalf("unexpected Idempotency-Key header %q", got)
	}
}

func TestValidateMatch(t *testing.T) {
	client, capture := newCaptureClient(`{"success":true,"code":200,"data":{"valid":true,"match":{"result":"likely","score":90,"fields":{"company_name":{"result":"match","score":100,"vies":null}}}}}`)
	res, err := client.Validate.Check(context.TODO(), vatsense.ValidateCheckParams{
		VatNumber:        vatsense.String("GB288305674"),
		MatchCompanyName: vatsense.String("Weio Ltd"),
		MatchPostcode:    vatsense.String("SW3 1GL"),
	})
	if err != nil {
		t.Fatal(err)
	}
	q := capture.requests[0].URL.Query()
	if q.Get("match_company_name") != "Weio Ltd" || q.Get("match_postcode") != "SW3 1GL" {
		t.Fatalf("query: %v", q)
	}
	if res.Data.Match.Result != "likely" || res.Data.Match.Score != 90 || res.Data.Match.Fields["company_name"].Score != 100 {
		t.Fatalf("match: %+v", res.Data.Match)
	}
}

func TestRateListProvinces(t *testing.T) {
	client, capture := newCaptureClient(`{"success":true,"code":200,"data":[]}`)
	_, _ = client.Rates.List(context.TODO(), vatsense.RateListParams{Provinces: vatsense.Bool(false)})
	if got := capture.requests[0].URL.Query().Get("provinces"); got != "false" {
		t.Fatalf("provinces: got %q", got)
	}
}
