package entity

import (
	"errors"
	"phonePe/util"
)

type EmailRequest struct {
	DefaultRequest
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
}

func (e EmailRequest) Type() RequestType {
	return EMAIL
}

func (e EmailRequest) Validate() (bool, error) {
	if e.Sender == "" {
		return false, errors.New("Missing Sender")
	}
	if e.Receiver == "" {
		return false, errors.New("Missing Receiver")
	}
	if e.Subject == "" {
		return false, errors.New("Missing Subject")
	}
	if e.Message == "" {
		return false, errors.New("Missing Message")
	}
	return true, nil
}

func (e EmailRequest) Prefix() string {
	return "EM"
}

func (e EmailRequest) Id() string {
	return e.RequestId
}

func (e EmailRequest) SetId() {
	e.RequestId = e.Prefix() + util.GenerateRequestId()
}
