package dao

type ResourceRepositoryInterface interface {
}

type ResourceRepository struct {
}

func NewResourceRepository() ResourceRepositoryInterface {
	return &ResourceRepository{}
}
