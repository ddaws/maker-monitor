#!/bin/bash

echo "Installing Go Dependancies"
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/ethereum/go-ethereum
go get github.com/deckarep/golang-set
go get github.com/gorilla/websocket
go get github.com/rs/cors
go get github.com/ilyakaznacheev/cleanenv
