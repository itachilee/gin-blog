package morks

import "time"

type HeartBeat struct {
	MessageID      int       `json:"MessageId"`
	MessageVersion int       `json:"MessageVersion"`
	ClientID       int       `json:"ClientId"`
	ClientType     int       `json:"ClientType"`
	Timestamp      time.Time `json:"Timestamp"`
	Message        struct {
		TaskCount int `json:"TaskCount"`
		Busy      int `json:"Busy"`
	} `json:"Message"`
}

type AllwoPushOrder struct {
	MessageID      int    `json:"MessageId"`
	MessageVersion int    `json:"MessageVersion"`
	ClientID       int    `json:"ClientId"`
	ClientType     int    `json:"ClientType"`
	Timestamp      string `json:"Timestamp"`
	Message        struct {
		SubOrderID   string `json:"SubOrderId"`
		CanPushOrder int    `json:"CanPushOrder"`
	} `json:"Message"`
}
