module github.com/bredr/dapr-demo/services/operator

go 1.15

replace github.com/bredr/dapr-demo/services/common => ../common

require (
	github.com/bredr/dapr-demo/services/common v0.0.0-20210224220909-a987bf58711a
	github.com/dapr/go-sdk v1.0.0
	github.com/google/uuid v1.2.0
	github.com/spf13/viper v1.7.1
	go.mongodb.org/mongo-driver v1.4.6
)
