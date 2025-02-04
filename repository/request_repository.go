package repository

import (
	"phonePe/entity"
)

type RequestRepository interface {
	CreateRequest(request entity.Request) (entity.Request, error)
	UpdateRequest(string, bool) (entity.Request, error)
	GetRequest(string) (entity.Request, error)
}
