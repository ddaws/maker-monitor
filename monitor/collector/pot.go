package collector

import (
	"log"

	"github.com/ddaws/go-maker/maker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shopspring/decimal"
)

const (
	secondsPerYear = float64(365 * 24 * 60 * 60)
	roundPrecision = 0.001
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
	dsrAnnualizedDesc = prometheus.NewDesc(
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
	// Measure the total Dai savings accumulated
	if pieRad, err := c.pot.Pie(); err == nil {
		pie := decimal.NewFromBigInt(pieRad, -maker.RadScale)
		pieApprox, _ := pie.Float64()

		ch <- prometheus.MustNewConstMetric(
			pieDesc,
			prometheus.GaugeValue,
			pieApprox,
		)
	}
	// Measure the Dai Savings Rate
	if dsrRad, err := c.pot.Dsr(); err == nil {
		dsr := decimal.NewFromBigInt(dsrRad, -maker.RadScale)
		dsrApprox, _ := dsr.Float64()

		ch <- prometheus.MustNewConstMetric(
			dsrDesc,
			prometheus.GaugeValue,
			dsrApprox,
		)
		// Calculate annualized DSR and round to roundPrecision
		dsrAnnualized = math.Pow(dsrApprox, secondsPerYear)
		dsrAnnualized = math.Round(dsrAnnualized / roundPrecision) * roundPrecision

		ch <- prometheus.MustNewConstMetric(
			dsrAnnualizedDesc,
			prometheus.GaugeValue,
			dsrAnnualized
		)
	}
}