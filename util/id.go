package util

import (
	"strconv"
	"sync"
)

var (
	requestCount  int
	requestMutex  sync.Mutex
	providerCount int
	providerMutex sync.Mutex
)

func GenerateRequestId() string {
	requestMutex.Lock()
	defer requestMutex.Unlock()
	requestCount++
	return strconv.Itoa(requestCount)
}

func GenerateProviderId() string {
	providerMutex.Lock()
	defer providerMutex.Unlock()
	providerCount++
	return strconv.Itoa(providerCount)
}
