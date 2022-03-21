package mqtt

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"time"
)

type client struct {
	options *Options
	client  mqtt.Client
}

type Options struct {
	ConnectionString string
	Username         string
	Password         string
}

type ClientInterface interface {
	Connect()
	Sub(topic string, callback func(msg map[string]interface{}))
	Publish(topic, msg string)
}

func NewClient(options *Options) ClientInterface {
	return &client{
		options: options,
	}
}

func (c *client) Connect() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.options.ConnectionString)
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(c.options.Username)
	opts.SetPassword(c.options.Password)
	opts.OnConnect = func(client mqtt.Client) {
		log.Infof("Connected to %s", c.options.ConnectionString)
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Infof("Connect lost: %v", err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	c.client = client
}

func (c *client) Sub(topic string, callback func(msg map[string]interface{})) {
	token := c.client.Subscribe(topic, 1, func(c mqtt.Client, msg mqtt.Message) {
		log.Infof("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

		var i map[string]interface{}
		_ = json.Unmarshal(msg.Payload(), &i)

		callback(i)


	})
	token.Wait()

	log.Infof("Subscribed to topic %s", topic)
}

func (c *client) Publish(topic, msg string) {
	token := c.client.Publish(topic, 0, false, msg)
	token.Wait()
	time.Sleep(time.Second)
}

