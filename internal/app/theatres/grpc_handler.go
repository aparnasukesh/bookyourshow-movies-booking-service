package theatres

import (
	"context"

	"github.com/aparnasukesh/inter-communication/movie_booking"
	"google.golang.org/protobuf/types/known/timestamppb"
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

// theater type
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

// screen type
func (h *GrpcHandler) AddScreenType(ctx context.Context, req *movie_booking.AddScreenTypeRequest) (*movie_booking.AddScreenTypeResponse, error) {
	if err := h.svc.AddScreenType(ctx, ScreenType{
		ScreenTypeName: req.ScreenTypeName,
	}); err != nil {
		return &movie_booking.AddScreenTypeResponse{
			Message: "failed to add screen type",
		}, err
	}
	return &movie_booking.AddScreenTypeResponse{
		Message: "successfully added screen type",
	}, nil
}

func (h *GrpcHandler) DeleteScreenTypeByID(ctx context.Context, req *movie_booking.DeleteScreenTypeRequest) (*movie_booking.DeleteScreenTypeResponse, error) {
	if err := h.svc.DeleteScreenTypeByID(ctx, int(req.ScreenTypeId)); err != nil {
		return &movie_booking.DeleteScreenTypeResponse{
			Message: "failed to delete screen type",
		}, err
	}
	return &movie_booking.DeleteScreenTypeResponse{
		Message: "successfully deleted screen type",
	}, nil
}

func (h *GrpcHandler) DeleteScreenTypeByName(ctx context.Context, req *movie_booking.DeleteScreenTypeByNameRequest) (*movie_booking.DeleteScreenTypeByNameResponse, error) {
	if err := h.svc.DeleteScreenTypeByName(ctx, req.Name); err != nil {
		return &movie_booking.DeleteScreenTypeByNameResponse{
			Message: "failed to delete screen type",
		}, err
	}
	return &movie_booking.DeleteScreenTypeByNameResponse{
		Message: "successfully deleted screen type",
	}, nil
}

func (h *GrpcHandler) GetScreenTypeByID(ctx context.Context, req *movie_booking.GetScreenTypeByIDRequest) (*movie_booking.GetScreenTypeByIDResponse, error) {
	screenType, err := h.svc.GetScreenTypeByID(ctx, int(req.ScreenTypeId))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetScreenTypeByIDResponse{
		ScreenType: &movie_booking.ScreenType{
			Id:             int32(screenType.ID),
			ScreenTypeName: screenType.ScreenTypeName,
		},
	}, nil
}

func (h *GrpcHandler) GetScreenTypeByName(ctx context.Context, req *movie_booking.GetScreenTypeByNameRequest) (*movie_booking.GetScreenTypeByNameResponse, error) {
	screenType, err := h.svc.GetScreenTypeByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetScreenTypeByNameResponse{
		ScreenType: &movie_booking.ScreenType{
			Id:             int32(screenType.ID),
			ScreenTypeName: screenType.ScreenTypeName,
		},
	}, nil
}

func (h *GrpcHandler) UpdateScreenType(ctx context.Context, req *movie_booking.UpdateScreenTypeRequest) (*movie_booking.UpdateScreenTypeResponse, error) {
	err := h.svc.UpdateScreenType(ctx, int(req.Id), ScreenType{
		ScreenTypeName: req.ScreenTypeName,
	})
	if err != nil {
		return nil, err
	}
	return &movie_booking.UpdateScreenTypeResponse{}, nil
}

func (h *GrpcHandler) ListScreenTypes(ctx context.Context, req *movie_booking.ListScreenTypesRequest) (*movie_booking.ListScreenTypesResponse, error) {
	response, err := h.svc.ListScreenTypes(ctx)
	if err != nil {
		return nil, err
	}

	var grpcScreenTypes []*movie_booking.ScreenType
	for _, m := range response {
		grpcScreenType := &movie_booking.ScreenType{
			Id:             int32(m.ID),
			ScreenTypeName: m.ScreenTypeName,
		}
		grpcScreenTypes = append(grpcScreenTypes, grpcScreenType)
	}

	return &movie_booking.ListScreenTypesResponse{
		ScreenTypes: grpcScreenTypes,
	}, nil
}

// Seat categories
func (h *GrpcHandler) AddSeatCategory(ctx context.Context, req *movie_booking.AddSeatCategoryRequest) (*movie_booking.AddSeatCategoryResponse, error) {
	if err := h.svc.AddSeatCategory(ctx, SeatCategory{
		SeatCategoryName: req.SeatCategory.SeatCategoryName,
	}); err != nil {
		return &movie_booking.AddSeatCategoryResponse{
			Message: "failed to add seat category",
		}, err
	}
	return &movie_booking.AddSeatCategoryResponse{
		Message: "successfully added seat category",
	}, nil
}

func (h *GrpcHandler) DeleteSeatCategoryByID(ctx context.Context, req *movie_booking.DeleteSeatCategoryRequest) (*movie_booking.DeleteSeatCategoryResponse, error) {
	if err := h.svc.DeleteSeatCategoryByID(ctx, int(req.SeatCategoryId)); err != nil {
		return &movie_booking.DeleteSeatCategoryResponse{
			Message: "failed to delete seat category",
		}, err
	}
	return &movie_booking.DeleteSeatCategoryResponse{
		Message: "successfully deleted seat category",
	}, nil
}

func (h *GrpcHandler) DeleteSeatCategoryByName(ctx context.Context, req *movie_booking.DeleteSeatCategoryByNameRequest) (*movie_booking.DeleteSeatCategoryByNameResponse, error) {
	if err := h.svc.DeleteSeatCategoryByName(ctx, req.Name); err != nil {
		return &movie_booking.DeleteSeatCategoryByNameResponse{
			Message: "failed to delete seat category",
		}, err
	}
	return &movie_booking.DeleteSeatCategoryByNameResponse{
		Message: "successfully deleted seat category",
	}, nil
}

func (h *GrpcHandler) GetSeatCategoryByID(ctx context.Context, req *movie_booking.GetSeatCategoryByIDRequest) (*movie_booking.GetSeatCategoryByIDResponse, error) {
	seatCategory, err := h.svc.GetSeatCategoryByID(ctx, int(req.SeatCategoryId))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetSeatCategoryByIDResponse{
		SeatCategory: &movie_booking.SeatCategory{
			Id:               int32(seatCategory.ID),
			SeatCategoryName: seatCategory.SeatCategoryName,
		},
	}, nil
}

func (h *GrpcHandler) GetSeatCategoryByName(ctx context.Context, req *movie_booking.GetSeatCategoryByNameRequest) (*movie_booking.GetSeatCategoryByNameResponse, error) {
	seatCategory, err := h.svc.GetSeatCategoryByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetSeatCategoryByNameResponse{
		SeatCategory: &movie_booking.SeatCategory{
			Id:               int32(seatCategory.ID),
			SeatCategoryName: seatCategory.SeatCategoryName,
		},
	}, nil
}

func (h *GrpcHandler) UpdateSeatCategory(ctx context.Context, req *movie_booking.UpdateSeatCategoryRequest) (*movie_booking.UpdateSeatCategoryResponse, error) {
	err := h.svc.UpdateSeatCategory(ctx, int(req.Id), SeatCategory{
		SeatCategoryName: req.SeatCategory.SeatCategoryName,
	})
	if err != nil {
		return nil, err
	}
	return &movie_booking.UpdateSeatCategoryResponse{
		Message: "successfully updated seat category",
	}, nil
}

func (h *GrpcHandler) ListSeatCategories(ctx context.Context, req *movie_booking.ListSeatCategoriesRequest) (*movie_booking.ListSeatCategoriesResponse, error) {
	response, err := h.svc.ListSeatCategories(ctx)
	if err != nil {
		return nil, err
	}

	var grpcSeatCategories []*movie_booking.SeatCategory
	for _, m := range response {
		grpcSeatCategory := &movie_booking.SeatCategory{
			Id:               int32(m.ID),
			SeatCategoryName: m.SeatCategoryName,
		}
		grpcSeatCategories = append(grpcSeatCategories, grpcSeatCategory)
	}

	return &movie_booking.ListSeatCategoriesResponse{
		SeatCategories: grpcSeatCategories,
	}, nil
}

// Theater handlers
func (h *GrpcHandler) AddTheater(ctx context.Context, req *movie_booking.AddTheaterRequest) (*movie_booking.AddTheaterResponse, error) {
	if err := h.svc.AddTheater(ctx, Theater{
		Name:            req.Name,
		Place:           req.Place,
		City:            req.City,
		District:        req.District,
		State:           req.State,
		OwnerID:         uint(req.OwnerId),
		NumberOfScreens: int(req.NumberOfScreens),
		TheaterTypeID:   int(req.TheaterTypeId),
	}); err != nil {
		return &movie_booking.AddTheaterResponse{}, err
	}
	return &movie_booking.AddTheaterResponse{}, nil
}

func (h *GrpcHandler) DeleteTheaterByID(ctx context.Context, req *movie_booking.DeleteTheaterRequest) (*movie_booking.DeleteTheaterResponse, error) {
	if err := h.svc.DeleteTheaterByID(ctx, int(req.TheaterId)); err != nil {
		return &movie_booking.DeleteTheaterResponse{}, err
	}
	return &movie_booking.DeleteTheaterResponse{}, nil
}

func (h *GrpcHandler) DeleteTheaterByName(ctx context.Context, req *movie_booking.DeleteTheaterByNameRequest) (*movie_booking.DeleteTheaterByNameResponse, error) {
	if err := h.svc.DeleteTheaterByName(ctx, req.Name); err != nil {
		return &movie_booking.DeleteTheaterByNameResponse{}, err
	}
	return &movie_booking.DeleteTheaterByNameResponse{}, nil
}

func (h *GrpcHandler) GetTheaterByID(ctx context.Context, req *movie_booking.GetTheaterByIDRequest) (*movie_booking.GetTheaterByIDResponse, error) {
	theater, err := h.svc.GetTheaterByID(ctx, int(req.TheaterId))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetTheaterByIDResponse{
		Theater: &movie_booking.Theater{
			TheaterId:       int32(theater.ID),
			Name:            theater.Name,
			Place:           theater.Place,
			City:            theater.City,
			District:        theater.District,
			State:           theater.State,
			OwnerId:         uint32(theater.OwnerID),
			NumberOfScreens: int32(theater.NumberOfScreens),
			TheaterTypeId:   int32(theater.TheaterTypeID),
		},
	}, nil
}
func (h *GrpcHandler) GetTheaterByName(ctx context.Context, req *movie_booking.GetTheaterByNameRequest) (*movie_booking.GetTheaterByNameResponse, error) {
	theaters, err := h.svc.GetTheaterByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	var theaterResponses []*movie_booking.Theater
	for _, theater := range theaters {
		theaterResponse := &movie_booking.Theater{
			TheaterId:       int32(theater.ID),
			Name:            theater.Name,
			Place:           theater.Place,
			City:            theater.City,
			District:        theater.District,
			State:           theater.State,
			OwnerId:         uint32(theater.OwnerID),
			NumberOfScreens: int32(theater.NumberOfScreens),
			TheaterTypeId:   int32(theater.TheaterTypeID),
		}
		theaterResponses = append(theaterResponses, theaterResponse)
	}
	return &movie_booking.GetTheaterByNameResponse{
		Theater: theaterResponses,
	}, nil
}

func (h *GrpcHandler) UpdateTheater(ctx context.Context, req *movie_booking.UpdateTheaterRequest) (*movie_booking.UpdateTheaterResponse, error) {

	err := h.svc.UpdateTheater(ctx, uint(req.TheaterId), TheaterUpdateInput{
		Name:            req.Name,
		Place:           req.Place,
		City:            req.City,
		District:        req.District,
		State:           req.State,
		OwnerID:         uint(req.OwnerId),
		NumberOfScreens: int(req.NumberOfScreens),
		TheaterTypeID:   int(req.TheaterTypeId),
	})
	if err != nil {
		return nil, err
	}
	return &movie_booking.UpdateTheaterResponse{}, nil
}

func (h *GrpcHandler) ListTheaters(ctx context.Context, req *movie_booking.ListTheatersRequest) (*movie_booking.ListTheatersResponse, error) {
	response, err := h.svc.ListTheaters(ctx)
	if err != nil {
		return nil, err
	}

	var grpcTheaters []*movie_booking.Theater
	for _, m := range response {
		grpcTheater := &movie_booking.Theater{
			TheaterId:       int32(m.ID),
			Name:            m.Name,
			Place:           m.Place,
			City:            m.City,
			District:        m.District,
			State:           m.State,
			OwnerId:         uint32(m.OwnerID),
			NumberOfScreens: int32(m.NumberOfScreens),
			TheaterTypeId:   int32(m.TheaterTypeID),
		}
		grpcTheaters = append(grpcTheaters, grpcTheater)
	}

	return &movie_booking.ListTheatersResponse{
		Theaters: grpcTheaters,
	}, nil
}

// Theater - Screen
func (h *GrpcHandler) AddTheaterScreen(ctx context.Context, req *movie_booking.AddTheaterScreenRequest) (*movie_booking.AddTheaterScreenResponse, error) {
	if err := h.svc.AddTheaterScreen(ctx, TheaterScreen{
		TheaterID:    int(req.TheaterScreen.TheaterID),
		ScreenNumber: int(req.TheaterScreen.ScreenNumber),
		SeatCapacity: int(req.TheaterScreen.SeatCapacity),
		ScreenTypeID: int(req.TheaterScreen.ScreenTypeID),
	}); err != nil {
		return &movie_booking.AddTheaterScreenResponse{}, err
	}
	return &movie_booking.AddTheaterScreenResponse{}, nil
}

func (h *GrpcHandler) DeleteTheaterScreenByID(ctx context.Context, req *movie_booking.DeleteTheaterScreenRequest) (*movie_booking.DeleteTheaterScreenResponse, error) {
	if err := h.svc.DeleteTheaterScreenByID(ctx, int(req.TheaterScreenId)); err != nil {
		return &movie_booking.DeleteTheaterScreenResponse{}, err
	}
	return &movie_booking.DeleteTheaterScreenResponse{}, nil
}

func (h *GrpcHandler) DeleteTheaterScreenByNumber(ctx context.Context, req *movie_booking.DeleteTheaterScreenByNumberRequest) (*movie_booking.DeleteTheaterScreenByNumberResponse, error) {
	if err := h.svc.DeleteTheaterScreenByNumber(ctx, int(req.TheaterID), int(req.ScreenNumber)); err != nil {
		return &movie_booking.DeleteTheaterScreenByNumberResponse{}, err
	}
	return &movie_booking.DeleteTheaterScreenByNumberResponse{}, nil
}

func (h *GrpcHandler) GetTheaterScreenByID(ctx context.Context, req *movie_booking.GetTheaterScreenByIDRequest) (*movie_booking.GetTheaterScreenByIDResponse, error) {
	theaterScreen, err := h.svc.GetTheaterScreenByID(ctx, int(req.TheaterScreenId))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetTheaterScreenByIDResponse{
		TheaterScreen: &movie_booking.TheaterScreen{
			ID:           uint32(theaterScreen.ID),
			TheaterID:    int32(theaterScreen.TheaterID),
			ScreenNumber: int32(theaterScreen.ScreenNumber),
			SeatCapacity: int32(theaterScreen.SeatCapacity),
			ScreenTypeID: int32(theaterScreen.ScreenTypeID),
		},
	}, nil
}

func (h *GrpcHandler) GetTheaterScreenByNumber(ctx context.Context, req *movie_booking.GetTheaterScreenByNumberRequest) (*movie_booking.GetTheaterScreenByNumberResponse, error) {
	theaterScreen, err := h.svc.GetTheaterScreenByNumber(ctx, int(req.TheaterID), int(req.ScreenNumber))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetTheaterScreenByNumberResponse{
		TheaterScreen: &movie_booking.TheaterScreen{
			ID:           uint32(theaterScreen.ID),
			TheaterID:    int32(theaterScreen.TheaterID),
			ScreenNumber: int32(theaterScreen.ScreenNumber),
			SeatCapacity: int32(theaterScreen.SeatCapacity),
			ScreenTypeID: int32(theaterScreen.ScreenTypeID),
		},
	}, nil
}

func (h *GrpcHandler) UpdateTheaterScreen(ctx context.Context, req *movie_booking.UpdateTheaterScreenRequest) (*movie_booking.UpdateTheaterScreenResponse, error) {
	err := h.svc.UpdateTheaterScreen(ctx, int(req.TheaterScreen.ID), TheaterScreen{
		TheaterID:    int(req.TheaterScreen.TheaterID),
		ScreenNumber: int(req.TheaterScreen.ScreenNumber),
		SeatCapacity: int(req.TheaterScreen.SeatCapacity),
		ScreenTypeID: int(req.TheaterScreen.ScreenTypeID),
	})
	if err != nil {
		return nil, err
	}
	return &movie_booking.UpdateTheaterScreenResponse{}, nil
}

func (h *GrpcHandler) ListTheaterScreens(ctx context.Context, req *movie_booking.ListTheaterScreensRequest) (*movie_booking.ListTheaterScreensResponse, error) {
	response, err := h.svc.ListTheaterScreens(ctx, int(req.TheaterID))
	if err != nil {
		return nil, err
	}

	var grpcTheaterScreens []*movie_booking.TheaterScreen
	for _, m := range response {
		grpcTheaterScreen := &movie_booking.TheaterScreen{
			ID:           uint32(m.ID),
			TheaterID:    int32(m.TheaterID),
			ScreenNumber: int32(m.ScreenNumber),
			SeatCapacity: int32(m.SeatCapacity),
			ScreenTypeID: int32(m.ScreenTypeID),
		}
		grpcTheaterScreens = append(grpcTheaterScreens, grpcTheaterScreen)
	}

	return &movie_booking.ListTheaterScreensResponse{
		TheaterScreens: grpcTheaterScreens,
	}, nil
}

// Showtime
func (h *GrpcHandler) AddShowtime(ctx context.Context, req *movie_booking.AddShowtimeRequest) (*movie_booking.AddShowtimeResponse, error) {
	if err := h.svc.AddShowtime(ctx, Showtime{
		MovieID:  int(req.Showtime.MovieId),
		ScreenID: int(req.Showtime.ScreenId),
		ShowDate: req.Showtime.ShowDate.AsTime(),
		ShowTime: req.Showtime.ShowTime.AsTime(),
	}); err != nil {
		return &movie_booking.AddShowtimeResponse{}, err
	}
	return &movie_booking.AddShowtimeResponse{}, nil
}

func (h *GrpcHandler) DeleteShowtimeByID(ctx context.Context, req *movie_booking.DeleteShowtimeRequest) (*movie_booking.DeleteShowtimeResponse, error) {
	if err := h.svc.DeleteShowtimeByID(ctx, int(req.ShowtimeId)); err != nil {
		return &movie_booking.DeleteShowtimeResponse{}, err
	}
	return &movie_booking.DeleteShowtimeResponse{}, nil
}

func (h *GrpcHandler) DeleteShowtimeByDetails(ctx context.Context, req *movie_booking.DeleteShowtimeByDetailsRequest) (*movie_booking.DeleteShowtimeByDetailsResponse, error) {
	if err := h.svc.DeleteShowtimeByDetails(ctx, int(req.MovieId), int(req.ScreenId), req.ShowDate.AsTime(), req.ShowTime.AsTime()); err != nil {
		return &movie_booking.DeleteShowtimeByDetailsResponse{}, err
	}
	return &movie_booking.DeleteShowtimeByDetailsResponse{}, nil
}

func (h *GrpcHandler) GetShowtimeByID(ctx context.Context, req *movie_booking.GetShowtimeByIDRequest) (*movie_booking.GetShowtimeByIDResponse, error) {
	showtime, err := h.svc.GetShowtimeByID(ctx, int(req.ShowtimeId))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetShowtimeByIDResponse{
		Showtime: &movie_booking.Showtime{
			Id:       uint32(showtime.ID),
			MovieId:  int32(showtime.MovieID),
			ScreenId: int32(showtime.ScreenID),
			ShowDate: timestamppb.New(showtime.ShowDate),
			ShowTime: timestamppb.New(showtime.ShowTime),
		},
	}, nil
}

func (h *GrpcHandler) GetShowtimeByDetails(ctx context.Context, req *movie_booking.GetShowtimeByDetailsRequest) (*movie_booking.GetShowtimeByDetailsResponse, error) {
	showtime, err := h.svc.GetShowtimeByDetails(ctx, int(req.MovieId), int(req.ScreenId), req.ShowDate.AsTime(), req.ShowTime.AsTime())
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetShowtimeByDetailsResponse{
		Showtime: &movie_booking.Showtime{
			Id:       uint32(showtime.ID),
			MovieId:  int32(showtime.MovieID),
			ScreenId: int32(showtime.ScreenID),
			ShowDate: timestamppb.New(showtime.ShowDate),
			ShowTime: timestamppb.New(showtime.ShowTime),
		},
	}, nil
}

func (h *GrpcHandler) UpdateShowtime(ctx context.Context, req *movie_booking.UpdateShowtimeRequest) (*movie_booking.UpdateShowtimeResponse, error) {
	err := h.svc.UpdateShowtime(ctx, int(req.Showtime.Id), Showtime{
		MovieID:  int(req.Showtime.MovieId),
		ScreenID: int(req.Showtime.ScreenId),
		ShowDate: req.Showtime.ShowDate.AsTime(),
		ShowTime: req.Showtime.ShowTime.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &movie_booking.UpdateShowtimeResponse{}, nil
}

func (h *GrpcHandler) ListShowtimes(ctx context.Context, req *movie_booking.ListShowtimesRequest) (*movie_booking.ListShowtimesResponse, error) {
	response, err := h.svc.ListShowtimes(ctx, int(req.MovieId))
	if err != nil {
		return nil, err
	}

	var grpcShowtimes []*movie_booking.Showtime
	for _, m := range response {
		grpcShowtime := &movie_booking.Showtime{
			Id:       uint32(m.ID),
			MovieId:  int32(m.MovieID),
			ScreenId: int32(m.ScreenID),
			ShowDate: timestamppb.New(m.ShowDate),
			ShowTime: timestamppb.New(m.ShowTime),
		}
		grpcShowtimes = append(grpcShowtimes, grpcShowtime)
	}

	return &movie_booking.ListShowtimesResponse{
		Showtimes: grpcShowtimes,
	}, nil
}
