package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang-smarthome/src/mqtt"
	xiaomi_cube "golang-smarthome/src/mqtt/device/xiaomi-cube"
	"os"
	"os/signal"
)

var (
	done <-chan bool
	on   bool
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

	client.Sub("zigbee2mqtt/0x00158d0001148311", func(msg map[string]interface{}) {
		event := &xiaomi_cube.Event{}
		_ = event.LoadFromMap(msg)

		log.Infof("%+v", event)

		if event.Action == "flip90" {
			if on {
				client.Publish("cmnd/tasmota_19EE45/POWER", "OFF")
				on = false
			} else {
				client.Publish("cmnd/tasmota_19EE45/POWER", "ON")
				on = true
			}
		}
	})

	/*	cfg := config.Load()

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
		}*/

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done

	return nil
}
