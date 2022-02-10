package mqtt

import (
	"collyD/pkg/setting"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

/**
 * @Description:订阅回调
 * @param client
 * @param msg
 */

func subCallBackFunc(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("订阅: 当前话题是 [%s]; 信息是 [%s] \n", msg.Topic(), string(msg.Payload()))
}
func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, subCallBackFunc)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func SetUp() {
	// var c = make(chan bool)
	client, err := createMqttClient()
	if err != nil {
		panic(err)
	}
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		// c <- true
		panic(token.Error())
	}
	sub(client)
	publish(client)
	// <-c
}

// 创建一个mqtt连接
func createMqttClient() (mqtt.Client, error) {
	var broker = setting.MqttSetting.Host
	var port = setting.MqttSetting.Port
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go-" + randStr(10))
	opts.SetUsername(setting.MqttSetting.Username)
	opts.SetPassword(setting.MqttSetting.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	return client, nil
}
