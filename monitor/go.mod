module github.com/ddaws/maker-monitor

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/ddaws/go-maker v0.0.0-20200319112206-e06005d816d3
	github.com/ethereum/go-ethereum v1.12.1
	github.com/gorilla/websocket v1.5.0
	github.com/ilyakaznacheev/cleanenv v1.2.1
	github.com/prometheus/client_golang v1.14.0
	github.com/shopspring/decimal v0.0.0-20200227202807-02e2044944cc
)

//replace github.com/ddaws/go-maker => ../../go-maker/
