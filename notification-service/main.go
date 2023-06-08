package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	_ "github.com/supabase-community/storage-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	. "notification_service/database"
	. "notification_service/notification"
	notificationGrpc "notification_service/proto/notification"
	"os"
	"os/signal"
	"syscall"
)

var notificationController NotificationController

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	var err error
	Client, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer Client.Close()

	// Handle WebSocket messages
	for {
		// Read message from the client
		_, message, err := Client.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}
		log.Printf("Received message: %s\n", message)
		// Send response back to the client
		err = Client.WriteMessage(websocket.TextMessage, []byte("Server received your message"))
		if err != nil {
			log.Println("Failed to send response:", err)
			break
		}
	}
}

func runWebSocketServer() {
	http.HandleFunc("/ws", handleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("WebSocket server error: ", err)
	}
}

func runGRPCServer() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("gRPC server error: ", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(CreateServerLogger()),
	)
	reflection.Register(grpcServer)

	db, err := InitDB()
	DeclareUnique(db, []UniqueField{})
	if err != nil {
		log.Fatal("Database error: ", err)
	}

	notificationRepository := &NotificationRepository{
		DB: db,
	}
	notificationService := &NotificationService{
		NotificationRepository: notificationRepository,
	}

	notificationController := NewNotificationController(notificationService)
	notificationGrpc.RegisterNotificationServiceServer(grpcServer, notificationController)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("gRPC server error: ", err)
	}
}

func main() {
	go runWebSocketServer()
	go runGRPCServer()

	fmt.Println("Server started")

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	fmt.Println("Server stopped")
}

func CreateServerLogger() grpc.UnaryServerInterceptor {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	entry := logrus.NewEntry(logger)
	return grpc_logrus.UnaryServerInterceptor(entry, grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel))
}
