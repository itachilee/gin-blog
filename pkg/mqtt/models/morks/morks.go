package morks

import (
	"collyD/pkg/mqtt/message"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mitchellh/mapstructure"
)

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

type Package struct {
	MessageID      int       `json:"MessageId"`
	MessageVersion int       `json:"MessageVersion"`
	ClientID       int       `json:"ClientId"`
	ClientType     int       `json:"ClientType"`
	Timestamp      time.Time `json:"Timestamp"`

	Message map[string]interface{}
}

func Deserialize(input string) *Package {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(input), &m)
	if err != nil {
		log.Fatalln(err)

	}

	if !message.CheckMessage(m["MessageId"].(int)) {
		log.Fatalf("该消息id非法 %s", m["MessageId"])
	}

	switch m["MessageId"].(int) {
	case message.ZMJ_HEART_BEAT:
		var heartBeat HeartBeat
		mapstructure.Decode(m, &heartBeat)
		fmt.Println("person", heartBeat)

	case message.ZMJ_ALLOW_PUSH_ORDER:
		var allwoPushOrder AllwoPushOrder
		mapstructure.Decode(m, &allwoPushOrder)
		fmt.Println("cat", allwoPushOrder)
	}

	return nil
}
