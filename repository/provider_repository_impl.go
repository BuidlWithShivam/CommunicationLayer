package repository

import (
	"errors"
	"phonePe/entity"
)

type ProviderRepositoryImpl struct {
	Providers map[string]*entity.Provider
}

func (p *ProviderRepositoryImpl) GetProvidersForRequest(request entity.Request) ([]*entity.Provider, error) {
	critical := request.IsCritical()
	communication := request.CriticalMessage()
	var providers []*entity.Provider
	for _, provider := range p.Providers {
		if !provider.State {
			continue
		}
		if !critical {
			for channel, _ := range provider.DefaultChannels {
				if channel == request.Type() {
					providers = append(providers, provider)
				}
			}
		} else {
			for comm, accounts := range provider.CriticalChannels {
				if comm == communication {
					for channel, _ := range accounts {
						if channel == request.Type() {
							providers = append(providers, provider)
						}
					}
				}
			}
		}
	}
	return providers, nil
}

func (p *ProviderRepositoryImpl) CreateProvider(provider *entity.Provider) (*entity.Provider, error) {
	p.Providers[provider.Id()] = provider
	return provider, nil
}

func (p *ProviderRepositoryImpl) UpdateProvider(provider *entity.Provider) (*entity.Provider, error) {
	_, ok := p.Providers[provider.Id()]
	if !ok {
		return provider, errors.New("provider not found")
	}
	p.Providers[provider.Id()] = provider
	return provider, nil
}

func (p *ProviderRepositoryImpl) DeleteProvider(provider *entity.Provider) error {
	delete(p.Providers, provider.Id())
	return nil
}

func (p *ProviderRepositoryImpl) GetProvider(providerId string) (*entity.Provider, error) {
	provider, ok := p.Providers[providerId]
	if !ok {
		return provider, errors.New("provider not found")
	}
	return provider, nil
}

func (p *ProviderRepositoryImpl) UpdateState(providerId string, state bool) error {
	provider, ok := p.Providers[providerId]
	if !ok {
		return errors.New("provider not found")
	}
	provider.State = state
	return nil
}

func NewProviderRepositoryImpl() *ProviderRepositoryImpl {
	return &ProviderRepositoryImpl{
		Providers: make(map[string]*entity.Provider),
	}
}
