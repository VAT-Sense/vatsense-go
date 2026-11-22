// Checking the signature of VAT Sense webhook events.

package vatsense

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
)

// WebhookSignatureHeader is the name of the header that carries the signature.
const WebhookSignatureHeader = "VatSense-Signature"

// DefaultWebhookTolerance is how old an event may be before it is refused.
const DefaultWebhookTolerance = 5 * time.Minute

// ErrWebhookSignature is wrapped by every error from [VerifyWebhookSignature]:
// the header is missing or malformed, the event is too old, or no signature
// matches.
var ErrWebhookSignature = errors.New("vatsense: webhook signature")

// WebhookEvent is an event as VAT Sense sends it.
type WebhookEvent struct {
	// Starts `evt_`; the same on every retry, so a repeat can be ignored.
	ID string `json:"id"`
	// Such as `validation.completed`, `vies.country.unavailable` or
	// `batch.completed`.
	Type string `json:"type"`
	// When the event happened, in Unix seconds.
	Created int64 `json:"created"`
	// False for test mode events.
	Livemode bool `json:"livemode"`
	Data     struct {
		// The event's object: a `pending_check`, a `vies_country` or a `batch`.
		// Unmarshal it into the type you need, such as [Batch].
		Object json.RawMessage `json:"object"`
	} `json:"data"`
}

// WebhookOptions changes how [VerifyWebhookSignature] checks an event.
type WebhookOptions struct {
	// Refuse events signed more than this long ago. Zero means
	// [DefaultWebhookTolerance]; a negative value skips the check.
	Tolerance time.Duration
	// The current time, for tests. Zero means time.Now().
	Now time.Time
}

func webhookError(message string) error {
	return errors.Join(ErrWebhookSignature, errors.New(message))
}

// VerifyWebhookSignature checks an event's `VatSense-Signature` header against
// its raw body, exactly as received, with the endpoint's signing secret
// (`whsec_...`). Every `v1` signature in the header is tried, so events signed
// during a secret roll verify with either secret. The returned error wraps
// [ErrWebhookSignature].
func VerifyWebhookSignature(payload []byte, header string, secret string, opts ...WebhookOptions) error {
	var o WebhookOptions
	if len(opts) > 0 {
		o = opts[0]
	}
	if header == "" {
		return webhookError("the VatSense-Signature header is missing")
	}
	if secret == "" {
		return webhookError("no signing secret was given")
	}

	timestamp := int64(-1)
	var signatures []string
	for _, part := range strings.Split(header, ",") {
		key, value, found := strings.Cut(strings.TrimSpace(part), "=")
		if !found {
			continue
		}
		switch key {
		case "t":
			t, err := strconv.ParseInt(value, 10, 64)
			if err != nil || t < 0 {
				return webhookError("the signature header's timestamp is not a number")
			}
			timestamp = t
		case "v1":
			if value != "" {
				signatures = append(signatures, value)
			}
		}
	}
	if timestamp < 0 {
		return webhookError("the signature header has no timestamp")
	}
	if len(signatures) == 0 {
		return webhookError("the signature header has no v1 signature")
	}

	tolerance := o.Tolerance
	if tolerance == 0 {
		tolerance = DefaultWebhookTolerance
	}
	now := o.Now
	if now.IsZero() {
		now = time.Now()
	}
	if tolerance > 0 && math.Abs(float64(now.Unix()-timestamp)) > tolerance.Seconds() {
		return webhookError("the event's timestamp is outside the tolerance")
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(timestamp, 10) + "."))
	mac.Write(payload)
	expected := []byte(hex.EncodeToString(mac.Sum(nil)))
	for _, signature := range signatures {
		if hmac.Equal(expected, []byte(signature)) {
			return nil
		}
	}
	return webhookError("no signature matches the payload")
}

// ConstructWebhookEvent checks the signature with [VerifyWebhookSignature],
// then parses the event.
func ConstructWebhookEvent(payload []byte, header string, secret string, opts ...WebhookOptions) (*WebhookEvent, error) {
	if err := VerifyWebhookSignature(payload, header, secret, opts...); err != nil {
		return nil, err
	}
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, errors.Join(ErrWebhookSignature, err)
	}
	return &event, nil
}
