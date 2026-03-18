// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/VAT-Sense/vatsense-go"
	"github.com/VAT-Sense/vatsense-go/internal/testutil"
	"github.com/VAT-Sense/vatsense-go/option"
)

func TestRateListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rates.List(context.TODO(), vatsense.RateListParams{
		CountryCode: vatsense.String("GB"),
		Eu:          vatsense.Bool(true),
		IPAddress:   vatsense.String("86.27.166.97"),
		Period:      vatsense.Time(time.Now()),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRateCalculatePriceWithOptionalParams(t *testing.T) {
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
	_, err := client.Rates.CalculatePrice(context.TODO(), vatsense.RateCalculatePriceParams{
		Price:        "20.00",
		TaxType:      vatsense.RateCalculatePriceParamsTaxTypeExcl,
		CountryCode:  vatsense.String("GB"),
		Eu:           vatsense.Bool(true),
		IPAddress:    vatsense.String("86.27.166.97"),
		ProvinceCode: vatsense.String("ON"),
		Type:         vatsense.String("ebooks"),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRateDetailsWithOptionalParams(t *testing.T) {
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
	_, err := client.Rates.Details(context.TODO(), vatsense.RateDetailsParams{
		CountryCode:  vatsense.String("GB"),
		Eu:           vatsense.Bool(true),
		IPAddress:    vatsense.String("86.27.166.97"),
		Period:       vatsense.Time(time.Now()),
		ProvinceCode: vatsense.String("ON"),
		Type:         vatsense.String("ebooks"),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRateFindWithOptionalParams(t *testing.T) {
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
	_, err := client.Rates.Find(context.TODO(), vatsense.RateFindParams{
		CountryCode:  vatsense.String("GB"),
		Eu:           vatsense.Bool(true),
		IPAddress:    vatsense.String("86.27.166.97"),
		Period:       vatsense.Time(time.Now()),
		ProvinceCode: vatsense.String("ON"),
		Type:         vatsense.String("ebooks"),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRateListTypes(t *testing.T) {
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
	_, err := client.Rates.ListTypes(context.TODO())
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
