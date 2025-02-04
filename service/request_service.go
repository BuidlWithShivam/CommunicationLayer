package service

import "phonePe/entity"

type RequestService interface {
	ProcessRequest(request entity.Request)
}
