package service

import (
	"errors"
	"math/rand"
	"phonePe/entity"
)

type RandomAllocationService struct {
}

func NewRandomAllocationService() *RandomAllocationService {
	return &RandomAllocationService{}
}

func (r *RandomAllocationService) AllocateProvider(providers []*entity.Provider) (*entity.Provider, error) {
	if len(providers) == 0 {
		return nil, errors.New("provider is empty")
	}
	return providers[rand.Intn(len(providers))], nil
}
