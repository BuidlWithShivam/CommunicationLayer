package repository

import (
	"errors"
	"phonePe/entity"
	"sync"
)

type ProviderRepositoryImpl struct {
	providers map[string]*entity.Provider
	mutex     *sync.Mutex
}

func NewProviderRepositoryImpl() *ProviderRepositoryImpl {
	return &ProviderRepositoryImpl{
		providers: make(map[string]*entity.Provider),
		mutex:     &sync.Mutex{},
	}
}

func (p *ProviderRepositoryImpl) CreateProvider(provider *entity.Provider) (*entity.Provider, error) {
	p.providers[provider.Id()] = provider
	return provider, nil
}

func (p *ProviderRepositoryImpl) UpdateProvider(provider *entity.Provider) (*entity.Provider, error) {
	_, ok := p.providers[provider.Id()]
	if !ok {
		return provider, errors.New("provider not found")
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.providers[provider.Id()] = provider
	return provider, nil
}

func (p *ProviderRepositoryImpl) GetProvider(providerId string) (*entity.Provider, error) {
	provider, ok := p.providers[providerId]
	if !ok {
		return provider, errors.New("provider not found")
	}
	return provider, nil
}

func (p *ProviderRepositoryImpl) UpdateState(providerId string, state bool) error {
	provider, ok := p.providers[providerId]
	if !ok {
		return errors.New("provider not found")
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	provider.State = state
	return nil
}

func (p *ProviderRepositoryImpl) GetProvidersForRequest(request entity.Request) ([]*entity.Provider, error) {
	critical := request.IsCritical()
	communication := request.CriticalMessage()
	var providers []*entity.Provider
	for _, provider := range p.providers {
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
