package resolver

import (
	"context"
	"fmt"
	"mock_project/ent"
	"mock_project/internal/util"
	"mock_project/pb"

	generated "mock_project/graphql"

	"github.com/99designs/gqlgen/graphql"
	ut "github.com/go-playground/universal-translator"

	"mock_project/grpc/services"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client               *ent.Client
	validator            *validator.Validate
	validationTranslator ut.Translator
	logger               *zap.Logger
	customerService      services.CustomerServiceClient
	accountService       services.AccountServiceClient
	flightService		 services.FlightServiceClient
	passengerService	 services.PassengerServiceClient
	ticketService	 	 services.TicketServiceClient
	bookingService		 services.BookingServiceClient
}

func NewSchema(client *ent.Client, validator *validator.Validate, validationTranslator ut.Translator, logger *zap.Logger, grpcCustomer pb.CustomerServiceRPCClient, grpcAccount pb.AccountServiceRPCClient, grpcFlight pb.FlightServiceRPCClient, grpcPassenger pb.PassengerServiceRPCClient, grpcTicket pb.TicketServiceRPCClient, grpcBooking pb.BookingServiceRPCClient) graphql.ExecutableSchema {
	customerService := services.NewCustomerHandler(grpcCustomer)
	accountService := services.NewAccountHandler(grpcAccount, grpcCustomer)
	flightService := services.NewFlightHandler(grpcFlight, grpcTicket, grpcBooking)
	passengerService := services.NewPassengerHandler(grpcPassenger)
	ticketService := services.NewTicketHandler(grpcTicket)
	bookingService := services.NewBookingHandler(grpcBooking, grpcFlight, grpcCustomer, grpcPassenger, grpcTicket)
	
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client: client, validator: validator, validationTranslator: validationTranslator, logger: logger, customerService: customerService, accountService: accountService, flightService: flightService, passengerService: passengerService, ticketService: ticketService, bookingService: bookingService},
	})
}

func (r *Resolver) ValidationResolver() func(ctx context.Context, obj interface{}, next graphql.Resolver, constrains string) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, constrains string) (interface{}, error) {
		val, err := next(ctx)
		if err != nil {
			r.logger.Error("Getting error when extract values from context", zap.Error(err))
			return nil, util.WrapGQLInternalError(ctx)
		}

		fieldName := *graphql.GetPathContext(ctx).Field

		err = r.validator.Var(val, constrains)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return nil, fmt.Errorf("%s:%s", fieldName, validationErrors[0].Translate(r.validationTranslator))
		}

		return val, nil
	}
}
