package service

import "phonePe/entity"

type ProviderService interface {
	AddProvider(request *entity.ProviderRequest) (*entity.Provider, error)
	GetProvider(string) (*entity.Provider, error)
	UpdateState(string, bool) error
	UpdateProvider(*entity.Provider) (*entity.Provider, error)
	GetProviderForRequest(entity.Request) (*entity.Provider, error)
}
