module github.com/bredr/dapr-demo/services/operator

go 1.15

replace github.com/bredr/dapr-demo/services/common => ../common

require (
	github.com/bredr/dapr-demo/services/common v0.0.0-00010101000000-000000000000
	github.com/dapr/go-sdk v1.0.0
	github.com/google/uuid v1.1.2
)