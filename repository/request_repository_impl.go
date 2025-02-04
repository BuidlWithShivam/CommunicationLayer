package repository

import (
	"errors"
	"phonePe/entity"
)

type RequestRepositoryImpl struct {
	Requests map[string]entity.Request
}

func NewRequestRepositoryImpl() *RequestRepositoryImpl {
	return &RequestRepositoryImpl{
		Requests: make(map[string]entity.Request),
	}
}

func (r *RequestRepositoryImpl) UpdateRequest(requestId string, processed bool) (entity.Request, error) {
	request, ok := r.Requests[requestId]
	if !ok {
		return nil, errors.New("Request Not Found")
	}
	request.Process(processed)
	r.Requests[requestId] = request
	return request, nil
}

func (r *RequestRepositoryImpl) GetRequest(requestId string) (entity.Request, error) {
	request, ok := r.Requests[requestId]
	if !ok {
		return nil, errors.New("Request Not Found")
	}
	return request, nil
}

func (r *RequestRepositoryImpl) CreateRequest(request entity.Request) (entity.Request, error) {
	r.Requests[request.Id()] = request
	return request, nil
}
