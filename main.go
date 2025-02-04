package main

import (
	"fmt"
	"phonePe/entity"
	"phonePe/repository"
	"phonePe/service"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	providerRepository := repository.NewProviderRepositoryImpl()
	providerService := service.NewProviderServiceImpl(providerRepository)

	requestRepository := repository.NewRequestRepositoryImpl()
	requestService := service.NewRequestServiceImpl(requestRepository, providerService)

	_, err := providerService.AddProvider(&entity.ProviderRequest{
		Name:     "Provider1",
		UserName: "test",
		Password: "test",
		DefaultChannels: map[entity.RequestType]entity.ApiEndpoint{
			entity.EMAIL: entity.ApiEndpoint{Api: "email1"},
			entity.SMS:   entity.ApiEndpoint{Api: "sms1"},
		},
		CriticalChannels: map[entity.CritialCommuncation]map[entity.RequestType]entity.ApiEndpoint{
			entity.OTPComm: {
				entity.EMAIL: entity.ApiEndpoint{Api: "email2"},
				entity.SMS:   entity.ApiEndpoint{Api: "sms2"},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Provider 1 created")

	request := &entity.EmailRequest{
		DefaultRequest: entity.DefaultRequest{},
		Sender:         "test1",
		Receiver:       "test2",
		Subject:        "abcd",
		Message:        "qwerty",
	}
	request.SetId()
	err = requestService.ProcessRequest(request)

	if err != nil {
		panic(err)
	}
	fmt.Println("Request 1 processed")
}

// Request > RequestTypes > DifferentRequest for each type > Callback and processed, critical
// Provider > Multiple Account > Account will have channel > state

// Data Layer:
// Requests with request id
// Provider : List for each channel
// List of critical providers

// service:
// Provider

// Main
