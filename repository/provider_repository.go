package repository

import "phonePe/entity"

type ProviderRepository interface {
	CreateProvider(*entity.Provider) (*entity.Provider, error)
	UpdateProvider(*entity.Provider) (*entity.Provider, error)
	DeleteProvider(*entity.Provider) error
	GetProvider(string) (*entity.Provider, error)
	UpdateState(string, bool) error
	GetProvidersForRequest(request entity.Request) ([]*entity.Provider, error)
}
