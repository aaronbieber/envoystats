package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aaronbieber/carbonclient"
	"github.com/aaronbieber/envoyclient"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Email               string `yaml:"email"`
	Password            string `yaml:"password"`
	EnvoySerialNo       string `yaml:"envoySerialNo"`
	EnvoyIP             string `yaml:"envoyIP"`
	CarbonIP            string `yaml:"carbonIP"`
	ProductionStatPath  string `yaml:"productionStatPath"`
	ConsumptionStatPath string `yaml:"consumptionStatPath"`
}

func init() {
	log.SetOutput(os.Stdout)
}

func readConfig() (Config, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %v", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding config file: %v", err)
	}

	return cfg, nil
}

func main() {
	log.Info("Starting envoystats...")
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	envoy, err := envoyclient.NewClient(
		envoyclient.Config{
			Email:         cfg.Email,
			Password:      cfg.Password,
			EnvoySerialNo: cfg.EnvoySerialNo,
			EnvoyIP:       cfg.EnvoyIP})
	if err != nil {
		log.WithFields(
			log.Fields{
				"email":         cfg.Email,
				"envoySerialNo": cfg.EnvoySerialNo,
				"envoyIP":       cfg.EnvoyIP,
			}).Fatal(err)
		panic("error creating envoy client")
	}

	data, err := envoy.GetProductionData()
	if err != nil {
		log.Fatal(err)
		panic("error getting production data")
	}

	log.WithFields(log.Fields{
		"productionWattsNow":  data.ProductionWattsNow,
		"consumptionWattsNow": data.ConsumptionWattsNow,
	}).Info("production data received")

	now := time.Now()
	var stat []carbonclient.TimedMetric
	stat = append(stat,
		carbonclient.TimedMetric{
			Path: cfg.ProductionStatPath,
			Value: carbonclient.TimedMetricValue{
				Timestamp: now,
				Value:     data.ProductionWattsNow}})
	stat = append(stat,
		carbonclient.TimedMetric{
			Path: cfg.ConsumptionStatPath,
			Value: carbonclient.TimedMetricValue{
				Timestamp: now,
				Value:     data.ConsumptionWattsNow}})

	client, err := carbonclient.NewCarbonClient(cfg.CarbonIP, carbonclient.PICKLE_PORT)
	if err != nil {
		log.WithFields(log.Fields{
			"carbonIP": cfg.CarbonIP,
			"port":     carbonclient.PICKLE_PORT,
		}).Fatal(err)
	}

	err = client.SendMetrics(stat[:])
	if err != nil {
		log.WithFields(log.Fields{
			"carbonIP":   cfg.CarbonIP,
			"picklePort": carbonclient.PICKLE_PORT,
		}).Fatal(err)

		panic("error sending metrics to carbon")
	}

	log.Info("metrics sent to carbon", stat[:])
}
