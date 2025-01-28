package repo

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct{}

// Get user by email
func (ur *userRepo) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
