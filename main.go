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
	allocationService := service.NewRandomAllocationService()
	providerService := service.NewProviderServiceImpl(providerRepository, allocationService)

	requestRepository := repository.NewRequestRepositoryImpl()
	requestService := service.NewRequestServiceImpl(requestRepository, providerService)

	_, err := providerService.AddProvider(&entity.ProviderRequest{
		Name:     "Provider1",
		UserName: "test",
		Password: "test",
		DefaultChannels: map[entity.RequestType]entity.ApiEndpoint{
			entity.EMAIL: {Api: "email1"},
			entity.SMS:   {Api: "sms1"},
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

	p2, err := providerService.AddProvider(&entity.ProviderRequest{
		Name:     "Provider2",
		UserName: "test1",
		Password: "test1",
		DefaultChannels: map[entity.RequestType]entity.ApiEndpoint{
			entity.EMAIL: {Api: "email2"},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Provider 2 created")

	err = providerService.UpdateState(p2.ProviderId, false)
	if err != nil {
		panic(err)
	}

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

	request = &entity.EmailRequest{
		DefaultRequest: entity.DefaultRequest{},
		Sender:         "test2",
		Receiver:       "test1",
		Subject:        "abcde",
		Message:        "qwerty1",
	}
	request.SetId()
	err = requestService.ProcessRequest(request)

	if err != nil {
		panic(err)
	}
	fmt.Println("Request 2 processed")

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
