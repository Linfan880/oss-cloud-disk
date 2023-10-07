package dao

type UserRepoInterface interface {
	ListResource() error
}

type UserRepo struct {
}

func NewUserRepo() UserRepoInterface {
	return &UserRepo{}
}

func (u *UserRepo) ListResource() error {
	return nil
}
