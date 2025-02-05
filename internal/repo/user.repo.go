package repo

import (
	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/database"
)

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct {
	sqlc *database.Queries
}

// Get user by email
func (ur *userRepo) GetUserByEmail(email string) bool {
	// user, err := ur.sqlc.GetUserByEmailSqlc(ctx, email)
	// if err != nil {
	// 	fmt.Printf("Error get user by email: %v\n", err)
	// 	return false
	// }

	return true
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlc: database.New(global.Mdbc),
	}
}
