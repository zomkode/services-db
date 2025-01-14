module dynamodb-service

go 1.22.10

require (
	github.com/aws/aws-sdk-go v1.55.5
	golang_client/dynamodb_service v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.69.4
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/protobuf v1.36.2 // indirect
)

replace golang_client/mysql_service => ../golang_client/mysql_service

replace golang_client/dynamodb_service => ../golang_client/dynamodb_service
