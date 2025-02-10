package repository

import (
	"errors"
	"phonePe/entity"
	"sync"
)

type RequestRepositoryImpl struct {
	requests map[string]entity.Request
	mutex    *sync.Mutex
}

func NewRequestRepositoryImpl() *RequestRepositoryImpl {
	return &RequestRepositoryImpl{
		requests: make(map[string]entity.Request),
		mutex:    &sync.Mutex{},
	}
}

func (r *RequestRepositoryImpl) UpdateRequest(requestId string, processed bool) (entity.Request, error) {
	request, ok := r.requests[requestId]
	if !ok {
		return nil, errors.New("Request Not Found")
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	request.Process(processed)
	r.requests[requestId] = request
	return request, nil
}

func (r *RequestRepositoryImpl) GetRequest(requestId string) (entity.Request, error) {
	request, ok := r.requests[requestId]
	if !ok {
		return nil, errors.New("Request Not Found")
	}
	return request, nil
}

func (r *RequestRepositoryImpl) CreateRequest(request entity.Request) (entity.Request, error) {
	r.requests[request.Id()] = request
	return request, nil
}
