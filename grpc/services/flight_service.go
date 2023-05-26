package services

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/flight"
	"mock_project/internal/util"
	"mock_project/internal/validation"
	"mock_project/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightServiceClient interface {
	GetFlightByFlightCode(ctx context.Context, fcode string) (*ent.Flight, error)
	CreateFlight(ctx context.Context, input ent.NewFlightInput) (*ent.Flight, error)
	UpdateFlightByFlightCode(ctx context.Context, input ent.UpdateFlightByFlightCodeInput) (*ent.Flight, error)
	UpdateAvailableSeat(ctx context.Context, input ent.UpdateAvailableSeatInput) (*ent.Flight, error)
	UpdateFlightStatus(ctx context.Context, input ent.UpdateFlightStatusInput) (*ent.Flight, error)
	DeleteFlight(ctx context.Context,  id string) (bool, error)
	FindFlightByDateRange(ctx context.Context, input ent.GetFlightByDateRangeInput) ([]*ent.Flight, error)
}

type FlightHandler struct {
	FlightClient pb.FlightServiceRPCClient
	TicketClient pb.TicketServiceRPCClient
	BookingClient pb.BookingServiceRPCClient
}

func NewFlightHandler(FlightClient pb.FlightServiceRPCClient, TicketClient pb.TicketServiceRPCClient, BookingClient pb.BookingServiceRPCClient) FlightServiceClient {
	return &FlightHandler{
		FlightClient: FlightClient,
		TicketClient: TicketClient,
		BookingClient: BookingClient,
	}
}


