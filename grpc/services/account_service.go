package services

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/account"
	"mock_project/ent/customer"
	"mock_project/grpc/account_grpc/response"
	"mock_project/internal/util"
	"mock_project/pb"
	"mock_project/middleware/auth"
	"github.com/google/uuid"
)

type AccountServiceClient interface {
	GetAccountByID(ctx context.Context, id uuid.UUID) (*ent.Account, error)
	CreateAccount(ctx context.Context, input ent.NewAccountInput) (*ent.Account, error)
	AccountResetPassword(ctx context.Context, input ent.AccountResetPasswordInput) (string, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateAccount(ctx context.Context, input ent.UpdateAccountByIDInput) (*ent.Account, error)
	UpdateAccountStatus(ctx context.Context, input ent.UpdateAccountStatusRequest) (*ent.Account, error)
	Login(ctx context.Context, input ent.AccountLogin) (response.LoginResponse, error)
}

type AccountHandler struct {
	AccountClient  pb.AccountServiceRPCClient
	CustomerClient pb.CustomerServiceRPCClient
}

func NewAccountHandler(AccountClient pb.AccountServiceRPCClient, CustomerClient pb.CustomerServiceRPCClient) AccountServiceClient {
	return &AccountHandler{
		AccountClient:  AccountClient,
		CustomerClient: CustomerClient,
	}
}

func (h *AccountHandler) CreateAccount(ctx context.Context, input ent.NewAccountInput) (*ent.Account, error) {

	hashPassword, _ := util.HashPassword(input.Password)

	customer, err := h.CustomerClient.GetCustomerByCitizenId(ctx, &pb.GetCustomerByCitizenIdRequest{CitizenId: input.Customer.CitizenID})

	if customer == nil {
		newCustomer, err := h.CustomerClient.CreateCustomer(ctx, &pb.CreateCustomerRequest{
			Name: input.Customer.Name, 
			Address: input.Customer.Address, 
			CitizenId: input.Customer.CitizenID, 
			Phone: input.Customer.Phone, 
			Gender: pb.CreateCustomerRequest_Gender(pb.CreateCustomerRequest_Gender_value[string(input.Customer.Gender)])})
		if err != nil {
			log.Fatalf("Created failed: %s", err)
		}

		res, err := h.AccountClient.CreateAccount(ctx, &pb.CreateAccountRequest{Email: input.Email, Password: hashPassword, Role: pb.CreateAccountRequest_Role(pb.CreateAccountRequest_Role_value[string(input.Role)]), AccOwner: newCustomer.Id})

		if err != nil {
			return nil, err
		}

		return &ent.Account{
			ID:        uuid.MustParse(res.Id),
			Email:     res.Email,
			Role:      account.Role(res.Role.String()),
			Status: account.Status(res.Status.String()),
			CreatedAt: res.CreatedAt.AsTime(),
			UpdatedAt: res.UpdatedAt.AsTime(),
		}, nil
	}

	log.Println("Customer is already in use:", err)
	res, err := h.AccountClient.CreateAccount(ctx, &pb.CreateAccountRequest{Email: input.Email, Password: hashPassword, Role: pb.CreateAccountRequest_Role(pb.CreateAccountRequest_Role_value[string(input.Role)]), AccOwner: customer.Id})

	log.Println("NEW ACCOUNT SERVICE", res)

	if err != nil {
		return nil, err
	}

	return &ent.Account{
		ID:        uuid.MustParse(res.Id),
		Email:     res.Email,
		Role:      account.Role(res.Role.String()),
		Status: account.Status(res.Status.String()),
		CreatedAt: res.CreatedAt.AsTime(),
		UpdatedAt: res.UpdatedAt.AsTime(),
	}, nil
}

func (h *AccountHandler) GetAccountByID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {


	res, err := h.AccountClient.GetAccountByID(ctx, &pb.GetAccountByIdRequest{Id: id.String()})

	if err != nil {
		log.Println("Error account service",err)
		return nil, err
	}

	return &ent.Account{
		ID:        uuid.MustParse(res.Id),
		Email:     res.Email,
		Role:      account.Role(res.Role.String()),
		Status: account.Status(res.Status.String()),
		Edges:     ent.AccountEdges{AccOwner: &ent.Customer{ID: uuid.MustParse(res.AccOwner.Id), Name: res.AccOwner.Name, CitizenID: res.AccOwner.CitizenId, Phone: res.AccOwner.Phone, Address: res.AccOwner.Address, Gender: customer.Gender(res.AccOwner.Gender.String()), DateOfBirth: res.AccOwner.Dob.AsTime()}},
		CreatedAt: res.CreatedAt.AsTime(),
		UpdatedAt: res.UpdatedAt.AsTime(),
	}, nil
}

func (h *AccountHandler) AccountResetPassword(ctx context.Context, input ent.AccountResetPasswordInput) (string, error) {
	res, err := h.AccountClient.AccountResetPassword(ctx, &pb.AccountResetPasswordRequest{Email: input.Email, CurrentPassword: input.CurrentPassword, NewPassword: input.NewPassword})
	if err != nil {
		return res.IsUpdate, err
	}
	return res.IsUpdate, nil
}

func (h *AccountHandler) UpdateAccount(ctx context.Context, input ent.UpdateAccountByIDInput) (*ent.Account, error) {
	
	res, err := h.AccountClient.UpdateAccount(ctx, &pb.UpdateAccountRequest{Id: input.ID, Role: pb.UpdateAccountRequest_Role(pb.UpdateAccountRequest_Role_value[input.Role.String()])})
	if err != nil {
		return nil, err
	}
	return &ent.Account{
		ID:        uuid.MustParse(res.Id),
		Email:     res.Email,
		Role:      account.Role(res.Role.String()),
		Status:    account.Status(res.Status.String()),
		CreatedAt: res.CreatedAt.AsTime(),
		UpdatedAt: res.UpdatedAt.AsTime(),
	}, nil
}

func (h *AccountHandler) UpdateAccountStatus(ctx context.Context, input ent.UpdateAccountStatusRequest) (*ent.Account, error) {
	
	res, err := h.AccountClient.UpdateAccountStatus(ctx, &pb.UpdateAccountStatusRequest{
		Id: input.ID, 
		Status: pb.UpdateAccountStatusRequest_AccountStatus(pb.UpdateAccountStatusRequest_AccountStatus_value[input.Status.String()])})
	if err != nil {
		return nil, err
	}
	return &ent.Account{
		ID:        uuid.MustParse(res.Id),
		Email:     res.Email,
		Role:      account.Role(res.Role.String()),
		Status:    account.Status(res.Status.String()),
		CreatedAt: res.CreatedAt.AsTime(),
		UpdatedAt: res.UpdatedAt.AsTime(),
	}, nil
}

func (h *AccountHandler) DeleteAccount(ctx context.Context, id uuid.UUID) (bool, error) {
	is_authenticated := auth.ForContext(ctx)

	if is_authenticated == nil {
			return false, util.WrapGQLUnauthenticatedError(ctx)
	}
	res, err := h.AccountClient.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: id.String()})
	
	if err != nil {
		return res.IsDeleted, err
	}

	return res.IsDeleted, nil
}

func (h *AccountHandler) Login(ctx context.Context, input ent.AccountLogin) (response.LoginResponse, error) {
	res, err := h.AccountClient.Login(ctx, &pb.LoginRequest{Email: input.Email, Password: input.Password})
	if err != nil || !res.Status{
		log.Println("Login fail:", err)
		return response.LoginResponse{
			Token:  res.Token,
			Status: res.Status,
		}, err
	}

	log.Println("Login success:")
	return response.LoginResponse{
		Token:  res.Token,
		Status: res.Status,
	}, nil
}


