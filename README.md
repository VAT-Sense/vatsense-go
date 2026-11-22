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

If you would rather have an answer than wait for a consultation number when VIES is down, set `ConsultationOptional: vatsense.Bool(true)`. Where a national register can confirm the number, you get the answer with no consultation number instead of the `412`.

### Check many numbers at once (batches)

A batch checks a list of VAT numbers, as the dashboard's bulk validator does with a file. Batch calls are free and each number checked counts as one request. With a test key the batch is answered from the test numbers and costs nothing.

```go
created, err := client.Batches.New(ctx, vatsense.BatchNewParams{
	Numbers:        []string{"GB288305674", "FR12345678901"},
	Reference:      vatsense.String("November customers"),
	IdempotencyKey: vatsense.String("november-customers"),
})
if err != nil {
	return err
}
id := created.Data.ID

// Poll, or take the batch.completed webhook instead.
for {
	batch, err := client.Batches.Get(ctx, id)
	if err != nil {
		return err
	}
	if s := batch.Data.Status; s == "complete" || s == "paused" || s == "expired" {
		break
	}
	time.Sleep(10 * time.Second)
}

params := vatsense.BatchResultsParams{Limit: vatsense.Int(1000)}
for {
	page, err := client.Batches.Results(ctx, id, params)
	if err != nil {
		return err
	}
	for _, row := range page.Data.Data {
		fmt.Println(row.VatNumber, row.Result, row.CompanyName)
	}
	if !page.Data.HasMore {
		break
	}
	params.After = vatsense.Int(page.Data.NextAfter)
}

csv, err := client.Batches.ResultsCsv(ctx, id)
```

Set `Rows: []vatsense.BatchRowParam{{VatNumber: ..., CompanyName: vatsense.String(...)}}` instead of `Numbers` for company matching. A batch can also be paused, continued, completed and, while a draft, deleted.

### Check a webhook's signature

Every webhook event carries a `VatSense-Signature` header. Check it against the raw request body with the endpoint's signing secret before trusting the event:

```go
http.HandleFunc("/webhooks/vatsense", func(w http.ResponseWriter, r *http.Request) {
	payload, _ := io.ReadAll(r.Body)
	event, err := vatsense.ConstructWebhookEvent(payload, r.Header.Get(vatsense.WebhookSignatureHeader), os.Getenv("VATSENSE_WEBHOOK_SECRET"))
	if err != nil {
		http.Error(w, "bad signature", http.StatusBadRequest)
		return
	}
	if event.Type == "batch.completed" {
		var batch vatsense.Batch
		_ = json.Unmarshal(event.Data.Object, &batch)
		// ...
	}
	w.WriteHeader(http.StatusOK)
})
```

Events more than 5 minutes old are refused (change it with `vatsense.WebhookOptions{Tolerance: ...}`). During the 24 hours after a secret is rolled, events signed with either secret are accepted.

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
| `client.Batches`     | Check many VAT numbers in one batch              |
| `client.Invoice`     | Create and manage VAT-compliant invoices         |
| `client.Usage`       | Check your API usage                             |

## Documentation

Full API documentation is available at [vatsense.com/documentation](https://vatsense.com/documentation).

## Versioning

This package follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions. As the library is in initial development and has a major version of `0`, APIs may change at any time.

## Contributing

See [the contributing documentation](https://github.com/VAT-Sense/vatsense-go/tree/main/CONTRIBUTING.md).
