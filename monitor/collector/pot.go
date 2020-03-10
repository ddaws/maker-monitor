package collector

import (
	"log"

	"github.com/ddaws/go-maker/maker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shopspring/decimal"
)

// TODO: Add support for measuring Ilks, urns, gem, dai, and sin
var (
	pieDesc = prometheus.NewDesc(
		"mkr_pot_pie",
		"Maker Pot pie, aka to total Dai savings accrued (rad)"
	)
	dsrDesc = prometheus.NewDesc(
		"mkr_pot_dsr",
		"Maker Pot dsr, aka the Dai Savings Rate (rad)"
	)
	dsrAnnualized = prometheus.NewDesc(
		"mkr_pot_dsr_apy",
		"Maker Pot dsr annualized, aka the DSR as an annual percent return"
	)
	rowDesc = prometheus.NewDesc(
		"mkr_pot_row",
		"Maker Pot row, aka the last drip call"
	)
)

type potCollector struct {
	pot *maker.PotCaller
}

func (c *potCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pieDesc
	ch <- dsrDesc
	ch <- dsrAnnualized
	ch <- rowDesc
}

func (c *potCollector) Collect(ch chan<- prometheus.Metric) {
	// TODO...
}