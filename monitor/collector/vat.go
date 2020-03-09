package collector

import (
	"github.com/ddaws/go-maker/maker"
	"github.com/prometheus/client_golang/prometheus"
)

// TODO: Add support for measuring Ilks, urns, gem, dai, and sin
var (
	debtDesc = prometheus.NewDesc(
		"mkr_vat_debt",
		"Maker Vat debt casted from a uint256 value (rad)",
		nil, nil,
	)
	viceDesc = prometheus.NewDesc(
		"mkr_vat_vice",
		"Maker Vat vice casted from a uint256 value (rad)",
		nil, nil,
	)
	lineDesc = prometheus.NewDesc(
		"mkr_vat_line",
		"Maker Vat line casted from a uint256 (rad)",
		nil, nil,
	)
)

type vatCollector struct {
	vat *maker.VatCaller
}

// NewVatCollector reates a new collecotr for collecting Vat specific stats
func NewVatCollector(vat *maker.VatCaller) prometheus.Collector {
	return &vatCollector{vat}
}

func (c *vatCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- debtDesc
	ch <- viceDesc
	ch <- lineDesc
}

func (c *vatCollector) Collect(ch chan<- prometheus.Metric) {
	if _, err := c.vat.Debt(nil); err == nil {
		ch <- prometheus.MustNewConstMetric(
			debtDesc,
			prometheus.GaugeValue,
			0.0, // FIXME: Need to determine the correct way to convert a Rad to a float64
		)
	}
}