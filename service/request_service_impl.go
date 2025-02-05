package service

import (
	"fmt"
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
	fmt.Println("Provider selected : ", provider.Name)
	err = provider.ProcessRequest(request)
	return err
}

func (r *RequestServiceImpl) Callback(requestId string, data []byte) error {
	fmt.Println("callback for request ", requestId)
	_, err := r.RequestRepository.UpdateRequest(requestId, true)
	defer fmt.Println("Something to do with response data : ", data)
	return err
}
