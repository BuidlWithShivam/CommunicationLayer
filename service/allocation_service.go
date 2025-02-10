package service

import "phonePe/entity"

type AllocationService interface {
	AllocateProvider([]*entity.Provider) (*entity.Provider, error)
}
