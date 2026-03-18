# Rates

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#FindRate">FindRate</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#Rate">Rate</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateWithTaxRate">RateWithTaxRate</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#TaxRate">TaxRate</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateListResponse">RateListResponse</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateCalculatePriceResponse">RateCalculatePriceResponse</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateListTypesResponse">RateListTypesResponse</a>

Methods:

- <code title="get /rates">client.Rates.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateListParams">RateListParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateListResponse">RateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/price">client.Rates.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateService.CalculatePrice">CalculatePrice</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateCalculatePriceParams">RateCalculatePriceParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateCalculatePriceResponse">RateCalculatePriceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/tax_rate">client.Rates.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateService.Details">Details</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateDetailsParams">RateDetailsParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#FindRate">FindRate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/rate">client.Rates.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateFindParams">RateFindParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#FindRate">FindRate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rates/types">client.Rates.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateService.ListTypes">ListTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#RateListTypesResponse">RateListTypesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Countries

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#Country">Country</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryListResponse">CountryListResponse</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryListProvincesResponse">CountryListProvincesResponse</a>

Methods:

- <code title="get /countries">client.Countries.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryListParams">CountryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryListResponse">CountryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /countries/provinces">client.Countries.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryService.ListProvinces">ListProvinces</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryListProvincesParams">CountryListProvincesParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CountryListProvincesResponse">CountryListProvincesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Validate

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#ValidateCheckResponse">ValidateCheckResponse</a>

Methods:

- <code title="get /validate">client.Validate.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#ValidateService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#ValidateCheckParams">ValidateCheckParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#ValidateCheckResponse">ValidateCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Currency

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#VatPrice">VatPrice</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyListResponse">CurrencyListResponse</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyCalculateVatPriceResponse">CurrencyCalculateVatPriceResponse</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyConvertResponse">CurrencyConvertResponse</a>

Methods:

- <code title="get /currency">client.Currency.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyListParams">CurrencyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyListResponse">CurrencyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /currency/price">client.Currency.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyService.CalculateVatPrice">CalculateVatPrice</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyCalculateVatPriceParams">CurrencyCalculateVatPriceParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyCalculateVatPriceResponse">CurrencyCalculateVatPriceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /currency/convert">client.Currency.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyService.Convert">Convert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyConvertParams">CurrencyConvertParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CurrencyConvertResponse">CurrencyConvertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoice

Params Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#CreateInvoiceParam">CreateInvoiceParam</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceBusinessInputParam">InvoiceBusinessInputParam</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceConversionInputParam">InvoiceConversionInputParam</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceCustomerInputParam">InvoiceCustomerInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#Invoice">Invoice</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceConversionInput">InvoiceConversionInput</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceResponse">InvoiceResponse</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceListResponse">InvoiceListResponse</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceDeleteResponse">InvoiceDeleteResponse</a>

Methods:

- <code title="post /invoice">client.Invoice.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceNewParams">InvoiceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoice/{invoice_id}">client.Invoice.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /invoice/{invoice_id}">client.Invoice.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceUpdateParams">InvoiceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoice">client.Invoice.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceListParams">InvoiceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceListResponse">InvoiceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /invoice/{invoice_id}">client.Invoice.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceDeleteResponse">InvoiceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Item

Params Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemInputParam">InvoiceItemInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItem">InvoiceItem</a>
- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemGetResponse">InvoiceItemGetResponse</a>

Methods:

- <code title="get /invoice/{invoice_id}/item/{item_id}">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemGetParams">InvoiceItemGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemGetResponse">InvoiceItemGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /invoice/{invoice_id}/item/{item_id}">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemUpdateParams">InvoiceItemUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /invoice/{invoice_id}/item/{item_id}">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemDeleteParams">InvoiceItemDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /invoice/{invoice_id}/item">client.Invoice.Item.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceItemAddParams">InvoiceItemAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Usage

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#UsageGetResponse">UsageGetResponse</a>

Methods:

- <code title="get /usage">client.Usage.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#UsageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#UsageGetResponse">UsageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sandbox

Response Types:

- <a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#SandboxGenerateKeyResponse">SandboxGenerateKeyResponse</a>

Methods:

- <code title="post /sandbox/key">client.Sandbox.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#SandboxService.GenerateKey">GenerateKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go">vatsense</a>.<a href="https://pkg.go.dev/github.com/VAT-Sense/vatsense-go#SandboxGenerateKeyResponse">SandboxGenerateKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
