package theatres

type service struct {
	repo Repository
}

type Service interface {
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
