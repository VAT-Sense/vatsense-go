// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/VAT-Sense/vatsense-go"
	"github.com/VAT-Sense/vatsense-go/internal/testutil"
	"github.com/VAT-Sense/vatsense-go/option"
)

func TestCountryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Countries.List(context.TODO(), vatsense.CountryListParams{
		CountryCode: vatsense.F("GB"),
		IPAddress:   vatsense.F("86.27.166.97"),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCountryListProvinces(t *testing.T) {
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
	_, err := client.Countries.ListProvinces(context.TODO(), vatsense.CountryListProvincesParams{
		CountryCode: vatsense.F("CA"),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
