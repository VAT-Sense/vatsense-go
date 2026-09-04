# Changelog

## 0.2.0 (2026-09-04)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/VAT-Sense/vatsense-go/compare/v0.1.0...v0.2.0)

### Features

* **client:** optimize json encoder for internal types ([dafa95d](https://github.com/VAT-Sense/vatsense-go/commit/dafa95d7639ef18acad2f770fffba3059d1d8a9e))
* **go:** add default http client with timeout ([0889248](https://github.com/VAT-Sense/vatsense-go/commit/088924873ef59c10ea1746723731b39991b6b91d))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([ba05b9a](https://github.com/VAT-Sense/vatsense-go/commit/ba05b9ae16d7ce83c7a199bd3a9686977a36b0a7))
* support setting headers via env ([bbe9372](https://github.com/VAT-Sense/vatsense-go/commit/bbe9372e77a48c3f9f94081ad7c035742ada8ad9))


### Bug Fixes

* fix issue with unmarshaling in some cases ([7c7a1f2](https://github.com/VAT-Sense/vatsense-go/commit/7c7a1f2aa6c9d390c7cab8de232be3eeb4ef6595))
* **go:** avoid panic when http.DefaultTransport is wrapped ([eb7681e](https://github.com/VAT-Sense/vatsense-go/commit/eb7681ea6ff8e775ea0e729db3cc1805f2998b7d))


### Chores

* avoid embedding reflect.Type for dead code elimination ([ec03883](https://github.com/VAT-Sense/vatsense-go/commit/ec03883685f76d5a29877df336c2447dd07a9803))
* **internal:** more robust bootstrap script ([6a9a62f](https://github.com/VAT-Sense/vatsense-go/commit/6a9a62f7fd6bbde8945af2ec8403cc228f8a8147))
* **internal:** version bump ([fcc327e](https://github.com/VAT-Sense/vatsense-go/commit/fcc327ec862a6d7ec1c8a1bcfc4fdc848356d2b6))
* redact api-key headers in debug logs ([5b4cc51](https://github.com/VAT-Sense/vatsense-go/commit/5b4cc515e4a523d2b5bf1c57f6aa7322119f74ad))
* update docs for api:"required" ([3f290bc](https://github.com/VAT-Sense/vatsense-go/commit/3f290bc98013d1015fbe3cf29c19a5af0cba8d7f))

## 0.1.0 (2026-03-28)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/VAT-Sense/vatsense-go/compare/v0.0.2...v0.1.0)

### Features

* **api:** api update ([b981eb3](https://github.com/VAT-Sense/vatsense-go/commit/b981eb3b7bd6d66e5ac1088d28228a82cb871f85))
* **api:** api update ([b17827b](https://github.com/VAT-Sense/vatsense-go/commit/b17827b9a595f5bde69a37c487a28c2ce3f11ac3))
* **internal:** support comma format in multipart form encoding ([fbe94eb](https://github.com/VAT-Sense/vatsense-go/commit/fbe94ebd81ca86a80ec467c6930f4bd256d691a7))


### Bug Fixes

* prevent duplicate ? in query params ([fb63e53](https://github.com/VAT-Sense/vatsense-go/commit/fb63e53f786d862d9cddcc1b58c658e37fd7b257))


### Chores

* **ci:** skip lint on metadata-only changes ([9f7aa7f](https://github.com/VAT-Sense/vatsense-go/commit/9f7aa7ff5cad4545258d04eabf2c5af38603b226))
* **ci:** support opting out of skipping builds on metadata-only commits ([df4083e](https://github.com/VAT-Sense/vatsense-go/commit/df4083e45596123e547f188a2d200f3616fa499a))
* **client:** fix multipart serialisation of Default() fields ([8c82987](https://github.com/VAT-Sense/vatsense-go/commit/8c82987195196b4c5db9af130d164c012d3bbe25))
* **internal:** support default value struct tag ([07e806d](https://github.com/VAT-Sense/vatsense-go/commit/07e806d6445acdfa928ffece568b0d802483d4e1))
* **internal:** update gitignore ([a668a2c](https://github.com/VAT-Sense/vatsense-go/commit/a668a2cab2c2d16c2cb2300edd6b2329ef985dbf))
* remove unnecessary error check for url parsing ([dd96877](https://github.com/VAT-Sense/vatsense-go/commit/dd96877234b58c5f68d2712974d0acf0b749b108))


### Documentation

* **readme:** tailor README to VAT Sense use cases ([f330cf6](https://github.com/VAT-Sense/vatsense-go/commit/f330cf6ead0d6c294043a1b7d28f321998614025))

## 0.0.2 (2026-03-18)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/VAT-Sense/vatsense-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([f18b1a8](https://github.com/VAT-Sense/vatsense-go/commit/f18b1a88fe896afb6d02bcf912ae8dae31edd489))
* update SDK settings ([6f421cd](https://github.com/VAT-Sense/vatsense-go/commit/6f421cd485061935270b9a7aea4e1122eed9d235))
