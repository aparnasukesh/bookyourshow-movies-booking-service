package theatres

import (
	"context"

	"github.com/aparnasukesh/inter-communication/movie_booking"
)

type GrpcHandler struct {
	svc Service
	movie_booking.UnimplementedTheatreServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (h *GrpcHandler) AddTheaterType(ctx context.Context, req *movie_booking.AddTheaterTypeRequest) (*movie_booking.AddTheaterTypeResponse, error) {
	if err := h.svc.AddTheaterType(ctx, TheaterType{
		TheaterTypeName: req.TheaterTypeName,
	}); err != nil {
		return &movie_booking.AddTheaterTypeResponse{
			Message: "failed to add theater type",
		}, err
	}
	return &movie_booking.AddTheaterTypeResponse{
		Message: "successfully added theater type",
	}, nil
}

func (h *GrpcHandler) DeleteTheaterTypeByID(ctx context.Context, req *movie_booking.DeleteTheaterTypeRequest) (*movie_booking.DeleteTheaterTypeResponse, error) {
	if err := h.svc.DeleteTheaterTypeByID(ctx, int(req.TheaterTypeId)); err != nil {
		return &movie_booking.DeleteTheaterTypeResponse{
			Message: "failed to delete theater type",
		}, err
	}
	return &movie_booking.DeleteTheaterTypeResponse{
		Message: "successfully deleted theater type",
	}, nil
}

func (h *GrpcHandler) DeleteTheaterTypeByName(ctx context.Context, req *movie_booking.DeleteTheaterTypeByNameRequest) (*movie_booking.DeleteTheaterTypeByNameResponse, error) {
	if err := h.svc.DeleteTheaterTypeByName(ctx, req.Name); err != nil {
		return &movie_booking.DeleteTheaterTypeByNameResponse{
			Message: "failed to delete theater type",
		}, err
	}
	return &movie_booking.DeleteTheaterTypeByNameResponse{
		Message: "successfully deleted theater type",
	}, nil
}

func (h *GrpcHandler) GetTheaterTypeByID(ctx context.Context, req *movie_booking.GetTheaterTypeByIDRequest) (*movie_booking.GetTheaterTypeByIDResponse, error) {
	theaterType, err := h.svc.GetTheaterTypeByID(ctx, int(req.TheaterTypeId))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetTheaterTypeByIDResponse{
		TheaterType: &movie_booking.TheaterType{
			Id:              int32(theaterType.ID),
			TheaterTypeName: theaterType.TheaterTypeName,
		},
	}, nil
}

func (h *GrpcHandler) GetTheaterTypeByName(ctx context.Context, req *movie_booking.GetTheaterTypeByNameRequest) (*movie_booking.GetTheaterTypeBynameResponse, error) {
	theaterType, err := h.svc.GetTheaterTypeByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetTheaterTypeBynameResponse{
		TheaterType: &movie_booking.TheaterType{
			Id:              int32(theaterType.ID),
			TheaterTypeName: theaterType.TheaterTypeName,
		},
	}, nil
}

func (h *GrpcHandler) UpdateTheaterType(ctx context.Context, req *movie_booking.UpdateTheaterTypeRequest) (*movie_booking.UpdateTheaterTypeResponse, error) {
	err := h.svc.UpdateTheaterType(ctx, int(req.Id), TheaterType{
		TheaterTypeName: req.TheaterTypeName,
	})
	if err != nil {
		return nil, err
	}
	return &movie_booking.UpdateTheaterTypeResponse{}, nil
}

func (h *GrpcHandler) ListTheaterTypes(ctx context.Context, req *movie_booking.ListTheaterTypesRequest) (*movie_booking.ListTheaterTypeResponse, error) {
	response, err := h.svc.ListTheaterTypes(ctx)
	if err != nil {
		return nil, err
	}

	var grpcTheaterTypes []*movie_booking.TheaterType
	for _, m := range response {
		grpcTheaterType := &movie_booking.TheaterType{
			Id:              int32(m.ID),
			TheaterTypeName: m.TheaterTypeName,
		}
		grpcTheaterTypes = append(grpcTheaterTypes, grpcTheaterType)
	}

	return &movie_booking.ListTheaterTypeResponse{
		TheaterTypes: grpcTheaterTypes,
	}, nil
}
