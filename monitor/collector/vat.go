package collector

import (
	"log"

	"github.com/ddaws/go-maker/maker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shopspring/decimal"
)

// TODO: Add support for measuring Ilks, urns, gem, dai, and sin
var (
	debtDesc = prometheus.NewDesc(
		"mkr_vat_debt",
		"Maker Vat debt, aka the total Dai in issuance (rad)",
		nil, nil,
	)
	viceDesc = prometheus.NewDesc(
		"mkr_vat_vice",
		"Maker Vat vice, aka total unbacked Dai (rad)",
		nil, nil,
	)
	lineDesc = prometheus.NewDesc(
		"mkr_vat_line",
		"Maker Vat line, aka the debt ceiling (rad)",
		nil, nil,
	)
)

type vatCollector struct {
	vat *maker.VatCaller
}

// NewVatCollector reates a new collecotr for collecting Vat specific stats.
//
// TODO: Add support for collecting metrics on Ilks, Urns, and individual Vaults
func NewVatCollector(vat *maker.VatCaller) prometheus.Collector {
	return &vatCollector{vat}
}

func (c *vatCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- debtDesc
	ch <- viceDesc
	ch <- lineDesc
}

func (c *vatCollector) Collect(ch chan<- prometheus.Metric) {
	// Measure the total Dai in issuance
	if debt, err := c.vat.Debt(nil); err == nil {
		debtDec := decimal.NewFromBigInt(debt, -maker.RadScale)
		debtApprox, _ := debtDec.Float64()

		log.Printf("vatCollector.Collect(debtApprox=%.4f)", debtApprox)

		ch <- prometheus.MustNewConstMetric(
			debtDesc,
			prometheus.GaugeValue,
			debtApprox,
		)
	}
	// Meausre the total unbacked Dai
	if vice, err := c.vat.Vice(nil); err == nil {
		viceDec := decimal.NewFromBigInt(vice, -maker.RadScale)
		viceApprox, _ := viceDec.Float64()

		ch <- prometheus.MustNewConstMetric(
			viceDesc,
			prometheus.GaugeValue,
			viceApprox,
		)
	}
	// Measure the Maker debt ceiling, aka the total Dai available for issuance
	if line, err := c.vat.Line(nil); err == nil {
		lineDec := decimal.NewFromBigInt(line, -maker.RadScale)
		lineApprox, _ := lineDec.Float64()

		ch <- prometheus.MustNewConstMetric(
			lineDesc,
			prometheus.GaugeValue,
			lineApprox,
		)
	}
}
