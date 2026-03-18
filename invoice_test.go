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
			Business: vatsense.F(vatsense.InvoiceBusinessInputParam{
				Address:       vatsense.F("123 Example Street\nLondon\nSW3 1GL\nUnited Kingdom"),
				Name:          vatsense.F("VAT Sense"),
				VatNumber:     vatsense.F("GB12345678"),
				BankAccount:   vatsense.F("bank_account"),
				CompanyNumber: vatsense.F("9839222"),
				Email:         vatsense.F("dev@stainless.com"),
				Logo:          vatsense.F("https://example.com"),
				Phone:         vatsense.F("phone"),
				Website:       vatsense.F("https://example.com"),
			}),
			CurrencyCode: vatsense.F("USD"),
			Date:         vatsense.F("2018-06-03 14:02:00"),
			Items: vatsense.F([]vatsense.InvoiceItemInputParam{{
				Item:         vatsense.F("Standard payment plan"),
				PriceEach:    vatsense.F(19.990000),
				Quantity:     vatsense.F(1.000000),
				VatRate:      vatsense.F(20.000000),
				DiscountRate: vatsense.F(40.000000),
			}}),
			TaxPoint: vatsense.F("2018-06-03 14:02:00"),
			Conversion: vatsense.F(vatsense.InvoiceConversionInputParam{
				CurrencyCode: vatsense.F("GBP"),
				Rate:         vatsense.F(1.523000),
			}),
			Customer: vatsense.F(vatsense.InvoiceCustomerInputParam{
				Name:          vatsense.F("Demo Co."),
				Address:       vatsense.F("65 Demo Road\nLondon\nSW1 3DE\nUnited Kingdom"),
				CompanyNumber: vatsense.F("5584922"),
				CountryCode:   vatsense.F("country_code"),
				Email:         vatsense.F("dev@stainless.com"),
				Logo:          vatsense.F("https://example.com"),
				VatNumber:     vatsense.F("GB912343332"),
			}),
			HasVat:           vatsense.F(true),
			InvoiceNumber:    vatsense.F("203"),
			IsCopy:           vatsense.F(true),
			IsReverseCharge:  vatsense.F(true),
			Notes:            vatsense.F("notes"),
			PadInvoiceNumber: vatsense.F(int64(2)),
			Serial:           vatsense.F("serial"),
			TaxType:          vatsense.F(vatsense.CreateInvoiceTaxTypeIncl),
			Type:             vatsense.F(vatsense.CreateInvoiceTypeSale),
			ZeroRated:        vatsense.F(true),
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
				Business: vatsense.F(vatsense.InvoiceBusinessInputParam{
					Address:       vatsense.F("123 Example Street\nLondon\nSW3 1GL\nUnited Kingdom"),
					Name:          vatsense.F("VAT Sense"),
					VatNumber:     vatsense.F("GB12345678"),
					BankAccount:   vatsense.F("bank_account"),
					CompanyNumber: vatsense.F("9839222"),
					Email:         vatsense.F("dev@stainless.com"),
					Logo:          vatsense.F("https://example.com"),
					Phone:         vatsense.F("phone"),
					Website:       vatsense.F("https://example.com"),
				}),
				CurrencyCode: vatsense.F("USD"),
				Date:         vatsense.F("2018-06-03 14:02:00"),
				Items: vatsense.F([]vatsense.InvoiceItemInputParam{{
					Item:         vatsense.F("Standard payment plan"),
					PriceEach:    vatsense.F(19.990000),
					Quantity:     vatsense.F(1.000000),
					VatRate:      vatsense.F(20.000000),
					DiscountRate: vatsense.F(40.000000),
				}}),
				TaxPoint: vatsense.F("2018-06-03 14:02:00"),
				Conversion: vatsense.F(vatsense.InvoiceConversionInputParam{
					CurrencyCode: vatsense.F("GBP"),
					Rate:         vatsense.F(1.523000),
				}),
				Customer: vatsense.F(vatsense.InvoiceCustomerInputParam{
					Name:          vatsense.F("Demo Co."),
					Address:       vatsense.F("65 Demo Road\nLondon\nSW1 3DE\nUnited Kingdom"),
					CompanyNumber: vatsense.F("5584922"),
					CountryCode:   vatsense.F("country_code"),
					Email:         vatsense.F("dev@stainless.com"),
					Logo:          vatsense.F("https://example.com"),
					VatNumber:     vatsense.F("GB912343332"),
				}),
				HasVat:           vatsense.F(true),
				InvoiceNumber:    vatsense.F("203"),
				IsCopy:           vatsense.F(true),
				IsReverseCharge:  vatsense.F(true),
				Notes:            vatsense.F("notes"),
				PadInvoiceNumber: vatsense.F(int64(2)),
				Serial:           vatsense.F("serial"),
				TaxType:          vatsense.F(vatsense.CreateInvoiceTaxTypeIncl),
				Type:             vatsense.F(vatsense.CreateInvoiceTypeSale),
				ZeroRated:        vatsense.F(true),
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
		Limit:  vatsense.F(int64(1)),
		Offset: vatsense.F(int64(0)),
		Search: vatsense.F("search"),
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
