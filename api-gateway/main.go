package main

import (
	"context"
	"gateway/proto/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	authConn, err := ConnectToService("AUTH_SERVICE_ADDRESS")
	accomodationConn, err := ConnectToService("ACCOMODATION_SERVICE_ADDRESS")

	//conn, err := grpc.DialContext(
	//	context.Background(),
	//	os.Getenv(),
	//	grpc.WithBlock(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)

	//conn1, err1 := grpc.DialContext(
	//	context.Background(),
	//	os.Getenv("ACCOMODATION_SERVICE_ADDRESS"),
	//	grpc.WithBlock(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	client := gateway.NewAuthServiceClient(authConn)
	err = gateway.RegisterAuthServiceHandlerClient(
		context.Background(),
		gwmux,
		client,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	client1 := gateway.NewAccomodationServiceClient(accomodationConn)
	err = gateway.RegisterAccomodationServiceHandlerClient(
		context.Background(),
		gwmux,
		client1,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":" + os.Getenv("GATEWAY_ADDRESS"),
		Handler: gwmux,
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}

func ConnectToService(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		os.Getenv(address),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn, err
}
