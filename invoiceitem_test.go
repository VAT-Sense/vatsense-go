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

func TestInvoiceItemGet(t *testing.T) {
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
	_, err := client.Invoice.Item.Get(
		context.TODO(),
		"in5aeae457cda2a",
		"ii5aeae457ce201",
	)
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceItemUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoice.Item.Update(
		context.TODO(),
		"in5aeae457cda2a",
		"ii5aeae457ce201",
		vatsense.InvoiceItemUpdateParams{
			InvoiceItemInput: vatsense.InvoiceItemInputParam{
				Item:         vatsense.F("Standard payment plan"),
				PriceEach:    vatsense.F(19.990000),
				Quantity:     vatsense.F(1.000000),
				VatRate:      vatsense.F(20.000000),
				DiscountRate: vatsense.F(40.000000),
			},
		},
	)
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceItemDelete(t *testing.T) {
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
	_, err := client.Invoice.Item.Delete(
		context.TODO(),
		"in5aeae457cda2a",
		"ii5aeae457ce201",
	)
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceItemAdd(t *testing.T) {
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
	_, err := client.Invoice.Item.Add(
		context.TODO(),
		"in5aeae457cda2a",
		vatsense.InvoiceItemAddParams{
			Items: vatsense.F([]vatsense.InvoiceItemInputParam{{
				Item:         vatsense.F("Standard payment plan"),
				PriceEach:    vatsense.F(19.990000),
				Quantity:     vatsense.F(1.000000),
				VatRate:      vatsense.F(20.000000),
				DiscountRate: vatsense.F(40.000000),
			}}),
		},
	)
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
