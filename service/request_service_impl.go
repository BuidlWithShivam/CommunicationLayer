package service

import (
	"phonePe/entity"
	"phonePe/repository"
)

type RequestServiceImpl struct {
	RequestRepository repository.RequestRepository
	Service           ProviderService
}

func NewRequestServiceImpl(requestRepository repository.RequestRepository, service ProviderService) *RequestServiceImpl {
	return &RequestServiceImpl{
		RequestRepository: requestRepository,
		Service:           service,
	}
}

func (r *RequestServiceImpl) ProcessRequest(request entity.Request) error {
	valid, err := request.Validate()
	if !valid {
		return err
	}
	_, err = r.RequestRepository.CreateRequest(request)
	if err != nil {
		return err
	}
	provider, err := r.Service.GetProviderForRequest(request)
	if err != nil {
		return err
	}
	err = provider.ProcessRequest(request)
	return err
}
