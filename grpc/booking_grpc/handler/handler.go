package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	_ "mock_project/ent/booking"
	"mock_project/grpc/booking_grpc/repository"
	"mock_project/pb"
)

type BookingServiceImp interface {
	GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*pb.Booking, error)
	CreateBooking(ctx context.Context, modelBooking *pb.CreateBookingRequest) (*pb.Booking, error)
	UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (*pb.UpdateBookingStatusResponse, error)
	DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (bool, error)
	CustomerBookingHistory(ctx context.Context, model *pb.CustomerBookingHistoryRequest) (*pb.Bookings, error)
}

type BookingService struct {
	BookingRepository repository.BookingRepository
	pb.UnimplementedBookingServiceRPCServer
}

func NewBookingrHandler(BookingRepository repository.BookingRepository) (*BookingService, error) {
	return &BookingService{
		BookingRepository: BookingRepository,
	}, nil
}

func (s *BookingService) GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*pb.Booking, error) {
	res, err := s.BookingRepository.GetBookingByID(ctx, model)

	if err != nil {
		log.Println("Can not find the desired boooking", err)
		return nil, err
	}

	return &pb.Booking{
		Id:                  res.ID.String(),
		TotalEconomyTicket:  int32(res.EconomyTickets),
		TotalBusinessTicket: int32(res.BusinessTickets),
		Status:              pb.Booking_BookingStatus(pb.Booking_BookingStatus_value[res.Status.String()]),
		Type:                pb.Booking_Type(pb.Booking_Type_value[res.Type.String()]),
		CreatedAt:           timestamppb.New(res.CreatedAt),
		UpdatedAt:           timestamppb.New(res.UpdatedAt),
	}, nil
}

func (s *BookingService) CustomerBookingHistory(ctx context.Context, model *pb.CustomerBookingHistoryRequest) (*pb.Bookings, error) {
	res, err := s.BookingRepository.CustomerBookingHistory(ctx, model)

	log.Println("HANDLERHANDLER",res)


	if err != nil {
		log.Println("Can not find the desired boooking", err)
		return nil, err
	}

	var pRes []*pb.Booking

	for _, e := range res {
		temp := &pb.Booking{
				Id: e.ID.String(),
				TotalEconomyTicket: int32(e.EconomyTickets),
				TotalBusinessTicket: int32(e.BusinessTickets),
				Status: pb.Booking_BookingStatus(pb.Booking_BookingStatus_value[e.Status.String()]),
				Type: pb.Booking_Type(pb.Booking_Type_value[e.Type.String()]),
				CreatedAt: timestamppb.New(e.CreatedAt),
				UpdatedAt: timestamppb.New(e.UpdatedAt),
		}

		pRes = append(pRes, temp)
	}

	log.Println("HANDLERHANDLER",pRes)

	return &pb.Bookings{Bookings: pRes}, nil
}

func (s *BookingService) CreateBooking(ctx context.Context, modelBooking *pb.CreateBookingRequest) (*pb.Booking, error) {
	res, err := s.BookingRepository.CreateBooking(ctx, modelBooking)

	if err != nil {
		log.Println("Create boooking failed", err)
		return nil, err
	}

	return &pb.Booking{
		Id:                  res.ID.String(),
		TotalEconomyTicket:  int32(res.EconomyTickets),
		TotalBusinessTicket: int32(res.BusinessTickets),
		Status:              pb.Booking_BookingStatus(pb.Booking_BookingStatus_value[res.Status.String()]),
		Type:                pb.Booking_Type(pb.Booking_Type_value[res.Type.String()]),
		CreatedAt:           timestamppb.New(res.CreatedAt),
		UpdatedAt:           timestamppb.New(res.UpdatedAt),
	}, nil
}

func (s *BookingService) UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (*pb.UpdateBookingStatusResponse, error) {
	

	res, err := s.BookingRepository.UpdateBookingStatus(ctx, model)

	if err != nil {
		return &pb.UpdateBookingStatusResponse{
			IsUpdated: res,
		}, err
	}

	return &pb.UpdateBookingStatusResponse{
		IsUpdated: res,
	}, nil
}

func (s *BookingService) DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (*pb.DeleteBookingResponse, error) {
	res, err := s.BookingRepository.DeleteBooking(ctx, model)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeleteBookingResponse{
			IsDeleted: res,
		}, err
	}
	return &pb.DeleteBookingResponse{
		IsDeleted: res,
	}, nil
}
