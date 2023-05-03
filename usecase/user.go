package usecase

type UserUseCase struct {
	UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo}
}
