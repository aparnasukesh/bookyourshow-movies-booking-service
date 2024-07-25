package theatres

import "github.com/aparnasukesh/inter-communication/movie_booking"

type GrpcHandler struct {
	svc Service
	movie_booking.UnimplementedTheatreServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}
