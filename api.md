# Rates

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#FindRate">FindRate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#Rate">Rate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateWithTaxRate">RateWithTaxRate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#TaxRate">TaxRate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateListResponse">RateListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateCalculatePriceResponse">RateCalculatePriceResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateListTypesResponse">RateListTypesResponse</a>

Methods:

- <code title="get /rates">client.Rates.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateListParams">RateListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateListResponse">RateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/price">client.Rates.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateService.CalculatePrice">CalculatePrice</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateCalculatePriceParams">RateCalculatePriceParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateCalculatePriceResponse">RateCalculatePriceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/tax_rate">client.Rates.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateService.Details">Details</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateDetailsParams">RateDetailsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#FindRate">FindRate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/rate">client.Rates.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateFindParams">RateFindParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#FindRate">FindRate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/types">client.Rates.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateService.ListTypes">ListTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#RateListTypesResponse">RateListTypesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Countries

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#Country">Country</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryListResponse">CountryListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryListProvincesResponse">CountryListProvincesResponse</a>

Methods:

- <code title="get /countries">client.Countries.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryListParams">CountryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryListResponse">CountryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /countries/provinces">client.Countries.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryService.ListProvinces">ListProvinces</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryListProvincesParams">CountryListProvincesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CountryListProvincesResponse">CountryListProvincesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Validate

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#ValidateCheckResponse">ValidateCheckResponse</a>

Methods:

- <code title="get /validate">client.Validate.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#ValidateService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#ValidateCheckParams">ValidateCheckParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#ValidateCheckResponse">ValidateCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Currency

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#VatPrice">VatPrice</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyListResponse">CurrencyListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyCalculateVatPriceResponse">CurrencyCalculateVatPriceResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyConvertResponse">CurrencyConvertResponse</a>

Methods:

- <code title="get /currency">client.Currency.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyListParams">CurrencyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyListResponse">CurrencyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /currency/price">client.Currency.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyService.CalculateVatPrice">CalculateVatPrice</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyCalculateVatPriceParams">CurrencyCalculateVatPriceParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyCalculateVatPriceResponse">CurrencyCalculateVatPriceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /currency/convert">client.Currency.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyService.Convert">Convert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyConvertParams">CurrencyConvertParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CurrencyConvertResponse">CurrencyConvertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoice

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#CreateInvoiceParam">CreateInvoiceParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceBusinessInputParam">InvoiceBusinessInputParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceConversionInputParam">InvoiceConversionInputParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceCustomerInputParam">InvoiceCustomerInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#Invoice">Invoice</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceConversionInput">InvoiceConversionInput</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceResponse">InvoiceResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceListResponse">InvoiceListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceDeleteResponse">InvoiceDeleteResponse</a>

Methods:

- <code title="post /invoice">client.Invoice.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceNewParams">InvoiceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoice/{invoice_id}">client.Invoice.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /invoice/{invoice_id}">client.Invoice.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceUpdateParams">InvoiceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoice">client.Invoice.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceListParams">InvoiceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceListResponse">InvoiceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /invoice/{invoice_id}">client.Invoice.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceDeleteResponse">InvoiceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Item

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemInputParam">InvoiceItemInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItem">InvoiceItem</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemGetResponse">InvoiceItemGetResponse</a>

Methods:

- <code title="get /invoice/{invoice_id}/item/{item_id}">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemGetParams">InvoiceItemGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemGetResponse">InvoiceItemGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /invoice/{invoice_id}/item/{item_id}">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemUpdateParams">InvoiceItemUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /invoice/{invoice_id}/item/{item_id}">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemDeleteParams">InvoiceItemDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /invoice/{invoice_id}/item">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceItemAddParams">InvoiceItemAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Usage

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#UsageGetResponse">UsageGetResponse</a>

Methods:

- <code title="get /usage">client.Usage.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#UsageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#UsageGetResponse">UsageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sandbox

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#SandboxGenerateKeyResponse">SandboxGenerateKeyResponse</a>

Methods:

- <code title="post /sandbox/key">client.Sandbox.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#SandboxService.GenerateKey">GenerateKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/vat-sense-go#SandboxGenerateKeyResponse">SandboxGenerateKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
