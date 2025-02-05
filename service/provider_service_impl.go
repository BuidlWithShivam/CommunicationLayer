package service

import (
	"errors"
	"math/rand"
	"phonePe/entity"
	"phonePe/repository"
	"phonePe/util"
)

type ProviderServiceImpl struct {
	providerRepository repository.ProviderRepository
}

func NewProviderServiceImpl(providerRepository repository.ProviderRepository) *ProviderServiceImpl {
	return &ProviderServiceImpl{
		providerRepository: providerRepository,
	}
}

func (p ProviderServiceImpl) AddProvider(request *entity.ProviderRequest) (*entity.Provider, error) {
	provider := &entity.Provider{
		Name:             request.Name,
		State:            true,
		Username:         request.UserName,
		Password:         request.Password,
		DefaultChannels:  request.DefaultChannels,
		CriticalChannels: request.CriticalChannels,
	}
	provider.ProviderId = util.GenerateProviderId()
	return p.providerRepository.CreateProvider(provider)
}

func (p ProviderServiceImpl) GetProvider(providerId string) (*entity.Provider, error) {
	return p.providerRepository.GetProvider(providerId)
}

func (p ProviderServiceImpl) UpdateState(providerId string, state bool) error {
	return p.providerRepository.UpdateState(providerId, state)
}

func (p ProviderServiceImpl) UpdateProvider(provider *entity.Provider) (*entity.Provider, error) {
	return p.providerRepository.UpdateProvider(provider)
}

func (p ProviderServiceImpl) GetProviderForRequest(request entity.Request) (*entity.Provider, error) {
	providers, _ := p.providerRepository.GetProvidersForRequest(request)
	return p.AllocateProvider(providers)
}

func (p ProviderServiceImpl) AllocateProvider(providers []*entity.Provider) (*entity.Provider, error) {
	if len(providers) == 0 {
		return nil, errors.New("provider is empty")
	}
	return providers[rand.Intn(len(providers))], nil
}
