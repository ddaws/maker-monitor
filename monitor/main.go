package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ddaws/go-maker/maker"
	"github.com/ddaws/maker-monitor/collector"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type configType struct {
	Port   string `yaml:"port" env:"PORT" env-default:"8080"`
	Host   string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
	Infura struct {
		Network   string `yaml:"network" env:"INFURA_NETWORK" env-default:"mainnet"`
		ProjectID string `yaml:"projectID" env:"INFURA_PROJECT_ID"`
	} `yaml:"infura"`
}

const (
	infuraWebSocketAddr = "wss://%s.infura.io/ws/v3/%s"
)

var (
	configFile string
	config     configType
	// Collectors
	headerCollector = collector.NewHeaderCollector()
)

func init() {
	prometheus.MustRegister(headerCollector)
	// Add Go module build info and stats
	//prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	//prometheus.MustRegister(prometheus.NewGoCollector())
}

func main() {
	log.Println("Starting the Maker Monitor metrics server...")

	flag.StringVar(&configFile, "config", "config.yml", "Path to the config.yml file")
	flag.Parse()

	if !filepath.IsAbs(configFile) {
		absConfigPath, err := filepath.Abs(configFile)
		if err != nil {
			log.Fatalln(err)
		}
		configFile = absConfigPath
	}

	log.Println("Reading config...")
	err := cleanenv.ReadConfig(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	if config.Infura.ProjectID == "" {
		log.Fatalln("An Infura project ID must be specified to subscribe to Infura")
	}

	// Connect to Infura
	log.Println("Connecting to Infura...")
	infuraEndpoint := fmt.Sprintf(infuraWebSocketAddr, config.Infura.Network, config.Infura.ProjectID)
	client, err := ethclient.Dial(infuraEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Infura!")

	// Load the Vat and collector
	vat, err := maker.LoadVatCaller(client)
	if err != nil {
		log.Fatalln(err)
	}
	vatCollector := collector.NewVatCollector(vat)
	prometheus.MustRegister(vatCollector)

	// Load the Pot and collector
	//pot, err := maker.LoadPotCaller(client)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//potCollector := collector.NewPotCollector(pot)
	//prometheus.MustRegister(potCollector)

	// Start listening for blocks mined
	headers := make(chan *types.Header)
	go listenForBlocks(client, headers, headerCollector)

	// Expose a metrics endpoint for Prometheus scraping
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), nil)
}

func listenForBlocks(client *ethclient.Client, headers chan *types.Header, headerCollector *collector.HeaderCollector) {
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
			headerCollector.Measure(header)
		}
	}
}
