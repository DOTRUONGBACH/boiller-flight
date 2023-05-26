package repository
import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/account"
	"mock_project/pb"
	"os"
	_ "github.com/lib/pq"
	"github.com/google/uuid"
	"mock_project/internal/util"
)

type AccountRepository interface {
	GetAccountByID(ctx context.Context, model *pb.GetAccountByIdRequest) (*ent.Account, *ent.Customer, error)
	GetAccountByEmail(ctx context.Context, model *pb.GetAccountByEmailRequest) (*ent.Account, error)
	CreateAccount(ctx context.Context, model *pb.CreateAccountRequest) (*ent.Account, error)
	UpdateAccount(ctx context.Context, model *pb.UpdateAccountRequest) (*ent.Account, error)
	UpdateAccountStatus(ctx context.Context, model *pb.UpdateAccountStatusRequest) (*ent.Account, error)
	AccountResetPassword(ctx context.Context, model *pb.AccountResetPasswordRequest) (string, error)
	DeleteAccount(ctx context.Context, model *pb.DeleteAccountRequest) (bool ,error)
	CloseDB()
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (AccountRepository, error) {
	client, err := ent.Open("postgres", connection_str)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to creating db schema from the automation migration tool")
		os.Exit(1)
	}

	return &PostgresDB{Client: client}, nil
}


func (r *PostgresDB) CloseDB() {
	r.Client.Close()
}

func (r *PostgresDB) GetAccountByID(ctx context.Context, model *pb.GetAccountByIdRequest) (*ent.Account, *ent.Customer, error){
	res1, err1 := r.Client.Account.Query().Where(account.ID(uuid.MustParse(model.Id))).Only(ctx)

	if err1 != nil {
		log.Println("Account not found", err1)
		return nil, nil,err1
	}

	res2, err2 := res1.QueryAccOwner().Only(ctx)

	if err2 != nil {
		log.Println("Account connection not found", err2)
		return nil, nil,err1
	}
	
	return res1, res2, nil
}

func (r *PostgresDB) GetAccountByEmail(ctx context.Context, model *pb.GetAccountByEmailRequest) (*ent.Account, error){
	res, err := r.Client.Account.Query().Where(account.Email(model.Email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) CreateAccount(ctx context.Context, model *pb.CreateAccountRequest) (*ent.Account, error) {
	res, err := r.Client.Account.Create().SetEmail(model.Email).SetPassword(model.Password).SetRole(account.Role(model.Role.String())).SetStatus(account.StatusInactive).SetAccOwnerID(uuid.MustParse(model.AccOwner)).Save(ctx)
	if err != nil{
		return nil, err
	}

	log.Println("NEW ACCOUNT REPO", res)
	return res, nil
}

func (r *PostgresDB) UpdateAccount(ctx context.Context, model *pb.UpdateAccountRequest) (*ent.Account, error) {
	acc, err1 := r.Client.Account.Query().Where(account.ID(uuid.MustParse(model.Id))).Only(ctx)
	if err1 != nil {
		return nil, err1
	}
	res, err2 := acc.Update().SetRole(account.Role(model.Role.String())).Save(ctx)
	if err2 != nil{
		return nil, err2
	}

	return res, nil
}

func (r *PostgresDB) UpdateAccountStatus(ctx context.Context, model *pb.UpdateAccountStatusRequest) (*ent.Account, error) {
	acc, err1 := r.Client.Account.Query().Where(account.ID(uuid.MustParse(model.Id))).Only(ctx)
	if err1 != nil {
		return nil, err1
	}
	res, err2 := acc.Update().SetStatus(account.Status(model.Status.String())).Save(ctx)
	if err2 != nil{
		return nil, err2
	}

	return res, nil
}

func (r *PostgresDB) AccountResetPassword(ctx context.Context, model *pb.AccountResetPasswordRequest) (string, error) {
	myAccount, err1 := r.Client.Account.Query().Where(account.Email(model.Email)).Only(ctx)
	if err1 != nil {
		return "Account does not exist", err1
	}

	log.Println(myAccount.Password)
	log.Println(model.CurrentPassword)

	if util.CheckPasswordHash(model.CurrentPassword, myAccount.Password){
		newPassword, _ := util.HashPassword(model.NewPassword)
		_, err2 := myAccount.Update().SetEmail(model.Email).SetPassword(newPassword).Save(ctx)
		if err2 != nil {
			return "Reset password fail 1", err2
		}
		return "Success", nil
	}

	return "Reset password fail 2", nil
}

func (r *PostgresDB) DeleteAccount(ctx context.Context, model *pb.DeleteAccountRequest) (bool, error) {
	err := r.Client.Account.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
