package entity

import "phonePe/util"

type SoundBoxRequest struct {
	DefaultRequest
}

func (s *SoundBoxRequest) Type() RequestType {
	return SOUNDBOX
}

func (s *SoundBoxRequest) Validate() (bool, error) {
	return true, nil
}

func (e *SoundBoxRequest) Prefix() string {
	return "SB"
}

func (e *SoundBoxRequest) Id() string {
	return e.RequestId
}

func (e *SoundBoxRequest) SetId() {
	e.RequestId = e.Prefix() + util.GenerateRequestId()
}
