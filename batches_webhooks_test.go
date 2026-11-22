package vatsense_test

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/VAT-Sense/vatsense-go"
	"github.com/VAT-Sense/vatsense-go/option"
)

const testBatch = `{"object":"batch","id":"01JM0000000000000000000000","number":4,"status":"processing","livemode":true,"rows":{"total":2,"to_check":2,"done":1},"quote":null}`

func TestBatchNewSendsBodyAndIdempotencyKey(t *testing.T) {
	client, capture := newCaptureClient(`{"success":true,"code":201,"data":` + testBatch + `}`)
	res, err := client.Batches.New(context.TODO(), vatsense.BatchNewParams{
		IdempotencyKey: vatsense.String("nov-1"),
		Rows:           []vatsense.BatchRowParam{{VatNumber: "GB288305674", CompanyName: vatsense.String("Weio Ltd")}},
		Shortfall:      vatsense.BatchShortfallPause,
		AutoContinue:   vatsense.Bool(false),
	})
	if err != nil {
		t.Fatal(err)
	}
	req := capture.requests[0]
	if req.Method != http.MethodPost || req.URL.Path != "/batches" {
		t.Fatalf("request: %s %s", req.Method, req.URL.Path)
	}
	if got := req.Header.Get("Idempotency-Key"); got != "nov-1" {
		t.Fatalf("Idempotency-Key header: got %q", got)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(capture.bodies[0]), &body); err != nil {
		t.Fatal(err)
	}
	want := `{"auto_continue":false,"rows":[{"company_name":"Weio Ltd","vat_number":"GB288305674"}],"shortfall":"pause"}`
	if got, _ := json.Marshal(body); string(got) != want {
		t.Fatalf("body: got %s", got)
	}
	if res.Data.ID != "01JM0000000000000000000000" || res.Data.Rows.Done != 1 || !res.Data.Livemode {
		t.Fatalf("batch: %+v", res.Data)
	}
}

func TestBatchActionsCallTheirPaths(t *testing.T) {
	id := "01JM0000000000000000000000"
	client, capture := newCaptureClient(`{"success":true,"code":200,"data":` + testBatch + `}`)
	ctx := context.TODO()
	steps := []func() error{
		func() error {
			_, err := client.Batches.List(ctx, vatsense.BatchListParams{Page: vatsense.Int(2), Limit: vatsense.Int(10)})
			return err
		},
		func() error { _, err := client.Batches.Get(ctx, id); return err },
		func() error {
			_, err := client.Batches.Start(ctx, id, vatsense.BatchStartParams{Shortfall: vatsense.BatchShortfallOverage})
			return err
		},
		func() error { _, err := client.Batches.Pause(ctx, id); return err },
		func() error { _, err := client.Batches.Continue(ctx, id, vatsense.BatchContinueParams{}); return err },
		func() error { _, err := client.Batches.Complete(ctx, id); return err },
		func() error { _, err := client.Batches.Delete(ctx, id); return err },
		func() error {
			_, err := client.Batches.Results(ctx, id, vatsense.BatchResultsParams{After: vatsense.Int(100), Limit: vatsense.Int(50)})
			return err
		},
	}
	for _, step := range steps {
		if err := step(); err != nil {
			t.Fatal(err)
		}
	}
	var got []string
	for _, req := range capture.requests {
		line := req.Method + " " + req.URL.Path
		if req.URL.RawQuery != "" {
			line += "?" + req.URL.RawQuery
		}
		got = append(got, line)
	}
	want := []string{
		"GET /batches?limit=10&page=2",
		"GET /batches/" + id,
		"POST /batches/" + id + "/start",
		"POST /batches/" + id + "/pause",
		"POST /batches/" + id + "/continue",
		"POST /batches/" + id + "/complete",
		"DELETE /batches/" + id,
		"GET /batches/" + id + "/results?after=100&limit=50",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("calls:\n%s", strings.Join(got, "\n"))
	}
	if capture.bodies[2] != `{"shortfall":"overage"}` {
		t.Fatalf("start body: %s", capture.bodies[2])
	}
	if _, err := client.Batches.Get(ctx, ""); err == nil {
		t.Fatal("an empty id should be refused")
	}
}

func TestBatchResultsParse(t *testing.T) {
	client, _ := newCaptureClient(`{"success":true,"code":200,"data":{"data":[{"row":1,"vat_number":"GB288305674","result":"valid","confirmation_number":null}],"has_more":true,"next_after":1}}`)
	res, err := client.Batches.Results(context.TODO(), "01JM0000000000000000000000", vatsense.BatchResultsParams{})
	if err != nil {
		t.Fatal(err)
	}
	if !res.Data.HasMore || res.Data.NextAfter != 1 || res.Data.Data[0].Result != "valid" {
		t.Fatalf("results: %+v", res.Data)
	}
}

