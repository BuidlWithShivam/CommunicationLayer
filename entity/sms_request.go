package entity

import (
	"errors"
	"phonePe/util"
)

type SmsRequest struct {
	DefaultRequest
	Mobile  string `json:"mobile"`
	Message string `json:"message"`
}

func (e SmsRequest) Type() RequestType {
	return SMS
}

func (e SmsRequest) Validate() (bool, error) {
	if e.Mobile == "" {
		return false, errors.New("Mobile is required")
	}
	if e.Message == "" {
		return false, errors.New("Message is required")
	}
	return true, nil
}

func (e SmsRequest) Prefix() string {
	return "SM"
}

func (e SmsRequest) Id() string {
	return e.RequestId
}

func (e SmsRequest) SetId() {
	e.RequestId = e.Prefix() + util.GenerateRequestId()
}
