package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang-smarthome/config"
	"golang-smarthome/mqtt"
	"os"
	"os/signal"
)

var (
	done <-chan bool
)

type Options struct {
	MqttHost     string
	MqttPort     int
	MqttUsername string
	MqttPassword string
}

func Run(options *Options) error {
	client := mqtt.NewClient(&mqtt.Options{
		ConnectionString: fmt.Sprintf("tcp://%s:%d", options.MqttHost, options.MqttPort),
		Username:         options.MqttUsername,
		Password:         options.MqttPassword,
	})
	client.Connect()

	cfg := config.Load()

	for _, dc := range cfg.Devices {
		if dc.Endpoints.Event != "" {
			log.Infof("Device [%s] initializing", dc.Name)
			client.Sub(dc.Endpoints.Event, func(msg map[string]interface{}) {
				topic := ""
				if msg["click"] == "single" {
					client.Publish(topic, "ON")
				}

				if msg["click"] == "double" {
					client.Publish(topic, "OFF")
				}
			})
		}
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done

	return nil
}
