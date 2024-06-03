package usecases

import "go-community/internal/repositories/pgsql"

type Dependencies struct {
	Repository	*pgsql.PostgreRepositories
}

type Usecases struct {
	Health	healthUsecase
	Campus	campusUsecase
	CoolCategory coolCategoryUsecase
	Location locationUsecase
}

func New(d Dependencies) *Usecases{
	health := NewHealthUsecase(d.Repository.Health)
	campus := NewCampusUsecase(d.Repository.Campus)
	coolCategory := NewCoolCategoryUsecase(d.Repository.CoolCategory)
	location := NewLocationUsecase(d.Repository.Location, d.Repository.Campus)

	return &Usecases{
		Health: *health,
		Campus: *campus,
		CoolCategory: *coolCategory,
		Location: *location,
	}
}