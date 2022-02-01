package activitygroups

type Service struct {
	Repository Repository
}

func LoadService(repository Repository) *Service {
	return &Service{
		Repository: repository,
	}
}
