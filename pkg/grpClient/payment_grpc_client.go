package grpclient

import (
	"log"

	pb "github.com/aparnasukesh/inter-communication/payment"

	"google.golang.org/grpc"
)

func NewBookingPaymentServiceClient(port string) (pb.PaymentServiceClient, error) {
	// conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	// if err != nil {
	// 	return nil, err
	// }
	address := "payment-svc.default.svc.cluster.local:" + port
	serviceConfig := `{"loadBalancingPolicy": "round_robin"}`
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(serviceConfig))
	if err != nil {
		log.Printf("Failed to connect to gRPC service: %v", err)
		return nil, err
	}
	return pb.NewPaymentServiceClient(conn), nil
}
