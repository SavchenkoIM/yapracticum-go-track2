module server

replace storage => ../../internal/storage

replace handlers/updateMetrics => ../../internal/handlers/updateMetrics

go 1.21.4

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	handlers/updateMetrics v0.0.0-00010101000000-000000000000 // indirect
	storage v0.0.0-00010101000000-000000000000 // indirect
)