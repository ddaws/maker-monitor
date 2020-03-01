package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	Port   string `yaml:"port" env:"PORT" env-default:"8080"`
	Host   string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
	Infura struct {
		Network   string `yaml:"network" env:"INFURA_NETWORK" env-default:"ropsten"`
		ProjectID string `yaml:"projectID" env:"INFURA_PROJECT_ID"`
	} `yaml:"infura"`
}

const (
	InfuraWebSocketAddr = "wss://%s.infura.io/ws/v3/%s"
)

var (
	configFile   string
	config       Config
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "maker_processed_ops_total",
		Help: "The total number of processed events",
		ConstLabels: map[string]string{
			"type": "Fake",
		},
	})
)

func init() {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

func main() {
	fmt.Println("Starting the Maker Monitor metrics server...")

	flag.StringVar(&configFile, "config", "config.yml", "Path to the config.yml file")
	flag.Parse()

	if !filepath.IsAbs(configFile) {
		absConfigPath, err := filepath.Abs(configFile)
		if err != nil {
			log.Fatalln(err)
		}
		configFile = absConfigPath
	}

	fmt.Println("Reading config...")
	err := cleanenv.ReadConfig(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	if config.Infura.ProjectID == "" {
		log.Fatalln("An Infura project ID must be specified to subscribe to Infura")
	}

	// Connect to Infura
	fmt.Println("Connecting to Infura...")
	infuraEndpoint := fmt.Sprintf(InfuraWebSocketAddr, config.Infura.Network, config.Infura.ProjectID)
	client, err := ethclient.Dial(fmt.Sprintf(infuraEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Infura!")

	headers := make(chan *types.Header)
	go listenForBlocks(client, headers)

	go recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), nil)
}

func recordMetrics() {
	for {
		opsProcessed.Inc()
		time.Sleep(2 * time.Second)
	}
}

func listenForBlocks(client *ethclient.Client, headers chan *types.Header) {
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
		}
	}
}
