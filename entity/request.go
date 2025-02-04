package entity

type RequestType int

const (
	EMAIL RequestType = iota
	SMS
	SOUNDBOX
)

func (r RequestType) String() string {
	switch r {
	case EMAIL:
		return "EMAIL"
	case SMS:
		return "SMS"
	case SOUNDBOX:
		return "SOUNDBOX"
	default:
		return "Unknown"
	}
}

type Request interface {
	Id() string
	Type() RequestType
	Validate() (bool, error)
	Prefix() string
	Process(bool)
	IsCritical() bool
	CriticalMessage() CritialCommuncation
	SetId()
}
