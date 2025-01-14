// Service 2: MySQL Service
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"golang_client/dynamodb_service"
	"golang_client/mysql_service"
)

const (
	MySQLDSN = "root:2c0d67537464ec621730a1da@tcp(localhost:3379)/databasename"
	Port2    = ":50052"
	HTTPPort = ":8080"
)

type MySQLService struct {
	mysql_service.UnimplementedMySQLServiceServer
	db *sql.DB
}

func NewMySQLService() *MySQLService {
	db, err := sql.Open("mysql", MySQLDSN)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}

	return &MySQLService{db: db}
}

func (s *MySQLService) StoreAndReplicateData(ctx context.Context, request *mysql_service.StoreAndReplicateRequest) (*mysql_service.StoreAndReplicateResponse, error) {
	// Store data in MySQL

	id, data := request.Id, request.Data
	_, err := s.db.Exec("INSERT INTO tablename (id, data) VALUES (?, ?)", id, data)
	if err != nil {
		return &mysql_service.StoreAndReplicateResponse{Success: err != nil}, err
	}

	// Connect to DynamoDB Service
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &mysql_service.StoreAndReplicateResponse{Success: err != nil}, err
	}
	defer conn.Close()

	dynamoClient := dynamodb_service.NewDynamoDBServiceClient(conn)
	_, err = dynamoClient.StoreData(ctx, &dynamodb_service.StoreDataRequest{
		Id:   id,
		Data: data,
	})

	return &mysql_service.StoreAndReplicateResponse{Success: err != nil}, err
}

func (s *MySQLService) handleInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload map[string]string
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	columns := []string{}
	values := []interface{}{}
	placeholders := []string{}

	for col, val := range payload {
		columns = append(columns, col)
		values = append(values, val)
		placeholders = append(placeholders, "?")
	}

	query := "INSERT INTO tablename (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ")"
	_, err := s.db.Exec(query, values...)
	if err != nil {
		http.Error(w, "Database insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Connect to DynamoDB Service
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()

	dynamoClient := dynamodb_service.NewDynamoDBServiceClient(conn)
	_, err = dynamoClient.StoreData(context.Background(), &dynamodb_service.StoreDataRequest{
		Id:   payload["id"],
		Data: payload["data"],
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Data inserted in MySQL but not in dynamoDB" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data inserted successfully"))

	return
}

func main() {
	listener, err := net.Listen("tcp", Port2)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := NewMySQLService()

	grpcServer := grpc.NewServer()
	mysql_service.RegisterMySQLServiceServer(grpcServer, s)

	go func() {
		log.Printf("MySQL Service running on %s", Port2)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	http.HandleFunc("/insert", s.handleInsert)
	log.Printf("HTTP Server running on %s", HTTPPort)
	if err := http.ListenAndServe("localhost"+HTTPPort, nil); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
