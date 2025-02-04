package entity

import "fmt"

type Provider struct {
	State            bool                                                `json:"state"`
	ProviderId       string                                              `json:"id"`
	Name             string                                              `json:"name"`
	Username         string                                              `json:"username"`
	Password         string                                              `json:"password"`
	DefaultChannels  map[RequestType]ApiEndpoint                         `json:"defaultChannels"`
	CriticalChannels map[CritialCommuncation]map[RequestType]ApiEndpoint `json:"criticalChannels"`
}

func (provider *Provider) Id() string {
	return provider.ProviderId
}

type ApiEndpoint struct {
	Api string `json:"api"`
}

func (p Provider) ProcessRequest(request Request) error {
	if request.Type() == EMAIL {
		processEmailRequest(request)
	} else if request.Type() == SMS {
		processSmsRequest(request)
	} else if request.Type() == SOUNDBOX {
		processSoundBoxRequest(request)
	} else {
		fmt.Println("Unknown Request Type")
	}
	return nil
}

func processEmailRequest(request Request) {
	fmt.Println("Email Request Processed")
}

func processSmsRequest(request Request) {
	fmt.Println("Sms Request Processed")
}

func processSoundBoxRequest(request Request) {
	fmt.Println("SoundBox Request Processed")
}
