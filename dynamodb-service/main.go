// Service 1: DynamoDB Service
package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc"

	"golang_client/dynamodb_service"
)

const (
	DynamoDBTableName = "tablename"
	Port1             = ":8081"
)

type DynamoDBService struct {
	dynamodb_service.UnimplementedDynamoDBServiceServer
	session *dynamodb.DynamoDB
}

func NewDynamoDBService() *DynamoDBService {

	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")

	creds := credentials.NewEnvCredentials()
	creds.Get()

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: creds,
	}))

	db := dynamodb.New(sess)
	return &DynamoDBService{session: db}
}

func (s *DynamoDBService) StoreData(ctx context.Context, req *dynamodb_service.StoreDataRequest) (*dynamodb_service.StoreDataResponse, error) {
	_, err := s.session.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(DynamoDBTableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String(req.GetId())},
			"data": {S: aws.String(req.GetData())},
		},
	})
	if err != nil {
		return nil, err
	}
	return &dynamodb_service.StoreDataResponse{Success: true}, nil
}

func main() {
	listener, err := net.Listen("tcp", Port1)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dynamodb_service.RegisterDynamoDBServiceServer(grpcServer, NewDynamoDBService())

	log.Printf("DynamoDB Service running on %s", Port1)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
