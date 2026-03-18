// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vatsense_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/vat-sense-go"
	"github.com/stainless-sdks/vat-sense-go/internal/testutil"
	"github.com/stainless-sdks/vat-sense-go/option"
)

func TestInvoiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoice.New(context.TODO(), vatsense.InvoiceNewParams{
		CreateInvoice: vatsense.CreateInvoiceParam{
			Business: vatsense.InvoiceBusinessInputParam{
				Address:       "123 Example Street\nLondon\nSW3 1GL\nUnited Kingdom",
				Name:          "VAT Sense",
				VatNumber:     "GB12345678",
				BankAccount:   vatsense.String("bank_account"),
				CompanyNumber: vatsense.String("9839222"),
				Email:         vatsense.String("dev@stainless.com"),
				Logo:          vatsense.String("https://example.com"),
				Phone:         vatsense.String("phone"),
				Website:       vatsense.String("https://example.com"),
			},
			CurrencyCode: "USD",
			Date:         "2018-06-03 14:02:00",
			Items: []vatsense.InvoiceItemInputParam{{
				Item:         "Standard payment plan",
				PriceEach:    19.99,
				Quantity:     1,
				VatRate:      20,
				DiscountRate: vatsense.Float(40),
			}},
			TaxPoint: "2018-06-03 14:02:00",
			Conversion: vatsense.InvoiceConversionInputParam{
				CurrencyCode: "GBP",
				Rate:         1.523,
			},
			Customer: vatsense.InvoiceCustomerInputParam{
				Name:          "Demo Co.",
				Address:       vatsense.String("65 Demo Road\nLondon\nSW1 3DE\nUnited Kingdom"),
				CompanyNumber: vatsense.String("5584922"),
				CountryCode:   vatsense.String("country_code"),
				Email:         vatsense.String("dev@stainless.com"),
				Logo:          vatsense.String("https://example.com"),
				VatNumber:     vatsense.String("GB912343332"),
			},
			HasVat:           vatsense.Bool(true),
			InvoiceNumber:    vatsense.String("203"),
			IsCopy:           vatsense.Bool(true),
			IsReverseCharge:  vatsense.Bool(true),
			Notes:            vatsense.String("notes"),
			PadInvoiceNumber: vatsense.Int(2),
			Serial:           vatsense.String("serial"),
			TaxType:          vatsense.CreateInvoiceTaxTypeIncl,
			Type:             vatsense.CreateInvoiceTypeSale,
			ZeroRated:        vatsense.Bool(true),
		},
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceGet(t *testing.T) {
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
	_, err := client.Invoice.Get(context.TODO(), "in5aeae457cda2a")
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoice.Update(
		context.TODO(),
		"in5aeae457cda2a",
		vatsense.InvoiceUpdateParams{
			CreateInvoice: vatsense.CreateInvoiceParam{
				Business: vatsense.InvoiceBusinessInputParam{
					Address:       "123 Example Street\nLondon\nSW3 1GL\nUnited Kingdom",
					Name:          "VAT Sense",
					VatNumber:     "GB12345678",
					BankAccount:   vatsense.String("bank_account"),
					CompanyNumber: vatsense.String("9839222"),
					Email:         vatsense.String("dev@stainless.com"),
					Logo:          vatsense.String("https://example.com"),
					Phone:         vatsense.String("phone"),
					Website:       vatsense.String("https://example.com"),
				},
				CurrencyCode: "USD",
				Date:         "2018-06-03 14:02:00",
				Items: []vatsense.InvoiceItemInputParam{{
					Item:         "Standard payment plan",
					PriceEach:    19.99,
					Quantity:     1,
					VatRate:      20,
					DiscountRate: vatsense.Float(40),
				}},
				TaxPoint: "2018-06-03 14:02:00",
				Conversion: vatsense.InvoiceConversionInputParam{
					CurrencyCode: "GBP",
					Rate:         1.523,
				},
				Customer: vatsense.InvoiceCustomerInputParam{
					Name:          "Demo Co.",
					Address:       vatsense.String("65 Demo Road\nLondon\nSW1 3DE\nUnited Kingdom"),
					CompanyNumber: vatsense.String("5584922"),
					CountryCode:   vatsense.String("country_code"),
					Email:         vatsense.String("dev@stainless.com"),
					Logo:          vatsense.String("https://example.com"),
					VatNumber:     vatsense.String("GB912343332"),
				},
				HasVat:           vatsense.Bool(true),
				InvoiceNumber:    vatsense.String("203"),
				IsCopy:           vatsense.Bool(true),
				IsReverseCharge:  vatsense.Bool(true),
				Notes:            vatsense.String("notes"),
				PadInvoiceNumber: vatsense.Int(2),
				Serial:           vatsense.String("serial"),
				TaxType:          vatsense.CreateInvoiceTaxTypeIncl,
				Type:             vatsense.CreateInvoiceTypeSale,
				ZeroRated:        vatsense.Bool(true),
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

func TestInvoiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoice.List(context.TODO(), vatsense.InvoiceListParams{
		Limit:  vatsense.Int(1),
		Offset: vatsense.Int(0),
		Search: vatsense.String("search"),
	})
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceDelete(t *testing.T) {
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
	_, err := client.Invoice.Delete(context.TODO(), "in5aeae457cda2a")
	if err != nil {
		var apierr *vatsense.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
