
# podcast
Package podcast generates a fully compliant iTunes and RSS 2.0 podcast feed
for GoLang using a simple API.

Full documentation with detailed examples located at <a href="https://godoc.org/github.com/eduncan911/podcast">https://godoc.org/github.com/eduncan911/podcast</a>

### Usage

see [example_test.go](example_test.go)

### Go Modules
This library is supported on GoLang 1.7 and higher.

We have implemented Go Modules support and the CI pipeline shows it working with
new installs, tested with Go 1.13.  To keep 1.7 compatibility, we use
`go mod vendor` to maintain the `vendor/` folder for older 1.7 and later runtimes.

If either runtime has an issue, please create an Issue and I will address.