type csvTransport struct{}

func (csvTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Header:     http.Header{"Content-Type": []string{"text/csv; charset=UTF-8"}},
		Body:       io.NopCloser(strings.NewReader("row,vat_number\n1,GB288305674\n")),
		Request:    req,
	}, nil
}

func TestBatchResultsCsv(t *testing.T) {
	client := vatsense.NewClient(
		option.WithBaseURL("http://localhost:4010"),
		option.WithUsername("user"),
		option.WithPassword("key"),
		option.WithHTTPClient(&http.Client{Transport: csvTransport{}}),
	)
	csv, err := client.Batches.ResultsCsv(context.TODO(), "01JM0000000000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if string(csv) != "row,vat_number\n1,GB288305674\n" {
		t.Fatalf("csv: %q", csv)
	}
}

func TestValidateConsultationOptionalAndConfirmationNumber(t *testing.T) {
	client, capture := newCaptureClient(`{"success":true,"code":200,"data":{"valid":true,"confirmation_number":"VSC-ABC"}}`)
	res, err := client.Validate.Check(context.TODO(), vatsense.ValidateCheckParams{
		VatNumber:            vatsense.String("FR40303265045"),
		RequesterVatNumber:   vatsense.String("GB288305674"),
		ConsultationOptional: vatsense.Bool(true),
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := capture.requests[0].URL.Query().Get("consultation_optional"); got != "true" {
		t.Fatalf("consultation_optional: got %q", got)
	}
	if res.Data.ConfirmationNumber != "VSC-ABC" {
		t.Fatalf("confirmation_number: got %q", res.Data.ConfirmationNumber)
	}
}

const webhookSecret = "whsec_test_secret"
const webhookPayload = `{"id":"evt_1","type":"batch.completed","created":1793365200,"livemode":true,"data":{"object":{"object":"batch","id":"01JM0000000000000000000000"}}}`

func sign(timestamp int64, body string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	fmt.Fprintf(mac, "%d.%s", timestamp, body)
	return hex.EncodeToString(mac.Sum(nil))
}

func TestWebhookSignatureAcceptsAGoodEvent(t *testing.T) {
	header := "t=1793365200,v1=" + sign(1793365200, webhookPayload, webhookSecret)
	event, err := vatsense.ConstructWebhookEvent([]byte(webhookPayload), header, webhookSecret, vatsense.WebhookOptions{Now: time.Unix(1793365260, 0)})
	if err != nil {
		t.Fatal(err)
	}
	var batch vatsense.Batch
	if err := json.Unmarshal(event.Data.Object, &batch); err != nil {
		t.Fatal(err)
	}
	if event.Type != "batch.completed" || batch.ID != "01JM0000000000000000000000" {
		t.Fatalf("event: %+v", event)
	}
}

func TestWebhookSignatureAcceptsTheOldSecretDuringARoll(t *testing.T) {
	header := "t=1793365200,v1=" + sign(1793365200, webhookPayload, "whsec_new") + ",v1=" + sign(1793365200, webhookPayload, webhookSecret)
	if err := vatsense.VerifyWebhookSignature([]byte(webhookPayload), header, webhookSecret, vatsense.WebhookOptions{Now: time.Unix(1793365200, 0)}); err != nil {
		t.Fatal(err)
	}
}

func TestWebhookSignatureRefusesBadEvents(t *testing.T) {
	cases := []struct {
		header string
		now    int64
	}{
		{"", 1793365200},
		{"v1=abc", 1793365200},
		{"t=1793365200", 1793365200},
		{"t=soon,v1=abc", 1793365200},
		{"t=1793365200,v1=" + sign(1793365200, webhookPayload, "whsec_other"), 1793365200},
		{"t=1793365200,v1=" + sign(1793365200, webhookPayload, webhookSecret), 1793365200 + 301},
	}
	for _, c := range cases {
		err := vatsense.VerifyWebhookSignature([]byte(webhookPayload), c.header, webhookSecret, vatsense.WebhookOptions{Now: time.Unix(c.now, 0)})
		if !errors.Is(err, vatsense.ErrWebhookSignature) {
			t.Fatalf("%q: got %v", c.header, err)
		}
	}
	header := "t=1793365200,v1=" + sign(1793365200, webhookPayload, webhookSecret)
	changed := strings.Replace(webhookPayload, "true", "false", 1)
	if err := vatsense.VerifyWebhookSignature([]byte(changed), header, webhookSecret, vatsense.WebhookOptions{Now: time.Unix(1793365200, 0)}); !errors.Is(err, vatsense.ErrWebhookSignature) {
		t.Fatalf("changed body: got %v", err)
	}
}