func (h *FlightHandler) CreateFlight(ctx context.Context, input ent.NewFlightInput) (*ent.Flight, error) {
	inputCapacity := input.FlightInput.Capacity
	inputEconomySeat := input.FlightInput.EconomyAvailableSeat
	inputBusinessSeat := input.FlightInput.BusinessAvailableSeat

	isValidSeat := validation.ValidateAvailableSeat(inputCapacity,inputEconomySeat,inputBusinessSeat)

	if !isValidSeat{
			return nil, util.WrapGQLInvalidInputError(ctx)
	}

	res, err := h.FlightClient.CreateFlight(ctx, &pb.CreateFlightRequest{FlightInput: &pb.FlightInput{
		FlightCode: input.FlightInput.FlightCode,
		From: input.FlightInput.From,
		To:   input.FlightInput.To, 
		DepartureDate: timestamppb.New(input.FlightInput.DepartureDate),
		DepartureTime: timestamppb.New(input.FlightInput.DepartureTime),
		Duration:      int32(input.FlightInput.Duration),
		Capacity:      int32(inputCapacity),
		EcomomyAvailableSeat:  int32(inputEconomySeat),
		BusinessAvailableSeat: int32(inputBusinessSeat),
		Status:        pb.FlightInput_Status(pb.FlightInput_Status_value[input.FlightInput.Status.String()])}})
	if err != nil {
		return nil, err
	}

	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		EconomyAvailableSeat: int(res.EcomomyAvailableSeat),
		BusinessAvailableSeat: int(res.BusinessAvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil
}

func (h *FlightHandler) GetFlightByFlightCode(ctx context.Context, fcode string) (*ent.Flight, error) {
	res, err := h.FlightClient.GetFlightByFlightCode(ctx, &pb.GetFlightByFlightCodeRequest{FlightCode: fcode})
	if err != nil {
		log.Println("Can not find desired flight", err)
		return nil, err
	}

	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		EconomyAvailableSeat: int(res.EcomomyAvailableSeat),
		BusinessAvailableSeat: int(res.BusinessAvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil
}

func (h *FlightHandler) UpdateFlightByFlightCode(ctx context.Context, input ent.UpdateFlightByFlightCodeInput) (*ent.Flight, error) {
	res, err := h.FlightClient.UpdateFlightByFlightCode(ctx, &pb.UpdateFlightByFlightCodeRequest{FlightInput: &pb.FlightInput{
		FlightCode: input.FlightInput.FlightCode,
		From: input.FlightInput.From,
		To:   input.FlightInput.To, 
		DepartureDate: timestamppb.New(input.FlightInput.DepartureDate),
		DepartureTime: timestamppb.New(input.FlightInput.DepartureTime),
		Duration:      int32(input.FlightInput.Duration),
		Capacity:      int32(input.FlightInput.Capacity),
		EcomomyAvailableSeat: int32(input.FlightInput.EconomyAvailableSeat),
		BusinessAvailableSeat: int32(input.FlightInput.BusinessAvailableSeat),
		Status:        pb.FlightInput_Status(pb.FlightInput_Status_value[input.FlightInput.Status.String()])}})
	if err != nil {
		log.Println("Update flight failed", err)
		return nil, err
	}
	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		EconomyAvailableSeat: int(res.EcomomyAvailableSeat),
		BusinessAvailableSeat: int(res.BusinessAvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil
}

func (h *FlightHandler) UpdateAvailableSeat(ctx context.Context, input ent.UpdateAvailableSeatInput) (*ent.Flight, error) {
	res, err := h.FlightClient.UpdateAvailableSeat(ctx, &pb.UpdateFlightAvailableSeatRequest{
		FlightCode: input.FlightCode,
		EcomomyAvailableSeat: int32(input.EconomyAvailableSeat),
		BusinessAvailableSeat: int32(input.BusinessAvailableSeat),
	})
	if err != nil {
		log.Println("Update flight failed", err)
		return nil, err
	}
	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		EconomyAvailableSeat: int(res.EcomomyAvailableSeat),
		BusinessAvailableSeat: int(res.BusinessAvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil
}

func (h *FlightHandler) UpdateFlightStatus(ctx context.Context, input ent.UpdateFlightStatusInput) (*ent.Flight, error) {
	
	res, err := h.FlightClient.UpdateFlightStatus(ctx, &pb.UpdateFlightStatusRequest{
		Id: input.ID,
		Status: pb.UpdateFlightStatusRequest_Status(pb.UpdateFlightStatusRequest_Status_value[input.Status.String()]),
	})

	if err != nil {
		return nil, util.WrapGQLInvalidActionError(ctx)
	}

	if input.Status.String() == "Canceled"{
		_, err1 := h.TicketClient.UpdateTicketStatusByFlightId(ctx, &pb.UpdateTicketStatusByFlightIdRequest{
			Id: input.ID,
			Status: pb.UpdateTicketStatusByFlightIdRequest_Canceled,
		})

		if err1 != nil {
			log.Println("ERROR:",err1)
			return nil, util.WrapGQLInvalidActionError(ctx)
		}
		
		_, err2 := h.BookingClient.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{
			Id: input.ID,
			Status: pb.UpdateBookingStatusRequest_Canceled,
		}) 
		
		if err2 != nil {
			log.Println("ERROR:",err2)
			return nil, util.WrapGQLInvalidActionError(ctx)
		}
		
	} 

	return &ent.Flight{
		ID:            uuid.MustParse(res.Id),
		FlightCode:    res.FlightCode,
		From:          res.From,
		To:            res.To,
		DepartureDate: res.DepartureDate.AsTime(),
		DepartureTime: res.DepartureTime.AsTime(),
		Duration:      int(res.Duration),
		Capacity:      int(res.Capacity),
		EconomyAvailableSeat: int(res.EcomomyAvailableSeat),
		BusinessAvailableSeat: int(res.BusinessAvailableSeat),
		Status:        flight.Status(res.Status.String()),
	}, nil
}

func (h *FlightHandler) DeleteFlight(ctx context.Context, id string) (bool, error) {
	res, err := h.FlightClient.DeleteFlight(ctx, &pb.DeleteFlightRequest{Id: id})
	if err != nil {
		return res.IsDeleted, err
	}
	return res.IsDeleted, nil
}

func (h *FlightHandler) FindFlightByDateRange(ctx context.Context, input ent.GetFlightByDateRangeInput) ([]*ent.Flight, error) {	
	log.Println("StartDate", timestamppb.New(input.StartDate))
	log.Println("StartDate", input.StartDate)
	log.Println("EndDate", timestamppb.New(input.EndDate))
	log.Println("StartDate", input.EndDate)
	
	pRes, err := h.FlightClient.FindFlightByDateRange(ctx, &pb.FindFlightByDateRangeRequest{
		FirstDate:  timestamppb.New(input.StartDate),
		SecondDate: timestamppb.New(input.EndDate),
	})
	
	if err != nil {
		log.Println("ERRROR", err)
		return nil, err
	}

	var res []*ent.Flight

	for _, e := range(pRes.Flights){
		temp := &ent.Flight{
			ID:            uuid.MustParse(e.Id),
			FlightCode:    e.FlightCode,
			From:          e.From,
			To:            e.To,
			DepartureDate: e.DepartureDate.AsTime(),
			DepartureTime: e.DepartureTime.AsTime(),
			Duration:      int(e.Duration),
			Capacity:      int(e.Capacity),
			EconomyAvailableSeat: int(e.EcomomyAvailableSeat),
			BusinessAvailableSeat: int(e.BusinessAvailableSeat),
			Status:        flight.Status(e.Status.String()),
		}
		res = append(res, temp)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
