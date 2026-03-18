// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/vat-sense-go"
	"github.com/stainless-sdks/vat-sense-go/internal/testutil"
	"github.com/stainless-sdks/vat-sense-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vatsense.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	rates, err := client.Rates.List(context.TODO(), vatsense.RateListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", rates.Code)
}
