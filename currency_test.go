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

func TestCurrencyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Currency.List(context.TODO(), vatsense.CurrencyListParams{
		From: vatsense.F("USD,CAD,AUD"),
		To:   vatsense.F(vatsense.CurrencyListParamsToGbp),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCurrencyCalculateVatPrice(t *testing.T) {
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
	_, err := client.Currency.CalculateVatPrice(context.TODO(), vatsense.CurrencyCalculateVatPriceParams{
		Price:   vatsense.F("20.00"),
		TaxType: vatsense.F(vatsense.CurrencyCalculateVatPriceParamsTaxTypeExcl),
		VatRate: vatsense.F(5.000000),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCurrencyConvert(t *testing.T) {
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
	_, err := client.Currency.Convert(context.TODO(), vatsense.CurrencyConvertParams{
		Amount: vatsense.F("39.99"),
		From:   vatsense.F("USD"),
		To:     vatsense.F(vatsense.CurrencyConvertParamsToGbp),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
