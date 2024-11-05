package grpclient

import (
	pb "github.com/aparnasukesh/inter-communication/payment"

	"google.golang.org/grpc"
)

func NewBookingPaymentServiceClient(port string) (pb.PaymentServiceClient, error) {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewPaymentServiceClient(conn), nil
}
