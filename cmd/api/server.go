package api

import (
	"context"
	"log"
	"mock_project/config"
	"mock_project/ent"
	"mock_project/middleware"
	"mock_project/middleware/auth"
	"mock_project/pb"
	"mock_project/resolver"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewServerCmd(configs *config.Configurations, logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with graphql",
		Run: func(cmd *cobra.Command, args []string) {
			// Connect to postgresql database
			db, err := ent.Open("postgres", configs.Postgres.ConnectionString)
			if err != nil {
				logger.Error("Getting error connect to postgresql database", zap.Error(err))
				os.Exit(1)
			}
			defer db.Close()

			// Run the automation migration tool
			if err := db.Schema.Create(context.Background()); err != nil {
				logger.Error("Failed to creating db schema from the automation migration tool", zap.Error(err))
				os.Exit(1)
			}
			
			//Conenct to Customer server
			conn1, err := grpc.Dial(":2223", grpc.WithTransportCredentials(insecure.NewCredentials()))
    		if err != nil {
        		log.Fatalf("Failed connecting to Customer server: %s", err)
    		}
    		defer conn1.Close()

			//Connect to Account server
			conn2, err := grpc.Dial(":2224", grpc.WithTransportCredentials(insecure.NewCredentials()))
    		if err != nil {
        		log.Fatalf("Failed connecting to Account server: %s", err)
    		}
    		defer conn2.Close()

			//Connect to Flight server
			conn3, err := grpc.Dial(":2225", grpc.WithTransportCredentials(insecure.NewCredentials()))
    		if err != nil {
        		log.Fatalf("Failed connecting to Flight server: %s", err)
    		}
    		defer conn3.Close()

			//Connect to Passenger server
			conn4, err := grpc.Dial(":2226", grpc.WithTransportCredentials(insecure.NewCredentials()))
    		if err != nil {
        		log.Fatalf("Failed connecting to Passenger server: %s", err)
    		}
    		defer conn4.Close()

			//Connect to Ticket server
			conn5, err := grpc.Dial(":2227", grpc.WithTransportCredentials(insecure.NewCredentials()))
    		if err != nil {
        		log.Fatalf("Failed connecting to Passenger server: %s", err)
    		}
    		defer conn5.Close()

			//Connect to Booking server
			conn6, err := grpc.Dial(":2228", grpc.WithTransportCredentials(insecure.NewCredentials()))
    		if err != nil {
        		log.Fatalf("Failed connecting to Passenger server: %s", err)
    		}
    		defer conn6.Close()

			// Create validator
			validator := validator.New()
			en := en.New()
			uni := ut.New(en, en)
			validationTranslator, _ := uni.GetTranslator("en")

			customerService := pb.NewCustomerServiceRPCClient(conn1)
			accountService := pb.NewAccountServiceRPCClient(conn2)
			flightService := pb.NewFlightServiceRPCClient(conn3)
			passengerService := pb.NewPassengerServiceRPCClient(conn4)
			ticketService := pb.NewTicketServiceRPCClient(conn5)
			bookingService := pb.NewBookingServiceRPCClient(conn6)


			// Register default translation for validator
			err = en_translations.RegisterDefaultTranslations(validator, validationTranslator)
			if err != nil {
				logger.Error("Getting error from register default translation", zap.Error(err))
				os.Exit(1)
			}

			// GraphQL schema resolver handler.
			resolverHandler := handler.NewDefaultServer(resolver.NewSchema(db, validator, validationTranslator, logger, customerService, accountService, flightService, passengerService, ticketService, bookingService))
			// Handler for GraphQL Playground
			playgroundHandler := playground.Handler("GraphQL Playground", "/graphql")

			// Handle a method not allowed.
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.HandleMethodNotAllowed = true

			// Use middlewares
			r.Use(
				auth.Auth(accountService),
				ginzap.Ginzap(logger, time.RFC3339, true),
				ginzap.RecoveryWithZap(logger, true),
				middleware.CorsMiddleware(),
				middleware.RequestCtxMiddleware(),
			)


			// Create a new GraphQL
			r.POST("/graphql", func(c *gin.Context) {
				resolverHandler.ServeHTTP(c.Writer, c.Request)
			})

			r.OPTIONS("/graphql", func(c *gin.Context) {
				c.Status(200)
			})

			// Enable playground for query/testing
			r.GET("/", func(c *gin.Context) {
				playgroundHandler.ServeHTTP(c.Writer, c.Request)
			})


			// Listen on port 8000
			logger.Info("Listening on port: 8000")
			if err := r.Run(":8000"); err != nil {
				logger.Error("Get error from run server", zap.Error(err))
			}
		},
	}
}
