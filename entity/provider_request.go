package entity

type ProviderRequest struct {
	Name             string                                              `json:"name"`
	UserName         string                                              `json:"userName"`
	Password         string                                              `json:"password"`
	DefaultChannels  map[RequestType]ApiEndpoint                         `json:"defaultChannels"`
	CriticalChannels map[CritialCommuncation]map[RequestType]ApiEndpoint `json:"criticalChannels"`
}
