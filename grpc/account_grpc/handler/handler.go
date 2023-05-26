package handler

import (
	"context"
	"log"
	_ "mock_project/ent/account"
	"mock_project/grpc/account_grpc/repository"
	"mock_project/internal/jwt"
	"mock_project/internal/util"
	"mock_project/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServiceImp interface {
	Login(ctx context.Context, model *pb.LoginRequest) (*pb.LoginResponse, error)
	GetAccountByID(ctx context.Context, model *pb.GetAccountByIdRequest) (*pb.Account, error)
	GetAccountByEmail(ctx context.Context, model *pb.GetAccountByEmailRequest) (*pb.Account, error)
	CreateAccount(ctx context.Context, model *pb.CreateAccountRequest) (*pb.Account, error)
	UpdateAccount(ctx context.Context, model *pb.UpdateAccountRequest) (*pb.Account, error)
	UpdateAccountStatus(ctx context.Context, model *pb.UpdateAccountStatusRequest) (*pb.Account, error)
	AccountResetPassword(ctx context.Context, model *pb.AccountResetPasswordRequest) (*pb.AccountResetPasswordResponse, error)
	DeleteAccount(ctx context.Context, model *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error)
}

type AccountService struct {
	AccountRepository repository.AccountRepository
	pb.UnimplementedAccountServiceRPCServer
}

func NewAccountrHandler(AccountRepository repository.AccountRepository) (*AccountService, error) {
	return &AccountService{
		AccountRepository: AccountRepository,
	}, nil
}

//Get Account information by Email
func (us *AccountService) GetAccountByID(ctx context.Context, model *pb.GetAccountByIdRequest) (*pb.Account, error){
	
	acc, customer, err := us.AccountRepository.GetAccountByID(ctx, model)

	if err != nil {
		log.Println("Can not fine user by ID: ", err)
		return nil, err
	}

	return &pb.Account{
		Id:   acc.ID.String(),
		Email: acc.Email,
		Role: pb.Account_Role(pb.CreateAccountRequest_Role_value[acc.Role.String()]),
		Status: pb.Account_AccountStatus(pb.Account_AccountStatus_value[acc.Status.String()]),
		AccOwner: &pb.Customer{Id: customer.ID.String(), Name: customer.Name, CitizenId:customer.CitizenID, Phone: customer.Phone, Address: customer.Address, Gender: pb.Customer_Gender(pb.Customer_Gender_value[customer.Gender.String()]), Dob: timestamppb.New(customer.DateOfBirth)},
		CreatedAt: timestamppb.New(acc.CreatedAt),
		UpdatedAt: timestamppb.New(acc.UpdatedAt),
	}, nil
}


//Get Account information by Email
func (us *AccountService) GetAccountByEmail(ctx context.Context, model *pb.GetAccountByEmailRequest) (*pb.Account, error){
	res, err := us.AccountRepository.GetAccountByEmail(ctx, model)

	if err != nil {
		log.Println("Can not fine user by email: ", err)
		return nil, err
	}

	// ownerId, _ := res.AccOwner(ctx)
	return &pb.Account{
		Id:   res.ID.String(),
		Email: res.Email,
		Role: pb.Account_Role(pb.CreateAccountRequest_Role_value[res.Role.String()]),
		Status: pb.Account_AccountStatus(pb.Account_AccountStatus_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}


func (us *AccountService) CreateAccount(ctx context.Context, model *pb.CreateAccountRequest) (*pb.Account, error){
	res, err := us.AccountRepository.CreateAccount(ctx, model)

	log.Println("NEW ACCOUNT HANDLER", res)

	if err != nil {
		log.Println("Created failed: ", err)
		return nil, err
	}

	// ownerId, _ := res.AccOwner(ctx)

	return &pb.Account{
		Id:   res.ID.String(),
		Email: res.Email,
		Role: pb.Account_Role(pb.CreateAccountRequest_Role_value[res.Role.String()]),
		Status: pb.Account_AccountStatus(pb.Account_AccountStatus_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}


// Delete account service
func (us *AccountService) DeleteAccount(ctx context.Context, model *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error){
	res, err := us.AccountRepository.DeleteAccount(ctx, model)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeleteAccountResponse{
			IsDeleted: res,
		},err
	}
	return &pb.DeleteAccountResponse{
		IsDeleted: res,
	},nil
}


// Login service
func(us *AccountService) Login(ctx context.Context, model *pb.LoginRequest) (*pb.LoginResponse, error){
	
	myUser, err := us.AccountRepository.GetAccountByEmail(ctx,&pb.GetAccountByEmailRequest{Email: model.Email})
	if err != nil {
		log.Println("Account does not exist: ", err)
		return &pb.LoginResponse{
			Token: "",
			Status: false,
		}, err
	}

	if util.CheckPasswordHash(model.Password, myUser.Password) && myUser.Status.String() == "Active"{
		token, err := jwt.GenerateToken(myUser.Email)
		if err != nil {
			log.Println("Generate token fail: ", err)
			return &pb.LoginResponse{
				Token: "",
				Status: false,
			}, err
		}
		return &pb.LoginResponse{
			Token: token,
			Status: true,
		},nil
	}

	return &pb.LoginResponse{
		Token: "",
		Status: false,
	}, err
}


//Update Account information 
func (us *AccountService) UpdateAccount(ctx context.Context, model *pb.UpdateAccountRequest) (*pb.Account, error){
	res, err := us.AccountRepository.UpdateAccount(ctx, model)
	if err != nil {
		log.Println("Update account failed: ", err)
		return nil, err
	}

	return &pb.Account{
		Id:   res.ID.String(),
		Email: res.Email,
		Role: pb.Account_Role(pb.CreateAccountRequest_Role_value[res.Role.String()]),
		Status: pb.Account_AccountStatus(pb.Account_AccountStatus_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}

func (us *AccountService) UpdateAccountStatus(ctx context.Context, model *pb.UpdateAccountStatusRequest) (*pb.Account, error){
	res, err := us.AccountRepository.UpdateAccountStatus(ctx, model)
	if err != nil {
		log.Println("Update status account failed: ", err)
		return nil, err
	}

	return &pb.Account{
		Id:   res.ID.String(),
		Email: res.Email,
		Role: pb.Account_Role(pb.CreateAccountRequest_Role_value[res.Role.String()]),
		Status: pb.Account_AccountStatus(pb.Account_AccountStatus_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}


//Change password service
func (us *AccountService) AccountResetPassword(ctx context.Context, model *pb.AccountResetPasswordRequest) (*pb.AccountResetPasswordResponse, error){
	res, err := us.AccountRepository.AccountResetPassword(ctx, model)
	if err != nil {
		log.Println("Reset password failed: ", err)
		return nil, err
	}

	return &pb.AccountResetPasswordResponse{
		IsUpdate: res,
	}, nil
}