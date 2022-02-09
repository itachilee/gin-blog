package models

const (
	MorkD = iota + 1
	Ice
	Server
	Other
	MorkS
)

type DeviceType int
type PackageHeader struct {
	MessageId      int        `json:"MessageId"`
	MessageVersion byte       `json:"MessageVersion"`
	ClientId       int        `json:"ClientId"`
	DeviceType     DeviceType `json:"DeviceType"`
	DateTime       string     `json:"DateTime"`
}

type PackageBody struct {
}
