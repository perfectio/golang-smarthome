package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang-smarthome/src/server"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	app := cli.NewApp()
	app.Name = "golang-smarthome"
	app.Usage = ""
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			EnvVars: []string{"DEBUG"},
			Aliases: []string{"d"},
			Value:   false,
			Usage:   "Enable debug mode. Default false",
		},
		&cli.StringFlag{
			Name:     "mqtt-host",
			EnvVars:  []string{"MQTT_HOST"},
			Required: true,
			Usage:    "Mqtt broker host",
		},
		&cli.StringFlag{
			Name:     "mqtt-port",
			EnvVars:  []string{"MQTT_PORT"},
			Required: true,
			Usage:    "Mqtt broker port",
		},
		&cli.StringFlag{
			Name:     "mqtt-username",
			EnvVars:  []string{"MQTT_USERNAME"},
			Required: true,
			Usage:    "Mqtt broker username",
		},
		&cli.StringFlag{
			Name:     "mqtt-password",
			EnvVars:  []string{"MQTT_PASSWORD"},
			Required: true,
			Usage:    "Mqtt broker password",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		err := server.Run(&server.Options{
			MqttHost:     ctx.String("mqtt-host"),
			MqttPort:     ctx.Int("mqtt-port"),
			MqttUsername: ctx.String("mqtt-username"),
			MqttPassword: ctx.String("mqtt-password"),
		})

		return err
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
