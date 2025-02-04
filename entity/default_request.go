package entity

import "time"

type CritialCommuncation int

const (
	OTPComm CritialCommuncation = iota
)

func (c CritialCommuncation) String() string {
	switch c {
	case OTPComm:
		return "OTP"
	default:
		return "Unknown"
	}
}

type DefaultRequest struct {
	RequestId   string              `json:"requestId"`
	Date        time.Time           `json:"date"`
	Processed   bool                `json:"processed"`
	Critical    bool                `json:"critical"`
	CriticalMsg CritialCommuncation `json:"criticalMsg"`
	CallbackUrl string              `json:"callbackUrl"`
}

func (d *DefaultRequest) Process(processed bool) {
	d.Processed = processed
}

func (d *DefaultRequest) IsCritical() bool {
	return d.Critical
}

func (d *DefaultRequest) CriticalMessage() CritialCommuncation {
	return d.CriticalMsg
}
