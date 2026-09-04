# VAT Sense Go SDK

<!-- x-release-please-start-version -->

<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go"><img src="https://pkg.go.dev/badge/github.com/VAT-Sense/vatsense-go.svg" alt="Go Reference"></a>

<!-- x-release-please-end -->

The official Go library for the [VAT Sense](https://vatsense.com) REST API. Validate VAT/EORI numbers, look up VAT/GST rates, calculate prices, convert currencies, and generate VAT-compliant invoices.

## Installation

```sh
go get -u 'github.com/VAT-Sense/vatsense-go@v0.1.0'
```

## Quick start

Create a client using your API key from the [VAT Sense dashboard](https://vatsense.com/dashboard). The API uses HTTP Basic Auth with `user` as the username and your API key as the password.

```go
package main

import (
	"context"
	"fmt"

	"github.com/VAT-Sense/vatsense-go"
	"github.com/VAT-Sense/vatsense-go/option"
)

func main() {
	client := vatsense.NewClient(
		option.WithUsername("user"),
		option.WithPassword("your_api_key"),
	)
}
```

You can also set the `VAT_SENSE_USERNAME` and `VAT_SENSE_PASSWORD` environment variables and the client will pick them up automatically.

### Validate a VAT number

```go
response, err := client.Validate.Check(context.TODO(), vatsense.ValidateCheckParams{
	VatNumber: vatsense.String("GB288305674"),
})
if err != nil {
	panic(err)
}

if response.Data.Valid {
	fmt.Println(response.Data.Company.AsValidateCheckResponseDataCompanyValidationCompany().CompanyName)
	fmt.Println(response.Data.Company.AsValidateCheckResponseDataCompanyValidationCompany().CompanyAddress)
	fmt.Println(response.Data.Company.AsValidateCheckResponseDataCompanyValidationCompany().CountryCode)
}
```

VAT validation works for the UK, EU, Australia, Norway, Switzerland, South Africa, and Brazil.

### Get a consultation number

If you need an official consultation number from VIES (EU) or HMRC (UK), provide your own VAT number as the requester:

```go
response, err := client.Validate.Check(context.TODO(), vatsense.ValidateCheckParams{
	VatNumber:          vatsense.String("FR12345678901"),
	RequesterVatNumber: vatsense.String("FR98765432101"),
})

fmt.Println(response.Data.ConsultationNumber)
```

> **Note:** GB requester numbers only work for GB validations, and EU requester numbers only work for EU validations. Cross-region requests are not supported.

### Find the VAT rate for a country

```go
rate, err := client.Rates.Find(context.TODO(), vatsense.RateFindParams{
	CountryCode: vatsense.String("DE"),
})

fmt.Println(rate.Data.CountryName)    // "Germany"
fmt.Println(rate.Data.TaxRate.Rate)   // 19
fmt.Println(rate.Data.TaxRate.Class)  // "standard"
```

### Find a rate for a specific product type

```go
rate, err := client.Rates.Find(context.TODO(), vatsense.RateFindParams{
	CountryCode: vatsense.String("DE"),
	Type:        vatsense.String("ebooks"),
})

fmt.Println(rate.Data.TaxRate.Rate)   // 7
fmt.Println(rate.Data.TaxRate.Class)  // "reduced"
```

### Find a rate by IP address

Useful for determining the correct rate based on your customer's location:

```go
rate, err := client.Rates.Find(context.TODO(), vatsense.RateFindParams{
	IPAddress: vatsense.String("185.86.151.11"),
})

fmt.Println(rate.Data.CountryCode)    // "GB"
fmt.Println(rate.Data.TaxRate.Rate)   // 20
```

### Calculate a VAT-inclusive price

```go
result, err := client.Rates.CalculatePrice(context.TODO(), vatsense.RateCalculatePriceParams{
	Price:       "100.00",
	TaxType:     vatsense.RateCalculatePriceParamsTaxTypeExcl,
	CountryCode: vatsense.String("FR"),
})

fmt.Println(result.Data.VatPrice.PriceInclVat)  // Price including VAT
fmt.Println(result.Data.VatPrice.PriceExclVat)  // Price excluding VAT
fmt.Println(result.Data.VatPrice.VatRate)        // VAT rate applied
fmt.Println(result.Data.VatPrice.Vat)            // VAT amount
```

### List all VAT rates

```go
rates, err := client.Rates.List(context.TODO(), vatsense.RateListParams{})

for _, rate := range rates.Data {
	fmt.Printf("%s: %s\n", rate.CountryCode, rate.CountryName)
}

// Filter to EU countries only
euRates, err := client.Rates.List(context.TODO(), vatsense.RateListParams{
	Eu: vatsense.Bool(true),
})
```

## Handling errors

When the API returns an error, the library returns it as an `*vatsense.Error`:

```go
response, err := client.Validate.Check(context.TODO(), vatsense.ValidateCheckParams{
	VatNumber: vatsense.String("GB288305674"),
})
if err != nil {
	var apiErr *vatsense.Error
	if errors.As(err, &apiErr) {
		fmt.Println(apiErr.StatusCode)
		fmt.Println(apiErr.Message)
	}
}
```

A `412` error means the upstream validation service (VIES, HMRC, etc.) is temporarily unavailable. These requests do not count against your usage quota.

## Retries

Failed requests are automatically retried up to 2 times with exponential backoff. This includes connection errors, timeouts, 429, and 5xx responses.

```go
// Disable retries
client := vatsense.NewClient(
	option.WithUsername("user"),
	option.WithPassword("your_api_key"),
	option.WithMaxRetries(0),
)

// Or configure per request
response, err := client.Validate.Check(
	context.TODO(),
	vatsense.ValidateCheckParams{VatNumber: vatsense.String("GB288305674")},
	option.WithMaxRetries(5),
)
```

## Available services

| Service              | Description                                     |
| -------------------- | ----------------------------------------------- |
| `client.Validate`    | Validate VAT and EORI numbers                   |
| `client.Rates`       | VAT/GST rate lookups, price calculations         |
| `client.Countries`   | Country data and province lookups                |
| `client.Currency`    | Exchange rates and currency conversion           |
| `client.Invoice`     | Create and manage VAT-compliant invoices         |
| `client.Usage`       | Check your API usage                             |

## Documentation

Full API documentation is available at [vatsense.com/documentation](https://vatsense.com/documentation).

## Versioning

This package follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions. As the library is in initial development and has a major version of `0`, APIs may change at any time.

## Contributing

See [the contributing documentation](https://github.com/VAT-Sense/vatsense-go/tree/main/CONTRIBUTING.md).
